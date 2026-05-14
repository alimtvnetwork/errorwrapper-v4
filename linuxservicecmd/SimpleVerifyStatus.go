package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3"
)

func SimpleVerifyStatus(serviceName string) *errorwrapper.Wrapper {
	res := Run(
		servicestate.Status,
		serviceName)

	return res.SimplifiedError()
}
