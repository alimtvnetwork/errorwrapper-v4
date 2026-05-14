package errcmd

import (
	"github.com/alimtvnetwork/errorwrapper-v3/internal/consts"
)

const (
	InvalidExitCode             = consts.InvalidExitCode
	SuccessfullyRunningExitCode = consts.CmdSuccessfullyRunningExitCode
	ScriptsMultiLineJoiner      = " && \\ \n" // " && \\ \n"
	SingleLineScriptsJoiner     = " && "      // " && "
	changeDirSpace              = "cd "
)
