package usi_test

import (
	"fmt"
	"prokishi/usi"
	"testing"
	"time"
)

func TestSender(t *testing.T) {

	s, err := usi.NewSender("../_cmd/prokishi.exe")
	if err != nil {
		t.Errorf("NewSencer() error is not nil: %v", err)
	}

	//defer s.Terminate()

	go func() {
		for now := range time.Tick(2 * time.Second) {
			s.Send(fmt.Sprintf("%v", now))
		}
	}()

	for txt := range s.OutCh {
		fmt.Println(txt)
	}

}
