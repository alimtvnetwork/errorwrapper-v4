package errint

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"

	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type ResultsWithErrorCollection struct {
	Values        []int
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

	it.Values = []int{}
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
	items, err := coredynamic.SliceItemsAsStringsAny(
		it.SafeValues())

	if err != nil {
		panic(err)
	}

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

func (it *ResultsWithErrorCollection) SafeValues() []int {
	return *it.SafeValuesPtr()
}

func (it *ResultsWithErrorCollection) SafeValuesPtr() *[]int {
	if it.HasIssuesOrEmpty() {
		return &[]int{}
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

func NewResultsWithErrorCollection(
	values []int,
	errCollection *errwrappers.Collection,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values:        values,
		ErrorWrappers: errCollection,
	}
}

func NewResultsWithErrorCollectionUsingErrorCollection(
	errCollection *errwrappers.Collection,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values:        []int{},
		ErrorWrappers: errCollection,
	}
}

func EmptyResultsWithErrorCollection() *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values:        []int{},
		ErrorWrappers: errwrappers.Empty(),
	}
}

func NewResultsWithErrorCollectionUsingTypeMessage(
	errType errtype.Variation,
	msg string,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values: []int{},
		ErrorWrappers: errwrappers.NewWithMessage(
			errType,
			msg),
	}
}

func NewResultsWithErrorCollectionUsingTypeError(
	errType errtype.Variation,
	err error,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values: []int{},
		ErrorWrappers: errwrappers.NewWithError(
			constants.ArbitraryCapacity2,
			errType,
			err),
	}
}

func NewResultsWithErrorCollectionUsingType(
	errType errtype.Variation,
) *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{
		Values: []int{},
		ErrorWrappers: errwrappers.NewWithType(
			errType),
	}
}
