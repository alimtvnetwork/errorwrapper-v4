package linuxservicecmd

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/enum-v10/servicestate"
)

func Run(
	action servicestate.Action,
	serviceName string,
) *Result {
	cmd := getCmdOnce(action, serviceName)
	cmdResult := cmd.CompiledResult()
	compiledCode := linuxservicestate.NewCode(cmdResult.ExitCode)
	errWrapper := cmdResult.ErrorWrapper()

	if errWrapper.HasError() {
		errWrapper = errWrapper.
			ConcatNew().
			MsgRefTwo(
				codestack.Skip1,
				"Service execution failed",
				"Service",
				serviceName,
				"Action",
				action.NameValue())
	}

	return &Result{
		Request: Request{
			ServiceName: serviceName,
			Action:      action,
		},
		CmdOnce:      cmd,
		ExitCode:     compiledCode,
		ErrorWrapper: errWrapper,
	}
}
