package pro

import (
	"errors"
	"os/exec"
	_ "unsafe" // for go:linkname
)

// Deprecated: 暂无法正常使用
//go:linkname cmdExec github.com/coulsonzero/gopkg/pro.Command
// CmdExec exec the cmd command
// eg: CmdExec("rm -f test.txt")
func cmdExec(command string) error {
	err := exec.Command("bash", "-c", command).Run()
	if err != nil {
		return errors.New("error: exec the cmd command")
	}
	return nil
}
