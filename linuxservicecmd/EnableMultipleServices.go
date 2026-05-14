package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

func EnableMultipleServices(
	isDetailedError,
	isContinueOnErr bool,
	errCollection *errwrappers.Collection,
	servicesNames ...string,
) (isSuccess bool) {
	return RunServices(
		isDetailedError,
		isContinueOnErr,
		false,
		errCollection,
		servicestate.Enable,
		servicesNames...)
}
