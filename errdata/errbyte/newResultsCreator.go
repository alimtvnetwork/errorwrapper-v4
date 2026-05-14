package errbyte

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
	values []byte,
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
	values ...byte,
) *Results {
	if len(values) == 0 {
		return it.ErrorWrapper(errWrapper)
	}

	return &Results{
		Values:       values,
		ErrorWrapper: errWrapper,
	}
}

func (it *newResultsCreator) String(
	valueString string,
) *Results {
	if len(valueString) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: []byte(valueString),
	}
}

func (it *newResultsCreator) StringWithErrorWrapper(
	valueString string,
	errorWrapper *errorwrapper.Wrapper,
) *Results {
	if len(valueString) == 0 {
		return it.Empty()
	}

	return &Results{
		Values:       []byte(valueString),
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultsCreator) ValuesOnly(
	values []byte,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}

func (it *newResultsCreator) SpreadValuesOnly(
	values ...byte,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}

func (it *newResultsCreator) Items(
	values []byte,
) *Results {
	if len(values) == 0 {
		return it.Empty()
	}

	return &Results{
		Values: values,
	}
}
