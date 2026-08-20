package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

func ConvertErrorFuncToIsSuccessCollectorFunc(
	errorType errtype.Variation,
	simpleErrFunc SimpleErrorFunc,
) IsSuccessCollectorFunc {
	if simpleErrFunc == nil {
		return nil
	}

	isSuccessCollectorFunc := func(errorCollection *errwrappers.Collection) (isSuccess bool) {
		err := simpleErrFunc()

		if err != nil {
			errorCollection.AddTypeError(errorType, err)
		}

		return err == nil
	}

	return isSuccessCollectorFunc
}
