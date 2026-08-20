package linuxservicecmd

import (
	"github.com/alimtvnetwork/core-v9/cmdconsts"
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v4/errcmd"
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
