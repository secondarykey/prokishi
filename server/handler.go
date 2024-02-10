package server

import (
	"context"
	"fmt"
	"log/slog"

	"prokishi/api"
	"prokishi/usi"

	"golang.org/x/xerrors"
)

// コネクションIDを発行
func (s *Server) Connection(ctx context.Context, r *api.ConnectionRequest) (*api.ConnectionResponse, error) {

	//CodeがOKかを確認
	if !s.verifyAuthentication(r.Code) {
		return nil, fmt.Errorf("can not be verified")
	}

	//EngineIDがOKかを確認
	id, err := s.startEngine(r.EngineId)
	if err != nil {
		return nil, xerrors.Errorf("startEngine() error: %w", err)
	}

	var res api.ConnectionResponse
	res.ConnectionId = id

	slog.Info(fmt.Sprintf("register Engine map:%s", id))

	return &res, nil
}

// エンジンに送信
func (s *Server) Send(ctx context.Context, r *api.SendRequest) (*api.SendResponse, error) {

	e, err := s.getEngine(r.Code, r.ConnectionId)
	if err != nil {
		return nil, xerrors.Errorf("getEngine() error: %w")
	}
	quit := false
	if r.Cmd == "quit" {
		quit = true
	}

	slog.Debug(fmt.Sprintf("USI(I):%s", r.Cmd))
	err = e.Send(r.Cmd)
	if err != nil {
		return nil, xerrors.Errorf("Send() error: %w")
	}

	if quit {
		slog.Info(fmt.Sprintf("remove Engine map:%s", r.ConnectionId))
		s.engines.Delete(r.ConnectionId)
	}

	var res api.SendResponse
	return &res, nil
}

// エンジンの出力をクライアントに送信
func (s *Server) Receive(r *api.ReceiveRequest, stream api.USIReceiveService_ReceiveServer) error {

	//エンジンを取得
	e, err := s.getEngine(r.Code, r.ConnectionId)
	if err != nil {
		return xerrors.Errorf("getEngine() error: %w")
	}

	//エンジンの出力を監視
	for v := range e.OutCh {
		slog.Debug(fmt.Sprintf("USI(O):%s", v))
		//ストリームに送る
		stream.Send(&api.ReceiveResponse{
			Cmd: v,
		})
	}

	return nil
}

// エンジン取得
func (s *Server) getEngine(code string, id string) (*usi.Sender, error) {
	//CodeがOKかを確認
	if !s.verifyAuthentication(code) {
		return nil, fmt.Errorf("can not be verified")
	}

	e, ok := s.engines.Load(id)
	if !ok {
		return nil, fmt.Errorf("connectionId is failed")
	}
	return e.(*usi.Sender), nil
}
