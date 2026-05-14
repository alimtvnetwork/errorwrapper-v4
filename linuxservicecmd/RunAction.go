package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3"
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
