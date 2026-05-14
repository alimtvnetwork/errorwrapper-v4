package errany

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type newResultCreator struct{}

func (it *newResultCreator) Empty() *Result {
	return &Result{}
}

func (it *newResultCreator) Item(
	item float32,
) *Result {
	return &Result{
		Value: item,
	}
}

func (it *newResultCreator) Float(
	item float32,
) *Result {
	return &Result{
		Value: item,
	}
}

func (it *newResultCreator) Error(
	errType errtype.Variation,
	err error,
) *Result {
	return &Result{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newResultCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) Create(
	result float32,
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		Value:        result,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) ValueOnly(
	result float32,
) *Result {
	return &Result{
		Value: result,
	}
}
