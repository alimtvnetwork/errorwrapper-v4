package errcmd

import (
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/enum/scripttype"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

var (
	NewShellScript = &newCmdOnceTypedScriptsCreator{
		scriptType: scripttype.Shell,
	}
	NewBashScript = &newCmdOnceTypedScriptsCreator{
		scriptType: scripttype.Bash,
	}

	New = &newCreator{
		ShellScript: NewShellScript,
		BashScript:  NewBashScript,
	}

	cmdNilErr = errnew.Messages.Many(
		errtype.CommandExecutionNotFound,
		errcore.FailedToCreateCmdType.String(),
		"Create() *CmdOnce",
		"Failed to create inner cmd.",
	)
)
