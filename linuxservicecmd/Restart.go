package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v4"
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
