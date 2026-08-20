package errcmd

import (
	"github.com/alimtvnetwork/errorwrapper-v4/internal/consts"
)

const (
	InvalidExitCode             = consts.InvalidExitCode
	SuccessfullyRunningExitCode = consts.CmdSuccessfullyRunningExitCode
	ScriptsMultiLineJoiner      = " && \\ \n" // " && \\ \n"
	SingleLineScriptsJoiner     = " && "      // " && "
	changeDirSpace              = "cd "
)
