package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func RunServices(
	isDetailedError,
	isContinueOnErr bool,
	isIgnoreUnknownService bool,
	errCollection *errwrappers.Collection,
	action servicestate.Action,
	servicesNames ...string,
) (isSuccess bool) {
	stateTracker := errCollection.StateTracker()

	for _, serviceName := range servicesNames {
		if isIgnoreUnknownService && IsUnknownServiceExist(serviceName) {
			continue
		}

		result := Run(action, serviceName)
		err := result.ErrorWrapperUsingOpt(isDetailedError)
		errCollection.AddWrapperPtr(err)

		if !isContinueOnErr && err.HasError() {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
