package errstr

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/anycmp"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coreutils/stringutil"
	"github.com/alimtvnetwork/errorwrapper-v3"
)

type Result struct {
	Value        string
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Result) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Result) IsAnyNull() bool {
	return it == nil
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.Value = constants.EmptyString
	it.ErrorWrapper.Dispose()
}

func (it *Result) IsEmpty() bool {
	return it == nil || it.Value == constants.EmptyString
}

func (it *Result) IsValid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Result) IsSuccess() bool {
	return it.HasSafeItems()
}

func (it *Result) IsFailed() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Result) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it Result) String() string {
	return it.Value
}

func (it Result) Bytes() []byte {
	return []byte(it.Value)
}

func (it Result) SafeBytes() []byte {
	if it.HasIssuesOrEmpty() {
		return []byte{}
	}

	return []byte(it.Value)
}

func (it *Result) IsEmptyOrWhitespace() bool {
	return it.IsEmpty() || stringutil.IsEmptyOrWhitespace(it.Value)
}

func (it *Result) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *Result) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *Result) IsEmptyError() bool {
	return it == nil || it.ErrorWrapper.IsEmpty()
}

func (it *Result) HasIssuesOrWhitespace() bool {
	return it.HasIssuesOrEmpty() ||
		stringutil.IsEmptyOrWhitespace(it.Value)
}

func (it *Result) Int() int {
	if it.HasIssuesOrEmpty() {
		return constants.Zero
	}

	v, _ := converters.StringToIntegerWithDefault(
		it.Value,
		constants.Zero)

	return v
}

func (it *Result) Byte() byte {
	return byte(it.Int())
}

func (it *Result) Bool() bool {
	if it == nil {
		return false
	}

	return stringutil.ToBool(it.Value)
}

func (it *Result) ValidValue() *corestr.ValidValue {
	if it == nil {
		return corestr.InvalidValidValueNoMessage()
	}

	return &corestr.ValidValue{
		Value:   it.Value,
		IsValid: it.IsSuccess(),
		Message: it.ErrorWrapper.FullString(),
	}
}

func (it *Result) SimpleStringOnce(
	isInit bool,
) corestr.SimpleStringOnce {
	if it == nil {
		return corestr.Empty.SimpleStringOnce()
	}

	return corestr.New.SimpleStringOnce.Create(
		it.Value, isInit)
}

func (it *Result) SimpleStringOnceInit() corestr.SimpleStringOnce {
	if it == nil {
		return corestr.Empty.SimpleStringOnce()
	}

	return corestr.New.SimpleStringOnce.Create(
		it.Value, true)
}

func (it *Result) SplitLines() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return strings.Split(
		it.Value,
		constants.NewLineUnix)
}

func (it *Result) IsEqualResult(right *Result) bool {
	cmp := anycmp.Cmp(it, right)

	if cmp.IsDefinedProperly() {
		return cmp.IsEqual()
	}

	if it.HasError() != right.HasError() {
		return false
	}

	if it.ErrorWrapper.IsNotEquals(right.ErrorWrapper) {
		return false
	}

	if it.Value != right.Value {
		return false
	}

	return true
}

func (it *Result) IsEqual(term string) bool {
	if it == nil {
		return false
	}

	return it.Value == term
}

func (it *Result) IsEqualIgnoreCase(term string) bool {
	if it == nil {
		return false
	}

	return strings.EqualFold(it.Value, term)
}

func (it *Result) SplitLinesSimpleSlice() *corestr.SimpleSlice {
	return corestr.New.SimpleSlice.Direct(
		false,
		it.SplitLines())
}

func (it *Result) ErrorWrapperInf() errorwrapper.ErrWrapper {
	return it.ErrorWrapper
}

func (it *Result) SafeString() string {
	if it == nil {
		return ""
	}

	return it.Value
}

func (it Result) Json() corejson.Result {
	return corejson.New(it)
}

func (it Result) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Result) JsonModelAny() interface{} {
	return it
}

func (it *Result) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Unmarshal(it)
}

func (it *Result) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Result) AsValueWithErrorWrapperBinder() errorwrapper.ValueWithErrorWrapperBinder {
	return it
}
