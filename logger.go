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

// 開発時は作業ディレクトリ、
// リリース動作の場合は実行しているパスを取得
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

func createLvLogger(lv slog.Level, w io.Writer) *slog.Logger {
	var opts slog.HandlerOptions
	opts.Level = lv
	h := slog.NewTextHandler(w, &opts)
	return slog.New(h)
}

func SetLog(lv slog.Level, w io.Writer) {
	slog.SetDefault(createLvLogger(lv, w))
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

	slog.SetDefault(createLvLogger(lv, fp))
	return fp
}

func timestamp() string {
	now := time.Now()
	return now.Format("20060102150405")
}
