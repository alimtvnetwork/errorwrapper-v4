package errcmd

import (
	"os/exec"

	"github.com/alimtvnetwork/core-v9/coreinterface"
)

func GetExitCode(
	errorWrapper coreinterface.IsEmptyChecker,
	exitError *exec.ExitError,
) int {
	if errorWrapper.IsEmpty() && exitError == nil {
		return SuccessfullyRunningExitCode
	}

	exitCode := InvalidExitCode

	if exitError != nil {
		exitCode = exitError.ExitCode()
	}

	return exitCode
}
