package errfloat64

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
	values []float64,
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
	values ...float64,
) *Results {
	if len(values) == 0 {
		return it.ErrorWrapper(errWrapper)
	}

	return &Results{
		Values:       values,
		ErrorWrapper: errWrapper,
	}
}

func (it *newResultsCreator) Float64s(
	values []float64,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}

func (it *newResultsCreator) ValuesOnly(
	values []float64,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}

func (it *newResultsCreator) SpreadValuesOnly(
	values ...float64,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}

func (it *newResultsCreator) Items(
	values []float64,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}
