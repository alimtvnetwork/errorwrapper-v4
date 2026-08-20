package errstr

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

type newResultsWithErrorCollectionCreator struct{}

func (it *newResultsWithErrorCollectionCreator) Empty() *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{}
}

func (it *newResultsWithErrorCollectionCreator) Error(
	errType errtype.Variation,
	err error,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		ErrorWrappers: errwrappers.
			NewCap1().
			AddWrapperPtr(
				errnew.Type.ErrorUsingStackSkip(
					codestack.Skip1,
					errType,
					err)),
	}
}

func (it *newResultsWithErrorCollectionCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		ErrorWrappers: errwrappers.
			NewCap1().
			AddWrapperPtr(errorWrapper),
	}
}

func (it *newResultsWithErrorCollectionCreator) Create(
	values []string,
	errorCollection *errwrappers.Collection,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values:        values,
		ErrorWrappers: errorCollection,
	}
}

func (it *newResultsWithErrorCollectionCreator) SpreadCreate(
	errorCollection *errwrappers.Collection,
	values ...string,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values:        values,
		ErrorWrappers: errorCollection,
	}
}

func (it *newResultsWithErrorCollectionCreator) ErrorWrapperWithValues(
	errWrapper *errorwrapper.Wrapper,
	values ...string,
) *ResultsWithErrorCollection {
	if len(values) == 0 {
		return it.ErrorWrapper(errWrapper)
	}

	return &ResultsWithErrorCollection{
		Values: values,
		ErrorWrappers: errwrappers.
			NewCap1().
			AddWrapperPtr(errWrapper),
	}
}

func (it *newResultsWithErrorCollectionCreator) SpreadItems(
	values ...string,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values: values,
	}
}

func (it *newResultsWithErrorCollectionCreator) ValuesOnly(
	values []string,
) *ResultsWithErrorCollection {
	if len(values) == 0 {
		return it.Empty()
	}

	return &ResultsWithErrorCollection{
		Values: values,
	}
}

func (it *newResultsWithErrorCollectionCreator) Strings(
	values []string,
) *ResultsWithErrorCollection {
	if len(values) == 0 {
		return it.Empty()
	}

	return &ResultsWithErrorCollection{
		Values: values,
	}
}

func (it *newResultsWithErrorCollectionCreator) ResultsWithErrorCollection(
	values []string,
	errCollection *errwrappers.Collection,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values:        values,
		ErrorWrappers: errCollection,
	}
}

func (it *newResultsWithErrorCollectionCreator) ErrorCollection(
	errCollection *errwrappers.Collection,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values:        []string{},
		ErrorWrappers: errCollection,
	}
}

func (it *newResultsWithErrorCollectionCreator) ErrorType(
	errType errtype.Variation,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values: []string{},
		ErrorWrappers: errwrappers.NewWithTypeUsingStackSkip(
			codestack.Skip1,
			errType),
	}
}

func (it *newResultsWithErrorCollectionCreator) ErrorWrapperUsingTypeMsg(
	errType errtype.Variation,
	msg string,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values: []string{},
		ErrorWrappers: errwrappers.NewWithMessageUsingStackSkip(
			codestack.Skip1,
			errType,
			msg),
	}
}

func (it *newResultsWithErrorCollectionCreator) TypeWithError(
	errType errtype.Variation,
	err error,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values: []string{},
		ErrorWrappers: errwrappers.NewWithError(
			constants.ArbitraryCapacity2,
			errType,
			err),
	}
}
