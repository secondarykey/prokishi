package prokishi

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/xerrors"
)

func GetRunDir(d bool) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", xerrors.Errorf("os.Getwd() error: %w", err)
	}
	//開発モードじゃない場合、実行位置に変更
	if !d {
		exe, err := os.Executable()
		if err != nil {
			return "", xerrors.Errorf("os.Executable() error: %w", err)
		}
		dir = filepath.Dir(exe)
	}
	return dir, nil
}

func SetLogFile(lv slog.Level, name string, d bool) io.Closer {

	dir, err := GetRunDir(d)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	//削除したりの運用面倒そうだから同じ名前にする
	fn := filepath.Join(dir, fmt.Sprintf("%s.log", name))

	fp, err := os.Create(fn)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	var opts slog.HandlerOptions
	opts.Level = lv

	h := slog.NewTextHandler(fp, &opts)
	slog.SetDefault(slog.New(h))

	return fp
}

func timestamp() string {
	now := time.Now()
	return now.Format("20060102150405")
}
