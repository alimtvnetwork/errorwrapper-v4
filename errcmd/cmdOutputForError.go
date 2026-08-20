package errcmd

import (
	"os/exec"

	"github.com/alimtvnetwork/core-v9/conditional"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

func cmdOutputForError(
	stdIn *StdIn,
	wholeCmdString string,
	allBytes []byte,
	cmd *exec.Cmd,
) *cmdCompiledOutput {
	hasStdOut := stdIn.HasStdOut()
	hasStdErr := stdIn.HasStdErr()

	stdOutString := conditional.IfTrueFuncString(
		hasStdOut,
		stdIn.StdOutString,
	)

	stdErrString := conditional.IfTrueFuncString(
		hasStdOut,
		stdIn.StdErrString)

	references := constants.EmptyString

	if hasStdErr || hasStdOut {
		references = errcore.VarTwoNoType(
			"OutBuffer", stdOutString,
			"ErrBuffer", stdErrString)
	}

	errFinal := errtype.FailedProcess.ErrorReferences(
		"Command:"+wholeCmdString,
		references)

	return &cmdCompiledOutput{
		Cmd:         cmd,
		Output:      corestr.New.SimpleStringOnce.Init(string(allBytes)),
		Error:       errFinal,
		ProcessName: cmd.Path,
		Arguments:   cmd.Args,
	}
}
