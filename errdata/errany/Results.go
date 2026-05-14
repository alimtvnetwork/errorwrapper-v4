package errany

import (
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"
)

type Results struct {
	Values       []interface{}
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Results) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Results) Clear() {
	if it == nil {
		return
	}

	it.Values = []interface{}{}
}

func (it *Results) IsAnyNull() bool {
	return it == nil || it.Values == nil
}

func (it *Results) Dispose() {
	if it == nil {
		return
	}

	it.Values = nil
	it.ErrorWrapper.Dispose()
}

func (it *Results) Length() int {
	if it == nil || it.Values == nil {
		return 0
	}

	return len(it.Values)
}

func (it *Results) HasAnyItem() bool {
	return it.Length() > 0
}

// HasSafeItems No errors and has items
func (it *Results) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *Results) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Results) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *Results) SafeValues() []interface{} {
	return *it.SafeValuesPtr()
}

func (it *Results) SafeValuesPtr() *[]interface{} {
	if it.IsEmpty() {
		return &[]interface{}{}
	}

	return &it.Values
}

func (it *Results) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it *Results) IsEmptyError() bool {
	return it == nil || it.ErrorWrapper.IsEmpty()
}

func (it *Results) IsValid() bool {
	return it.HasSafeItems()
}

func (it *Results) IsSuccess() bool {
	return it.HasSafeItems()
}

func (it *Results) IsFailed() bool {
	return it.HasIssuesOrEmpty()
}

func (it Results) String() string {
	return converters.AnyToValueString(it.SafeValues())
}

func (it *Results) ErrorWrapperInf() errorwrapper.ErrWrapper {
	return it.ErrorWrapper
}

func (it *Results) SafeString() string {
	if it == nil {
		return ""
	}

	return it.String()
}

func (it Results) Json() corejson.Result {
	return corejson.New(it)
}

func (it Results) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Results) JsonModelAny() interface{} {
	length := it.Length()

	if length == 0 {
		return corejson.Empty.ResultsCollection()
	}

	jsonCollection := corejson.NewResultsCollection.UsingCap(it.Length())

	return jsonCollection.AddAnyItems(it.Values...)
}

func (it *Results) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Unmarshal(it)
}

func (it *Results) AsResultsContractsBinder() errorwrapper.ResultsContractsBinder {
	return it
}

func (it *Results) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Results) AsValueWithErrorWrapperBinder() errorwrapper.ValueWithErrorWrapperBinder {
	return it
}

func (it *Results) AsResultsWithErrorWrapper() errorwrapper.ResultsContractsBinder {
	return it
}
