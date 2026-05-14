package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3"
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
