package errcmd

import (
	"os/exec"

	"gitlab.com/evatix-go/errorwrapper"
)

type newCmdCreator struct{}

func (it newCmdCreator) Cmd(
	process string,
	arguments ...string,
) (cmd *exec.Cmd, wholeCommand string) {
	cmd = exec.Command(
		process, arguments...)

	return cmd, ProcessArgsJoinAppend(
		process, arguments...)
}

func (it newCmdCreator) CmdOsStd(
	process string,
	arguments ...string,
) (cmd *exec.Cmd, wholeCommand string) {
	cmd = exec.Command(
		process, arguments...)

	cmd = CmdSetOsStandardWriters(cmd)

	return cmd, ProcessArgsJoinAppend(
		process, arguments...)
}

func (it newCmdCreator) CompiledErrorWrapper(
	process string,
	arguments ...string,
) *errorwrapper.Wrapper {
	_, errWp := New.OutputGetter.OutputErrWrapper(
		process,
		arguments...)

	return errWp
}
