package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper"
)

func RunAction(
	isDetailedError,
	isIgnoreUnknownService bool,
	action servicestate.Action,
	servicesName string,
) (*Result, *errorwrapper.Wrapper) {
	if isIgnoreUnknownService && IsUnknownServiceExist(servicesName) {
		return nil, nil
	}

	result := Run(action, servicesName)
	err := result.ErrorWrapperUsingOpt(isDetailedError)

	return result, err
}
