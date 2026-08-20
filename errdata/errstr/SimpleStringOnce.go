package errstr

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coreutils/stringutil"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

type SimpleStringOnce struct {
	Value        corestr.SimpleStringOnce
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *SimpleStringOnce) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *SimpleStringOnce) Dispose() {
	if it == nil {
		return
	}

	it.Value.Dispose()
	it.ErrorWrapper.Dispose()
}

func (it *SimpleStringOnce) IsEmpty() bool {
	return it == nil || it.Value.IsUninitialized()
}

func (it *SimpleStringOnce) IsValid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *SimpleStringOnce) IsSuccess() bool {
	return it.HasSafeItems()
}

func (it *SimpleStringOnce) IsFailed() bool {
	return it.HasIssuesOrEmpty()
}

func (it *SimpleStringOnce) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it SimpleStringOnce) String() string {
	return it.Value.String()
}

func (it *SimpleStringOnce) IsEmptyOrWhitespace() bool {
	return it.IsEmpty() || it.Value.IsWhitespace()
}

func (it *SimpleStringOnce) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *SimpleStringOnce) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *SimpleStringOnce) IsEmptyError() bool {
	return it == nil || it.ErrorWrapper.IsEmpty()
}

func (it *SimpleStringOnce) HasIssuesOrWhitespace() bool {
	return it.HasIssuesOrEmpty() ||
		it.Value.IsWhitespace()
}

func (it *SimpleStringOnce) Int() int {
	if it.HasIssuesOrEmpty() {
		return constants.Zero
	}

	v, _ := converters.StringTo.IntegerWithDefault(
		it.Value.Value(),
		constants.Zero)

	return v
}

func (it *SimpleStringOnce) Byte() byte {
	return byte(it.Int())
}

func (it *SimpleStringOnce) Bool() bool {
	if it == nil {
		return false
	}

	return stringutil.ToBool(it.Value.String())
}

func (it *SimpleStringOnce) ValidValue() *corestr.ValidValue {
	if it == nil {
		return corestr.InvalidValidValueNoMessage()
	}

	return &corestr.ValidValue{
		Value:   it.Value.String(),
		IsValid: it.IsSuccess(),
		Message: it.ErrorWrapper.FullString(),
	}
}

func (it *SimpleStringOnce) SplitLines() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return strings.Split(
		it.Value.String(),
		constants.NewLineUnix)
}

func (it *SimpleStringOnce) IsEqual(term string) bool {
	if it == nil {
		return false
	}

	return it.Value.Is(term)
}

func (it *SimpleStringOnce) IsEqualIgnoreCase(term string) bool {
	if it == nil {
		return false
	}

	return it.Value.IsEqualNonSensitive(term)
}

func (it *SimpleStringOnce) SplitLinesSimpleSlice() *corestr.SimpleSlice {
	return corestr.New.SimpleSlice.Direct(
		false,
		it.SplitLines())
}

func (it *SimpleStringOnce) ErrorWrapperInf() errorwrapper.ErrWrapper {
	return it.ErrorWrapper
}

func (it *SimpleStringOnce) SafeString() string {
	if it == nil {
		return ""
	}

	return it.String()
}

func (it SimpleStringOnce) Json() corejson.Result {
	return corejson.New(it)
}

func (it SimpleStringOnce) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it SimpleStringOnce) JsonModelAny() interface{} {
	return it
}

func (it *SimpleStringOnce) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Unmarshal(it)
}

func (it *SimpleStringOnce) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *SimpleStringOnce) AsValueWithErrorWrapperBinder() errorwrapper.ValueWithErrorWrapperBinder {
	return it
}
