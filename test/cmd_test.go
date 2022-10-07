package test

import (
	"github.com/coulsonzero/gopkg"
	"testing"
)

func TestCmd(t *testing.T) {
	if err := gopkg.CmdExec("ls"); err != nil {
		t.Error("error: exec cmd command !")
	}
}
