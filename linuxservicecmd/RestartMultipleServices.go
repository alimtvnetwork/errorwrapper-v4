package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func RestartMultipleServices(
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
		servicestate.Restart,
		servicesNames...)
}
