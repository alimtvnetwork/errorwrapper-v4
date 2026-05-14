package errconv

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/isany"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errcasted"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/internal/reflectinternal"
)

func GetPtr(wrapper interface{}) *errcasted.ResultPtr {
	if isany.Null(wrapper) {
		return errcasted.EmptyPtr()
	}

	pointerInfo := reflectinternal.GetPointerInfo(wrapper)

	if pointerInfo.IsPointer {
		actualWrapperPtr, isSuccess := wrapper.(*errorwrapper.Wrapper)

		if isSuccess && actualWrapperPtr != nil {
			return errcasted.NewPtr(actualWrapperPtr)
		}

		return errcasted.EmptyPtr()
	}

	actualWrapper, isSuccess := wrapper.(errorwrapper.Wrapper)

	if isSuccess {
		return errcasted.NewPtr(&actualWrapper)
	}

	switch castedNew := wrapper.(type) {
	case errcoreinf.BasicErrWrapper:
		return errcasted.NewPtr(
			errorwrapper.NewUsingBasicErr(castedNew))
	case errcoreinf.BaseErrorOrCollectionWrapper:
		return errcasted.NewPtr(
			errorwrapper.InterfaceToErrorWrapperUsingStackSkip(
				codestack.Skip1,
				errtype.Unknown,
				castedNew))
	}

	return errcasted.EmptyPtr()
}

func Get(wrapperIn interface{}) errcasted.Result {
	wrapperCastedPtr := GetPtr(wrapperIn)

	return wrapperCastedPtr.ToResult()
}
