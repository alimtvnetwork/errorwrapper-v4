package errbool

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"

	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

type ResultsWithErrorCollection struct {
	Values        []bool
	ErrorWrappers *errwrappers.Collection
}

func (it *ResultsWithErrorCollection) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *ResultsWithErrorCollection) IsAnyNull() bool {
	return it == nil || it.Values == nil
}

func (it *ResultsWithErrorCollection) Clear() {
	if it == nil {
		return
	}

	it.Values = []bool{}
}

func (it *ResultsWithErrorCollection) Dispose() {
	if it == nil {
		return
	}

	it.Values = nil
	it.ErrorWrappers.Dispose()
}

func (it *ResultsWithErrorCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *ResultsWithErrorCollection) IsValid() bool {
	return it.HasSafeItems()
}

func (it *ResultsWithErrorCollection) IsSuccess() bool {
	return it.HasSafeItems()
}

func (it *ResultsWithErrorCollection) IsFailed() bool {
	return it.HasIssuesOrEmpty()
}

func (it ResultsWithErrorCollection) String() string {
	items := converters.AnyTo.ToStrings(
		false, it.SafeValues())

	return strings.Join(
		items,
		constants.CommaUnixNewLine)
}

func (it *ResultsWithErrorCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *ResultsWithErrorCollection) Length() int {
	if it == nil || it.Values == nil {
		return 0
	}

	return len(it.Values)
}

// HasSafeItems No errors and has items
func (it *ResultsWithErrorCollection) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *ResultsWithErrorCollection) IsEmptyItems() bool {
	return it.Length() == 0
}

func (it *ResultsWithErrorCollection) IsEmptyError() bool {
	return it == nil || it.ErrorWrappers.IsEmpty()
}

func (it *ResultsWithErrorCollection) HasError() bool {
	return it != nil && it.ErrorWrappers.HasError()
}

func (it *ResultsWithErrorCollection) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it *ResultsWithErrorCollection) SafeValues() []bool {
	return *it.SafeValuesPtr()
}

func (it *ResultsWithErrorCollection) SafeValuesPtr() *[]bool {
	if it.HasIssuesOrEmpty() {
		return &[]bool{}
	}

	return &it.Values
}

func (it *ResultsWithErrorCollection) ErrorWrapperInf() errorwrapper.ErrWrapper {
	return it.ErrorWrappers.GetAsErrorWrapperPtr()
}

func (it *ResultsWithErrorCollection) SafeString() string {
	if it == nil {
		return ""
	}

	return it.String()
}

func (it ResultsWithErrorCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it ResultsWithErrorCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *ResultsWithErrorCollection) JsonModelAny() interface{} {
	return it
}

func (it *ResultsWithErrorCollection) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it *ResultsWithErrorCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it ResultsWithErrorCollection) AsValuesWithErrorWrapperCollectionBinder() errorwrapper.ValuesWithErrorWrapperCollectionBinder {
	return &it
}
