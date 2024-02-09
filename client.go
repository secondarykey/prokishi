package prokishi

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"prokishi/api"
	"prokishi/usi"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	ctx       context.Context
	conn      *grpc.ClientConn
	recvUSI   *usi.Receiver
	senderCli api.USISendServiceClient
	quit      chan os.Signal

	connectionId string
}

// クライアントの生成
// サーバと接続を行い、UI側との連携を生成
func NewClient(ctx context.Context, conf *Config, quit chan os.Signal) (*Client, error) {

	var cli Client

	cli.ctx = ctx
	cli.quit = quit

	err := cli.dial(conf.Host, conf.Port)
	if err != nil {
		return nil, xerrors.Errorf("dial() error: %w", err)
	}

	err = cli.connect(conf)
	if err != nil {
		return nil, xerrors.Errorf("connect() error: %w", err)
	}

	cli.recvUSI = usi.NewReceiver(os.Stdout, os.Stdin)
	go cli.receiveUSI(conf)

	return &cli, nil
}

// サーバとの接続
func (cli *Client) dial(host string, port int) error {

	addr := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return xerrors.Errorf("grpc.Dial() error: %w", err)
	}
	cli.conn = conn

	return nil
}

// 連携用の接続
func (cli *Client) connect(c *Config) error {

	connCli := api.NewConnectionServiceClient(cli.conn)
	req := &api.ConnectionRequest{
		Code:     c.Code,
		EngineId: c.EngineId,
	}

	//サーバからコネクションIDを取得
	res, err := connCli.Connection(cli.ctx, req)
	if err != nil {
		return xerrors.Errorf("api.Connection() error: %w", err)
	}

	cli.connectionId = res.ConnectionId
	recvReq := &api.ReceiveRequest{
		Code:         c.Code,
		ConnectionId: cli.connectionId,
	}

	//サーバ受信用のストリームを生成
	recvCli := api.NewUSIReceiveServiceClient(cli.conn)
	stream, err := recvCli.Receive(cli.ctx, recvReq)
	if err != nil {
		return xerrors.Errorf("api.Receive() error: %w", err)
	}
	//サーバの監視
	go cli.receiveServer(stream, c.Version)

	//送信用のクライアントを生成
	cli.senderCli = api.NewUSISendServiceClient(cli.conn)

	return nil
}

// 接続しているかを確認
func (cli *Client) isConnect() bool {
	if cli.connectionId == "" {
		return false
	}
	return true
}

// 接続を切る
func (cli *Client) disconnect() {
	cli.conn.Close()
	cli.connectionId = ""
	return
}

// USIから来た出力をサーバ側に送信
func (cli *Client) receiveUSI(conf *Config) {
	for in := range cli.recvUSI.InCh {
		err := cli.sendServer(conf, in)
		if err != nil {
			slog.Error(fmt.Sprintf("%+v", err))
		}
	}
}

// 終了をサーバに送信
func (cli *Client) sendQuit() {
	if cli.senderCli == nil {
		return
	}
	cli.sendServer(getConfig(), "quit")
}

// サーバに送信を行う
func (cli *Client) sendServer(conf *Config, cmd string) error {

	quit := false
	if cmd == "quit" {
		quit = true
	}

	slog.Debug(fmt.Sprintf("USI(I):%s", cmd))
	req := &api.SendRequest{
		Code:         conf.Code,
		ConnectionId: cli.connectionId,
		Cmd:          cmd,
	}

	//コマンドを送信
	_, err := cli.senderCli.Send(cli.ctx, req)
	if err != nil {
		slog.Error(fmt.Sprintf("%+v", err))
	}

	//終了フラグを設定
	if quit {
		cli.connectionId = ""
		cli.senderCli = nil
		cli.quit <- os.Interrupt
	}

	return nil
}

// サーバから受信を行い、標準出力を行う
// id name,id authorに関しては編集を行う
func (cli *Client) receiveServer(stream api.USIReceiveService_ReceiveClient, v string) {
	for {

		res, err := stream.Recv()
		if err != nil {
			break
		}

		cmd := res.Cmd
		//以下を編集
		if strings.Index(cmd, "id name") == 0 {
			cmd = fmt.Sprintf(cmd+"(%s)", "prokishi "+v)
		} else if strings.Index(cmd, "id author") == 0 {
			cmd = fmt.Sprintf(cmd+"(%s)", "secondarykey")
		}

		slog.Debug(fmt.Sprintf("USI(O): %s\n", cmd))
		//UIエンジンに送信
		cli.recvUSI.Send(cmd)
	}
}
