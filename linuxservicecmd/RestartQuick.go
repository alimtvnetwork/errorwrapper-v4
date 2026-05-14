package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper"
)

func RestartQuick(
	servicesName string,
) *errorwrapper.Wrapper {
	_, errWp := RunAction(
		true,
		false,
		servicestate.Restart,
		servicesName)

	return errWp
}
