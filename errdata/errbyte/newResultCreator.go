package errbyte

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newResultCreator struct{}

func (it *newResultCreator) Empty() *Result {
	return &Result{}
}

func (it *newResultCreator) Item(
	item byte,
) *Result {
	return &Result{
		Value: item,
	}
}

func (it *newResultCreator) Byte(
	item byte,
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
	result byte,
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		Value:        result,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) ValueOnly(
	result byte,
) *Result {
	return &Result{
		Value: result,
	}
}
