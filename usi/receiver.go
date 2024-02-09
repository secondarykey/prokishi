package usi

import (
	"bufio"
	"io"

	"golang.org/x/xerrors"
)

// UIからのレシーバ
// 標準出力と標準入力を監視
type Receiver struct {
	w    io.Writer
	r    io.Reader
	InCh chan string
}

// エンジン動作時はos.Stdout,os.Stdinのはず
func NewReceiver(w io.Writer, r io.Reader) *Receiver {

	var recv Receiver
	recv.w = w
	recv.r = r

	recv.InCh = make(chan string)
	go recv.monitorStdin()

	return &recv
}

// 入力を監視し、チャンネルに送信
func (r *Receiver) monitorStdin() {
	scanner := bufio.NewScanner(r.r)
	for scanner.Scan() {
		r.InCh <- scanner.Text()
	}
}

// コマンドを送信(出力に書き込む)
func (r *Receiver) Send(cmd string) error {
	_, err := r.w.Write([]byte(cmd + "\n"))
	if err != nil {
		return xerrors.Errorf("Write() error: %w", err)
	}
	return nil
}
