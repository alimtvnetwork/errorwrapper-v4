package errconv

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errcasted"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v4/internal/reflectinternal"
)

func GetPtr(wrapper interface{}) *errcasted.ResultPtr {
	if isany.Null(wrapper) {
		return errcasted.EmptyPtr()
	}

	pointerInfo := reflectinternal.GetPointerInfo(wrapper)

	if pointerInfo.IsPointer {
		actualWrapperPtr, isSuccess := wrapper.(*errorwrapper.Wrapper)

		if isSuccess {
			if actualWrapperPtr != nil {
				return errcasted.NewPtr(actualWrapperPtr)
			}
			// typed-nil *errorwrapper.Wrapper → empty, do not fall through.
			return errcasted.EmptyPtr()
		}
		// non-Wrapper pointer (e.g. *errwrappers.Collection): fall through
		// to the interface switch so it can match BasicErrWrapper /
		// BaseErrorOrCollectionWrapper.
	}

	actualWrapper, isSuccess := wrapper.(errorwrapper.Wrapper)

	if isSuccess {
		return errcasted.NewPtr(&actualWrapper)
	}

	switch castedNew := wrapper.(type) {
	case *errwrappers.Collection:
		return errcasted.NewPtr(castedNew.GetAsErrorWrapperPtr())
	case errwrappers.Collection:
		return errcasted.NewPtr(castedNew.GetAsErrorWrapperPtr())
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
