package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func RunServicesResults(
	isDetailedError,
	isContinueOnErr bool,
	isIgnoreUnknownService bool,
	errCollection *errwrappers.Collection,
	action servicestate.Action,
	servicesNames ...string,
) *Results {
	results := make([]*Result, len(servicesNames))

	for i, serviceName := range servicesNames {
		result := Run(action, serviceName)

		if isIgnoreUnknownService && result.IsUnknownService() {
			continue
		}

		err := result.ErrorWrapperUsingOpt(isDetailedError)
		results[i] = result
		errCollection.AddWrapperPtr(err)

		if !isContinueOnErr && err.HasError() {
			return &Results{Items: results}
		}
	}

	return &Results{Items: results}
}
