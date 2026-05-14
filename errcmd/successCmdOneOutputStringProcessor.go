package errcmd

import "gitlab.com/evatix-go/core/constants"

func successCmdOneOutputStringProcessor(
	cmdOnce *CmdOnce,
) (processedLine string, isTake, isBreak bool) {
	if cmdOnce.IsSuccessfullyExecuted() {
		return cmdOnce.CompiledOutput(), true, false
	}

	return constants.EmptyString, false, false
}
