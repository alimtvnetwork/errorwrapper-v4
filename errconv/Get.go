package errconv

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errcasted"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/internal/reflectinternal"
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
		// fall through to the interface switch below so non-Wrapper
		// pointers (e.g. *errwrappers.Collection) can still match
		// BasicErrWrapper / BaseErrorOrCollectionWrapper.
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
