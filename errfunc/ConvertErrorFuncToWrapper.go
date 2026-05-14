package errfunc

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
