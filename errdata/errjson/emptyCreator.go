package errjson

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
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
