package errstr

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type newResultTwoCreator struct{}

func (it *newResultTwoCreator) Empty() *Result2 {
	return &Result2{}
}

func (it *newResultTwoCreator) String(
	input string,
) *Result2 {
	return &Result2{
		Result: Result{
			Value: input,
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
	result1, result2 string,
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
	result2 string,
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
