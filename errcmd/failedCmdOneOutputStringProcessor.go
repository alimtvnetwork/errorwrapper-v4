package errcmd

import "github.com/alimtvnetwork/core-v9/constants"

func failedCmdOneOutputStringProcessor(
	cmdOnce *CmdOnce,
) (processedLine string, isTake, isBreak bool) {
	if cmdOnce.HasAnyIssues() {
		return cmdOnce.CompiledOutput(), true, false
	}

	return constants.EmptyString, false, false
}
