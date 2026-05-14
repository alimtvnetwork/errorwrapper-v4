package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

func ConvertLinuxErrorCollectorActionToErrWrapperFunc(
	action LinuxErrorCollectorAction,
) func() *errorwrapper.Wrapper {
	return func() *errorwrapper.Wrapper {
		errCollection := errwrappers.Empty()
		action.CollectorFunc(errCollection)

		if errCollection.HasAnyItem() {
			return errCollection.First()
		}

		return nil
	}
}
