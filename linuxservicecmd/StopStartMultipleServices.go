package linuxservicecmd

import (
	"time"

	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func StopStartMultipleServices(
	isDetailedError,
	isContinueOnErr bool,
	errCollection *errwrappers.Collection,
	sleep time.Duration,
	servicesNames ...string,
) (isSuccess bool) {
	isSuccess = StopMultipleServices(
		isDetailedError,
		isContinueOnErr,
		errCollection,
		servicesNames...)

	time.Sleep(sleep)

	if !isContinueOnErr && !isSuccess {
		return false
	}

	isSuccess = StartMultipleServices(
		isDetailedError,
		isContinueOnErr,
		errCollection,
		servicesNames...) && isSuccess

	return isSuccess
}
