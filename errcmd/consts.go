package errcmd

import (
	"gitlab.com/evatix-go/errorwrapper/internal/consts"
)

const (
	InvalidExitCode             = consts.InvalidExitCode
	SuccessfullyRunningExitCode = consts.CmdSuccessfullyRunningExitCode
	ScriptsMultiLineJoiner      = " && \\ \n" // " && \\ \n"
	SingleLineScriptsJoiner     = " && "      // " && "
	changeDirSpace              = "cd "
)
