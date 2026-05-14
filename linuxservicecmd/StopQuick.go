package linuxservicecmd

import (
	"gitlab.com/evatix-go/enum/servicestate"
	"gitlab.com/evatix-go/errorwrapper"
)

func StopQuick(
	servicesName string,
) *errorwrapper.Wrapper {
	_, errWp := RunAction(
		true,
		false,
		servicestate.Stop,
		servicesName)

	return errWp
}
