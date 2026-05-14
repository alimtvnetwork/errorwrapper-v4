package errcmd

import "gitlab.com/evatix-go/core/constants"

func failedCmdOneOutputStringProcessor(
	cmdOnce *CmdOnce,
) (processedLine string, isTake, isBreak bool) {
	if cmdOnce.HasAnyIssues() {
		return cmdOnce.CompiledOutput(), true, false
	}

	return constants.EmptyString, false, false
}
