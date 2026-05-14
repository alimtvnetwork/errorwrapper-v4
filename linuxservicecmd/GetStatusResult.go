package linuxservicecmd

import "gitlab.com/evatix-go/enum/servicestate"

func GetStatusResult(serviceName string) *Result {
	return Run(
		servicestate.Status,
		serviceName)
}
