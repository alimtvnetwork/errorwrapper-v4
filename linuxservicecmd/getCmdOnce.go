package linuxservicecmd

import (
	"gitlab.com/evatix-go/core/cmdconsts"
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func getCmdOnce(
	action servicestate.Action,
	serviceName string,
) *errcmd.CmdOnce {
	actionName := action.Name()

	if hasSystemCtlService {
		return bashArgsCmdCreator(
			cmdconsts.SystemCtl,
			actionName,
			serviceName,
		)
	}

	if hasService {
		return bashArgsCmdCreator(
			cmdconsts.Service,
			serviceName,
			actionName)
	}

	return bashArgsCmdCreator(
		cmdconsts.SystemCtl,
		actionName,
		serviceName,
	)
}
