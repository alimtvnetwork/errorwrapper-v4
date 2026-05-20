package errfloat

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newResultTwoCreator struct{}

func (it *newResultTwoCreator) Empty() *Result2 {
	return &Result2{}
}

func (it *newResultTwoCreator) Item(
	input float32,
) *Result2 {
	return &Result2{
		Result: Result{
			Value: normalizeFloat(input),
		},
	}
}

func (it *newResultTwoCreator) Error(
	errType errtype.Variation,
	err error,
) *Result2 {
	return &Result2{
		Result: Result{
			ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
				codestack.Skip1,
				errType,
				err),
		},
	}
}

func (it *newResultTwoCreator) ValueOnly(
	result1, result2 float32,
) *Result2 {
	return &Result2{
		Result: Result{
			Value: result1,
		},
		Value2: result2,
	}
}

func (it *newResultTwoCreator) Create(
	result,
	result2 float32,
	wrapper *errorwrapper.Wrapper,
) *Result2 {
	return &Result2{
		Result: Result{
			Value:        result,
			ErrorWrapper: wrapper,
		},
		Value2: result2,
	}
}

func (it *newResultTwoCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Result2 {
	return &Result2{
		Result: Result{
			ErrorWrapper: errorWrapper,
		},
	}
}
