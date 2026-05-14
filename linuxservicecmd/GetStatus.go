package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/linuxservicestate"
	"gitlab.com/evatix-go/enum/servicestate"
)

func GetStatus(serviceName string) linuxservicestate.ExitCode {
	res := Run(
		servicestate.Status,
		serviceName)

	return res.ExitCode
}
