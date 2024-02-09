package prokishi

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"golang.org/x/xerrors"
)

// クライアントを生成し、USIとサーバの連携を行う
func Run(host string, port int, opts ...Option) error {

	conf := getDefaultConfig(host, port)
	var optE error
	for _, opt := range opts {
		err := opt(conf)
		if err != nil {
			optE = errors.Join(optE, err)
		}
	}
	if optE != nil {
		return xerrors.Errorf("Options error: %w", optE)
	}
	setConfig(conf)

	ctx := context.Background()
	quit := make(chan os.Signal, 1)

	cli, err := NewClient(ctx, conf, quit)
	if err != nil {
		return xerrors.Errorf("NewClient() error: %w", err)
	}
	defer cli.disconnect()

	signal.Notify(quit, os.Interrupt)
	<-quit

	cli.sendQuit()

	return nil
}
