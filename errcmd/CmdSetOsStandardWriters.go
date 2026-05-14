package errcmd

import (
	"os"
	"os/exec"
)

func CmdSetOsStandardWriters(cmd *exec.Cmd) *exec.Cmd {
	if cmd == nil {
		return nil
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
