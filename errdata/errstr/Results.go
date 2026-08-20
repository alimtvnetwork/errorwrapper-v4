package errstr

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/anycmp"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/corecmp"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

type Results struct {
	Values       []string
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Results) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Results) IsAnyNull() bool {
	return it == nil || it.Values == nil
}

func (it *Results) Clear() {
	if it == nil {
		return
	}

	it.Values = []string{}
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

func (it *Results) SafeValues() []string {
	if it == nil {
		return []string{}
	}

	if it.Values == nil {
		it.Values = []string{}
	}

	return it.Values
}

func (it *Results) SafeValuesPtr() *[]string {
	values := it.SafeValues()

	return &values
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

func (it *Results) List() []string {
	return it.Strings()
}

func (it *Results) Lines() []string {
	return it.Strings()
}

func (it *Results) Items() []string {
	return it.Strings()
}

func (it *Results) Strings() []string {
	if it.IsAnyNull() {
		return []string{}
	}

	return it.Values
}

func (it *Results) String() string {
	return strings.Join(
		it.SafeValues(),
		constants.CommaUnixNewLine)
}

func (it *Results) IsEqualResultDefault(
	right *Results,
) bool {
	return it.IsEqualResult(
		false,
		false,
		right)
}

func (it *Results) IsDistinctEqualResult(
	right *Results,
) bool {
	return it.IsEqualResult(
		true,
		false,
		right)
}

func (it *Results) IsIgnoreOrderEqualResult(
	right *Results,
) bool {
	return it.IsEqualResult(
		false,
		true,
		right)
}

func (it *Results) IsEqualResult(
	isDistinctCompare,
	isIgnoreOrder bool,
	right *Results,
) bool {
	cmp := anycmp.Cmp(it, right)

	if cmp.IsDefinedProperly() {
		return cmp.IsEqual()
	}

	if it.Length() != right.Length() {
		return false
	}

	if it.HasError() != right.HasError() {
		return false
	}

	if it.ErrorWrapper.IsNotEquals(right.ErrorWrapper) {
		return false
	}

	if isIgnoreOrder {
		return corecmp.IsStringsEqualWithoutOrder(it.Values, right.Values)
	}

	if isDistinctCompare {
		return it.Hashset().HasAll(right.Values...)
	}

	return corecmp.IsStringsEqual(it.Values, right.Values)
}

func (it *Results) SimpleSlice() *corestr.SimpleSlice {
	return corestr.New.SimpleSlice.Direct(
		false,
		it.SafeValues())
}

func (it *Results) Hashset() *corestr.Hashset {
	if it.IsAnyNull() {
		return corestr.New.Hashset.Strings(nil)
	}

	return corestr.New.Hashset.Strings(
		it.Values)
}

func (it *Results) UniqueMap() map[string]bool {
	return converters.StringsTo.Hashset(it.SafeValues())
}

func (it *Results) UniqueMapPtr() *map[string]bool {
	itemsMap := it.UniqueMap()

	return &itemsMap
}

func (it *Results) StringCollection() *corestr.Collection {
	if it.IsAnyNull() {
		return corestr.Empty.Collection()
	}

	return corestr.New.Collection.Strings(
		it.Values,
	)
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
	return it
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
