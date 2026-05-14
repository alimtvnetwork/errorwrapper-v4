package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/enum-v10/servicestate"
)

func GetStatus(serviceName string) linuxservicestate.ExitCode {
	res := Run(
		servicestate.Status,
		serviceName)

	return res.ExitCode
}
