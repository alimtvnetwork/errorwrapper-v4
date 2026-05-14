package errfunc

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc(
	action LinuxIsSuccessPlusErrorCollectAction,
) func() *errorwrapper.Wrapper {
	return func() *errorwrapper.Wrapper {
		errCollection := errwrappers.Empty()

		if !action.IsSuccessCollectorFunc(errCollection) {
			return errCollection.First()
		}

		return nil
	}
}
