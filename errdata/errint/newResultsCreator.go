package errint

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newResultsCreator struct{}

func (it *newResultsCreator) Empty() *Results {
	return &Results{}
}

func (it *newResultsCreator) Error(
	errType errtype.Variation,
	err error,
) *Results {
	return &Results{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newResultsCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Results {
	return &Results{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultsCreator) Create(
	errWrapper *errorwrapper.Wrapper,
	values []int,
) *Results {
	if len(values) == 0 {
		return it.ErrorWrapper(errWrapper)
	}

	return &Results{
		Values:       values,
		ErrorWrapper: errWrapper,
	}
}

func (it *newResultsCreator) SpreadCreate(
	errWrapper *errorwrapper.Wrapper,
	values ...int,
) *Results {
	if len(values) == 0 {
		return it.ErrorWrapper(errWrapper)
	}

	return &Results{
		Values:       values,
		ErrorWrapper: errWrapper,
	}
}

func (it *newResultsCreator) ValuesOnly(
	values []int,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}

func (it *newResultsCreator) SpreadValuesOnly(
	values ...int,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}

func (it *newResultsCreator) Items(
	values []int,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}
