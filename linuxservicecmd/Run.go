package linuxservicecmd

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/enum/linuxservicestate"
	"gitlab.com/evatix-go/enum/servicestate"
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
