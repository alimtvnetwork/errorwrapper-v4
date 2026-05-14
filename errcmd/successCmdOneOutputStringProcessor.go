package errcmd

import "github.com/alimtvnetwork/core-v9/constants"

func successCmdOneOutputStringProcessor(
	cmdOnce *CmdOnce,
) (processedLine string, isTake, isBreak bool) {
	if cmdOnce.IsSuccessfullyExecuted() {
		return cmdOnce.CompiledOutput(), true, false
	}

	return constants.EmptyString, false, false
}
