package errstr

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

func (it *newResultCreator) String(
	input string,
) *Result {
	return &Result{
		Value: input,
	}
}

func (it *newResultCreator) Item(
	item string,
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
	result string,
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		Value:        result,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) ValueOnly(
	result string,
) *Result {
	return &Result{
		Value:        result,
		ErrorWrapper: nil,
	}
}
