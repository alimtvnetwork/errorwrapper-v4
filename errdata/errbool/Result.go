package errbool

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

type Result struct {
	Value        bool
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.ErrorWrapper.Dispose()
}

func (it *Result) IsAnyNull() bool {
	return it == nil
}

// IsEmpty
//
// nil or false returns true for empty
func (it *Result) IsEmpty() bool {
	return it == nil || !it.Value
}

func (it *Result) IsValid() bool {
	return !it.HasIssuesOrEmpty()
}

// IsSuccess
//
// Indicates NOT null || value false || has error
func (it *Result) IsSuccess() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *Result) IsFailed() bool {
	return it.HasIssuesOrEmpty()
}

// HasIssuesOrEmpty
//
// Indicates null || value false || has error
func (it *Result) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it Result) String() string {
	return converters.AnyToValueString(
		it.Value)
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

func (it *Result) IsTrue() bool {
	return it != nil && it.Value
}

// IsFalse
//
// it == nil || !it.Value
func (it *Result) IsFalse() bool {
	return it == nil || !it.Value
}

func (it *Result) IsApplicable() bool {
	return it.IsTrue()
}

func (it *Result) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Result) Int() int {
	if it.IsValid() {
		return constants.One
	}

	return constants.Zero
}

func (it *Result) Byte() byte {
	if it.IsValid() {
		return constants.One
	}

	return constants.Zero
}

func (it *Result) ErrorWrapperInf() errorwrapper.ErrWrapper {
	return it.ErrorWrapper
}

func (it *Result) SafeString() string {
	if it == nil {
		return "false"
	}

	return strconv.FormatBool(
		it.Value)
}

func (it Result) Json() corejson.Result {
	return corejson.New(it)
}

func (it Result) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Result) JsonModelAny() interface{} {
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
