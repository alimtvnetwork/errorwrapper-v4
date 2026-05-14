package linuxservicecmd

import "github.com/alimtvnetwork/enum-v10/servicestate"

func GetStatusResult(serviceName string) *Result {
	return Run(
		servicestate.Status,
		serviceName)
}
