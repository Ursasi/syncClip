package service

import (
	"github.com/atotto/clipboard"
	"syncClip/util"
)

func Send() (bool, error) {
	return false, nil
}

func Receive(request util.ReceiveRequest) (bool, error) {
	msg := request.Msg
	return Set(msg)
}

func Set(msg string) (bool, error) {
	err := clipboard.WriteAll(msg)
	if err != nil {
		return false, err
	}
	return true, nil
}
