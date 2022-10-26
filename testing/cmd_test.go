package testing

import (
	"github.com/coulsonzero/gopkg/pro"
	"testing"
)

func TestCmdExec(t *testing.T) {
	if err := pro.CmdExec("ls"); err != nil {
		t.Error("error: exec cmd command !")
	}
}

func TestCmdOutput(t *testing.T) {
	if _, err := pro.CmdOutput("ls"); err != nil {
		t.Error("error: print cmd output !")
	}
}
