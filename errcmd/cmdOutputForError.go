package errcmd

import (
	"os/exec"

	"gitlab.com/evatix-go/core/conditional"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func cmdOutputForError(
	stdIn *StdIn,
	wholeCmdString string,
	allBytes []byte,
	cmd *exec.Cmd,
) *cmdCompiledOutput {
	hasStdOut := stdIn.HasStdOut()
	hasStdErr := stdIn.HasStdErr()

	stdOutString := conditional.StringTrueFunc(
		hasStdOut,
		stdIn.StdOutString,
	)

	stdErrString := conditional.StringTrueFunc(
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
