package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

func ConvertErrorFuncToWrapper(
	errorType errtype.Variation,
	simpleErrFunc SimpleErrorFunc,
) WrapperFunc {
	if simpleErrFunc == nil {
		return nil
	}

	wrapperFunc := func() *errorwrapper.Wrapper {
		err := simpleErrFunc()

		if err == nil {
			return nil
		}

		return errnew.Type.Error(errorType, err)
	}

	return wrapperFunc
}
