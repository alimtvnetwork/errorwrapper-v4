package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

func VerifyExitCode(
	serviceName string,
	expectedExitCode linuxservicestate.ExitCode,
) *errorwrapper.Wrapper {
	return GetStatusResult(serviceName).
		VerifyExitCode(expectedExitCode)
}
