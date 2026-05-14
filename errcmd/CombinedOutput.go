package errcmd

import (
	"os/exec"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

// CombinedOutputError
//
// stdIn can be nil
func CombinedOutputError(
	stdIn *StdIn,
	processName string,
	args ...string,
) *cmdCompiledOutput {
	wholeCmdLines := stringslice.PrependLineNew(processName, args)
	wholeCmdString := strings.Join(wholeCmdLines, constants.Space)
	cmd := exec.Command(processName, args...)

	if cmd == nil {
		err := errtype.InvalidProcess.
			Error(
				"cmd nil",
				"command:",
				wholeCmdString)

		return InvalidCmdCompiledOutput(err, processName, args)
	}

	if stdIn.HasStdOut() {
		cmd.Stdout = stdIn.OutBuf
	}

	if stdIn.HasStdErr() {
		cmd.Stderr = stdIn.ErrBuf
	}

	allBytes, err := cmd.CombinedOutput()

	if err != nil {
		return cmdOutputForError(
			stdIn,
			wholeCmdString,
			allBytes,
			cmd)
	}

	return &cmdCompiledOutput{
		Cmd:         cmd,
		Output:      corestr.New.SimpleStringOnce.Init(string(allBytes)),
		Error:       nil,
		ProcessName: cmd.Path,
		Arguments:   cmd.Args,
	}
}
