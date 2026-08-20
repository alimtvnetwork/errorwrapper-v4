package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

func StartMultipleServices(
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
		servicestate.Start,
		servicesNames...)
}
