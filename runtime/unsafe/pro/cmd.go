package pro

import (
	"errors"
	"os/exec"
	_ "unsafe" // for go:linkname
)

//go:linkname cmdExec github.com/coulsonzero/gopkg/pro.CmdExec
func cmdExec(command string) error {
	err := exec.Command("bash", "-c", command).Run() // Run() wait the result, Start() cannot
	if err != nil {
		return errors.New("error: exec the cmd command")
	}
	return nil
}

//go:linkname cmdOutput github.com/coulsonzero/gopkg/pro.CmdOutput
func cmdOutput(command string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", command)
	res, err := cmd.Output()
	if err != nil {
		return nil, errors.New("exec: Stdout already set")
	}
	return res, nil
}
