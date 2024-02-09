package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"prokishi/api"
	"prokishi/usi"
	"strconv"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

func Run(host string, port int, opts ...Option) error {

	addr := net.JoinHostPort(host, strconv.Itoa(port))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return xerrors.Errorf("net.Listen() error: %w", err)
	}

	fmt.Println("IP Address ====================")
	printIPs()
	fmt.Println("===============================")
	fmt.Println("Listener Address:", listener.Addr())

	s := grpc.NewServer()
	srv := RegisterServiceServer(s)

	go func() {
		err := s.Serve(listener)
		if err != nil {
			log.Printf("Serve() error: %v", err)
		}
	}()

	srv.Wait()

	s.GracefulStop()
	return nil
}

func printIPs() {
	ift, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ifi := range ift {
		addrs, err := ifi.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, addr := range addrs {
			ip := getIP(addr)
			if !ip.IsLoopback() {
				fmt.Printf("%v\n", ip)
			}
		}
	}
}

func getIP(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	return ip
}

type Server struct {
	api.ConnectionServiceServer
	api.USISendServiceServer
	api.USIReceiveServiceServer

	engines *sync.Map
}

// GRPCサービスを登録
func RegisterServiceServer(r grpc.ServiceRegistrar) *Server {

	var serv Server

	api.RegisterConnectionServiceServer(r, &serv)
	api.RegisterUSISendServiceServer(r, &serv)
	api.RegisterUSIReceiveServiceServer(r, &serv)

	serv.engines = &sync.Map{}
	return &serv
}

// 認証
func (s *Server) verifyAuthentication(code string) bool {
	if code == "testCode" {
		//return true
	}
	return true
}

// コネクションIDでエンジンを実行し登録する
func (s *Server) startEngine(id string) (string, error) {

	if id != "testEngine" {
		//return "", fmt.Errorf("engine not found")
	}

	//エンジンを確認
	uid := uuid.New()
	rtn := uid.String()

	name := "D:\\Program Files\\ShogiGUI\\Suisho5-YaneuraOu-v7.5.0-windows\\YaneuraOu_NNUE-tournament-clang++-avx2.exe"
	//name := "D:\\Program Files\\ShogiGUI\\水匠5\\Suisho5-AVX2.exe"

	engine, err := usi.NewSender(name)
	if err != nil {
		return "", xerrors.Errorf("usi.NewSender() error: %w", err)
	}

	s.engines.Store(rtn, engine)
	return rtn, nil
}

func (s *Server) Wait() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
