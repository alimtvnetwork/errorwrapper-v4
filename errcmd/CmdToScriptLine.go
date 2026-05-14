package errcmd

import "os/exec"

func CmdToScriptLine(cmd *exec.Cmd) string {
	if cmd == nil || cmd.Path == "" {
		return ""
	}

	return ProcessArgsJoinAppend(
		cmd.Path,
		cmd.Args...)
}
