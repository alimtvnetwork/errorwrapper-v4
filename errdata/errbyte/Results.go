package errbyte

import (
	"errors"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errstr"
)

type Results struct {
	Values       []byte
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

	it.Values = []byte{}
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

func (it *Results) SafeValues() []byte {
	if it == nil {
		return []byte{}
	}

	if it.Values == nil {
		it.Values = []byte{}
	}

	return it.Values
}

func (it *Results) SafeValuesPtr() *[]byte {
	values := it.SafeValues()

	return &values
}

func (it *Results) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it *Results) ConvertToJsonResult() *corejson.Result {
	if it == nil {
		return corejson.
			Empty.
			ResultPtrWithErr(
				coredynamic.TypeName(Result{}),
				errors.New("errbyte results object is nil"))
	}

	return corejson.NewResult.Ptr(
		it.SafeValues(),
		it.ErrorWrapper.CompiledError(),
		"errbyte.Results")
}

func (it *Results) ValidValue() *corestr.ValidValue {
	if it == nil {
		return corestr.InvalidValidValueNoMessage()
	}

	return &corestr.ValidValue{
		Value:   it.SafeString(),
		IsValid: it.IsSuccess(),
		Message: it.ErrorWrapper.FullString(),
	}
}

func (it *Results) SimpleStringOnce(
	isInit bool,
) corestr.SimpleStringOnce {
	if it.IsAnyNull() {
		return corestr.Empty.SimpleStringOnce()
	}

	return corestr.New.SimpleStringOnce.Create(
		it.SafeString(), isInit)
}

func (it *Results) SplitLines() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return strings.Split(
		it.SafeString(),
		constants.NewLineUnix)
}

func (it *Results) IsEqual(term string) bool {
	if it == nil {
		return false
	}

	return it.SafeString() == term
}

func (it *Results) IsEqualIgnoreCase(term string) bool {
	if it == nil {
		return false
	}

	return strings.EqualFold(it.SafeString(), term)
}

func (it *Results) SplitLinesSimpleSlice() *corestr.SimpleSlice {
	return corestr.New.SimpleSlice.Direct(
		false,
		it.SplitLines())
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
	return string(it.SafeValues())
}

// NonEmptyString converter will only be called has bytes and no issues
func (it *Results) NonEmptyString(converter func(safeValues []byte) string) string {
	if it.HasIssuesOrEmpty() {
		return constants.EmptyString
	}

	return converter(it.SafeValues())
}

func (it *Results) ErrStr() *errstr.Result {
	if it == nil {
		return &errstr.Result{
			Value:        constants.EmptyString,
			ErrorWrapper: nil,
		}
	}

	return &errstr.Result{
		Value:        it.String(),
		ErrorWrapper: it.ErrorWrapper,
	}
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
