package errfunc

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
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
