package errjson

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type newResultsCollectionCreator struct{}

func (it *newResultsCollectionCreator) Empty() *ResultsCollection {
	return &ResultsCollection{}
}

func (it *newResultsCollectionCreator) Error(
	errType errtype.Variation,
	err error,
) *ResultsCollection {
	return &ResultsCollection{
		ErrorCollection: errwrappers.NewWithErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newResultsCollectionCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *ResultsCollection {
	return &ResultsCollection{
		ErrorCollection: errwrappers.NewCap1().AddWrapperPtr(errorWrapper),
	}
}

func (it *newResultsCollectionCreator) ErrorCollection(
	errCollection *errwrappers.Collection,
) *ResultsCollection {
	return &ResultsCollection{
		ErrorCollection: errCollection,
	}
}

func (it *newResultsCollectionCreator) ResultsWithErrorCollection(
	jsonResults *corejson.ResultsCollection,
	errCollection *errwrappers.Collection,
) *ResultsCollection {
	return &ResultsCollection{
		ResultsCollection: jsonResults,
		ErrorCollection:   errCollection,
	}
}

func (it *newResultsCollectionCreator) Create(
	jsonResults *corejson.ResultsCollection,
	errWrapper *errorwrapper.Wrapper,
) *ResultsCollection {
	if jsonResults.IsEmpty() {
		return it.ErrorWrapper(errWrapper)
	}

	return &ResultsCollection{
		ResultsCollection: jsonResults,
		ErrorCollection:   errwrappers.NewCap1().AddWrapperPtr(errWrapper),
	}
}

func (it *newResultsCollectionCreator) RawValuesOnly(
	jsonResults ...corejson.Result,
) *ResultsCollection {
	if len(jsonResults) == 0 {
		return it.Empty()
	}

	return &ResultsCollection{
		ResultsCollection: corejson.
			NewResultsCollection.UsingResults(
			jsonResults...),
	}
}

func (it *newResultsCollectionCreator) AnyItems(
	anyItems ...interface{},
) *ResultsCollection {
	if len(anyItems) == 0 {
		return it.Empty()
	}

	collection := corejson.NewResultsCollection.AnyItems(
		anyItems...)

	sliceErr, hasErr := collection.AllErrors()

	if hasErr {
		compiledErr := errorwrapper.NewUsingError(
			codestack.Skip1,
			errtype.Marshalling,
			errcore.ManyErrorToSingle(sliceErr))

		return &ResultsCollection{
			ResultsCollection: collection,
			ErrorCollection:   errwrappers.NewCap1().AddWrapperPtr(compiledErr),
		}
	}

	return &ResultsCollection{
		ResultsCollection: collection,
	}
}

func (it *newResultsCollectionCreator) ValueOnly(
	jsonCollection *corejson.ResultsCollection,
) *ResultsCollection {
	if jsonCollection == nil {
		return it.Empty()
	}

	allErrors, _ := jsonCollection.AllErrors()

	return &ResultsCollection{
		ResultsCollection: jsonCollection,
		ErrorCollection:   errwrappers.NewUsingErrors(allErrors...),
	}
}

func (it *newResultsCollectionCreator) UsingError(
	err error,
) *ResultsCollection {
	if err == nil {
		return it.Empty()
	}

	return &ResultsCollection{
		ResultsCollection: nil,
		ErrorCollection:   errwrappers.NewWithOnlyError(err),
	}
}

func (it *newResultsCollectionCreator) UsingTypeError(
	errType errtype.Variation,
	err error,
) *ResultsCollection {
	if err == nil {
		return it.Empty()
	}

	return &ResultsCollection{
		ResultsCollection: nil,
		ErrorCollection:   errwrappers.NewCap1().AddTypeError(errType, err),
	}
}

func (it *newResultsCollectionCreator) UsingTypeMsg(
	errType errtype.Variation,
	msg string,
) *ResultsCollection {
	return &ResultsCollection{
		ResultsCollection: nil,
		ErrorCollection:   errwrappers.NewCap1().AddUsingMessages(errType, msg),
	}
}
