package errjson

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type emptyCreator struct{}

func (it *emptyCreator) Result() *Result {
	return &Result{}
}

func (it *emptyCreator) ResultsCollection() *ResultsCollection {
	return &ResultsCollection{}
}

func (it *emptyCreator) ResultWithError(
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		ErrorWrapper: errorWrapper,
	}
}

func (it *emptyCreator) ResultsCollectionWithError(
	errorWrapper *errorwrapper.Wrapper,
) *ResultsCollection {
	return &ResultsCollection{
		ErrorCollection: errwrappers.NewUsingErrorWrappers(errorWrapper),
	}
}

func (it *emptyCreator) ResultsCollectionWithErrorCollection(
	errCollection *errwrappers.Collection,
) *ResultsCollection {
	return &ResultsCollection{
		ErrorCollection: errCollection,
	}
}

func (it *emptyCreator) ResultsCollectionWithValues(
	anyItems ...interface{},
) *ResultsCollection {
	return &ResultsCollection{
		ResultsCollection: corejson.
			NewResultsCollection.
			AnyItems(anyItems...),
	}
}

func (it *emptyCreator) ResultWithValue(jsonResult *corejson.Result) *Result {
	return &Result{
		Result: jsonResult,
		ErrorWrapper: errnew.
			Error.
			Type(
				errtype.Marshalling,
				jsonResult.MeaningfulError()),
	}
}
