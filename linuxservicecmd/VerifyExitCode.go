package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/linuxservicestate"
	"gitlab.com/evatix-go/errorwrapper"
)

func VerifyExitCode(
	serviceName string,
	expectedExitCode linuxservicestate.ExitCode,
) *errorwrapper.Wrapper {
	return GetStatusResult(serviceName).
		VerifyExitCode(expectedExitCode)
}
