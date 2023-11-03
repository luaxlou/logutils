package logutils

import (
	"errors"
	"testing"
)

func TestPrint(t *testing.T) {
	type args struct {
		obj interface{}
		msg string
	}

	type x struct {
		name string
	}
	PrintObjWithMsg(x{"xxx"}, "hi")

	PrintError(errors.New("I am creazy!!"), "haha")
}
