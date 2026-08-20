package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

func ConvertWrapperFuncToIsSuccessCollectorFunc(
	wrapperFunc WrapperFunc,
) IsSuccessCollectorFunc {
	if wrapperFunc == nil {
		return nil
	}

	isSuccessCollectorFunc := func(errorCollection *errwrappers.Collection) (isSuccess bool) {
		errWrapper := wrapperFunc()

		errorCollection.AddWrapperPtr(errWrapper)

		return errWrapper.IsEmptyError()
	}

	return isSuccessCollectorFunc
}
