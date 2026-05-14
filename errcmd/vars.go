package errcmd

import (
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
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
