package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper"
)

func Restart(
	isDetailedError,
	isIgnoreUnknownService bool,
	action servicestate.Action,
	servicesName string,
) (*Result, *errorwrapper.Wrapper) {
	return RunAction(
		isDetailedError,
		isIgnoreUnknownService,
		action,
		servicesName)
}
