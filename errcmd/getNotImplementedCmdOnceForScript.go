package errcmd

import (
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
