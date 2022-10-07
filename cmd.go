package gopkg

import (
	"errors"
	"os/exec"
)

// CmdExec exec the cmd command
// eg: CmdExec("rm -f test.txt")
func CmdExec(command string) error {
	err := exec.Command("bash", "-c", command).Run()
	if err != nil {
		return errors.New("error: exec the cmd command")
	}
	return nil
}
