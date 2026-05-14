package errcmd

import (
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func getNotImplementedCmdOnceForScript(scriptDefaultStringer coreinterface.Stringer) *CmdOnce {
	notImplErr := errnew.Messages.Many(
		errtype.NotImplemented,
		scriptDefaultStringer.String())

	return &CmdOnce{
		baseCmdWrapper: BaseCmdWrapper{
			initializeErrorWrapper: notImplErr,
		},
		Cmd: nil,
	}
}
