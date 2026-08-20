package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
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
