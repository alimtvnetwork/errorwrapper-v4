package errjson

import (
	"gitlab.com/evatix-go/core/coredata/corejson"

	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type ResultsCollection struct {
	*corejson.ResultsCollection
	ErrorCollection *errwrappers.Collection
}

func (it *ResultsCollection) IsAnyNull() bool {
	return it == nil || it.ResultsCollection == nil
}

func (it *ResultsCollection) HasError() bool {
	return it != nil && it.ErrorCollection.HasError()
}

func (it *ResultsCollection) IsEmpty() bool {
	return it.IsAnyNull() || it.ResultsCollection.IsEmpty()
}

func (it *ResultsCollection) Length() int {
	if it.IsAnyNull() {
		return 0
	}

	return it.ResultsCollection.Length()
}

func (it *ResultsCollection) Dispose() {
	if it == nil {
		return
	}

	it.ResultsCollection.Dispose()
	it.ErrorCollection.Dispose()
}
