// Frozen: prefer erranygen.Result[T] for new code; see docs/extensibility.md §6.3.
package errany

import (
	"fmt"
	"strconv"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"
)

type Result struct {
	Value        interface{}
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Result) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.Value = nil
	it.ErrorWrapper.Dispose()
}

func (it *Result) IsEmpty() bool {
	return it == nil || it.Value == nil || it.Value == 0
}

func (it *Result) IsValid() bool {
	return !it.HasError()
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

func (it *Result) String() string {
	return converters.AnyTo.ToValueString(
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

func (it *Result) Str() string {
	if it == nil {
		return constants.EmptyString
	}

	if value, ok := it.Value.(string); ok {
		return value
	}

	return constants.EmptyString
}

func (it *Result) Dynamic() *coredynamic.Dynamic {
	if it == nil {
		return coredynamic.InvalidDynamicPtr()
	}

	return coredynamic.NewDynamicPtr(
		it.Value,
		it.IsValid())
}

func (it *Result) Bool() bool {
	if it == nil {
		return false
	}

	value, ok := it.Value.(bool)
	return ok && value
}

func (it *Result) Int() int {
	if it == nil {
		return 0
	}

	switch value := it.Value.(type) {
	case int:
		return value
	case byte:
		return int(value)
	case float32:
		return int(value)
	case float64:
		return int(value)
	}

	return 0
}

func (it *Result) Byte() byte {
	if it == nil {
		return 0
	}

	return it.Value.(byte)
}

func (it *Result) Float32() float32 {
	if it == nil {
		return 0
	}

	return it.Value.(float32)
}

func (it *Result) Float64() float64 {
	if it == nil {
		return 0
	}

	return it.Value.(float64)
}

func (it *Result) ErrorWrapperInf() errorwrapper.ErrWrapper {
	return it.ErrorWrapper
}

func (it *Result) SafeString() string {
	if it == nil {
		return ""
	}

	return it.String()
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

func (it *Result) IsAnyNull() bool {
	return it == nil
}
