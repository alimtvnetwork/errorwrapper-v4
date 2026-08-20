package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

func StartQuick(
	servicesName string,
) *errorwrapper.Wrapper {
	_, errWp := RunAction(
		true,
		false,
		servicestate.Start,
		servicesName)

	return errWp
}
