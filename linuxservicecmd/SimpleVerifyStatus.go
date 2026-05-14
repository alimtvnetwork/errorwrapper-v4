package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper"
)

func SimpleVerifyStatus(serviceName string) *errorwrapper.Wrapper {
	res := Run(
		servicestate.Status,
		serviceName)

	return res.SimplifiedError()
}
