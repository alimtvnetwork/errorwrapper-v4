package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper"
)

func ReloadQuick(
	servicesName string,
) *errorwrapper.Wrapper {
	_, errWp := RunAction(
		true,
		false,
		servicestate.Reload,
		servicesName)

	return errWp
}
