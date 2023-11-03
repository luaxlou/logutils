package test

import (
	"errors"
	"github.com/luaxlou/logutils"
	"testing"
)

func TestPrint(t *testing.T) {

	type x struct {
		name string
	}
	logutils.PrintObjWithMsg(x{"xxx"}, "hi")

	logutils.PrintError(errors.New("I am creazy!!"), "haha")
	logutils.PrintErrorMsg("I am creazy too!!", "haha")
}
