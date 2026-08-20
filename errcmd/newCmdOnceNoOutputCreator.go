package errcmd

import (
	"os/exec"

	"github.com/alimtvnetwork/errorwrapper-v4"
)

type newCmdOnceNoOutputCreator struct{}

func (it *newCmdOnceNoOutputCreator) ExitCode(
	process string,
	arguments ...string,
) int {
	rs := New.Create(
		false,
		true,
		process,
		arguments...).
		CompiledResult()
	exitCode := rs.ExitCode
	go rs.Dispose()

	return exitCode
}

func (it *newCmdOnceNoOutputCreator) ExitCodeAsByte(
	process string,
	arguments ...string,
) byte {
	rs := New.Create(
		false,
		true,
		process,
		arguments...).
		CompiledResult()
	exitCode := rs.ExitCodeByte()
	go rs.Dispose()

	return exitCode
}

// Create (no output for error and std out)
//
// hasOutput false
// hasSecure true
func (it *newCmdOnceNoOutputCreator) Create(
	process string,
	arguments ...string,
) *CmdOnce {
	return New.Create(
		false,
		true,
		process,
		arguments...)
}

func (it *newCmdOnceNoOutputCreator) Result(
	process string,
	arguments ...string,
) *Result {
	cmdOnce := New.Create(
		false,
		true,
		process,
		arguments...)

	rs := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()

	return rs
}

func (it *newCmdOnceNoOutputCreator) ErrorWrapper(
	process string,
	arguments ...string,
) *errorwrapper.Wrapper {
	return New.Create(
		false,
		false,
		process,
		arguments...).
		CompiledFullErrorWrapper()
}

func (it *newCmdOnceNoOutputCreator) UsingCmdNoOutputVerboseError(
	cmd *exec.Cmd,
) *CmdOnce {
	return New.UsingCmd(
		false,
		false,
		cmd)
}

// CreateErrorVerbose (consider error std) alias for Default
//
// hasOutput false
// hasSecure false (error will verbose the error output)
func (it *newCmdOnceNoOutputCreator) CreateErrorVerbose(
	process string,
	arguments ...string,
) *CmdOnce {
	return New.Create(
		false,
		false,
		process,
		arguments...)
}

// Default (consider error std) alias for CreateErrorVerbose
//
// hasOutput false
// hasSecure false (error will verbose the error output)
func (it *newCmdOnceNoOutputCreator) Default(
	process string,
	arguments ...string,
) *CmdOnce {
	return New.Create(
		false,
		false,
		process,
		arguments...)
}
