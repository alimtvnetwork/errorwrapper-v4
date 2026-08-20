// Frozen: prefer erranygen.Result[T] for new code; see docs/extensibility.md §6.3.
package errfloat64

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/internal/consts"
)

type Result struct {
	Value        float64
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Result) IsInvalid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.Value = constants.Zero
	it.ErrorWrapper.Dispose()
}

func (it *Result) IsEmpty() bool {
	return it == nil || it.Value == 0
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

func (it *Result) IsValidRange(min, max float64) bool {
	if it == nil {
		return false
	}

	return it.Value >= min && it.Value <= max
}

func (it *Result) IsSafeValidRange(min, max float64) bool {
	if it == nil {
		return false
	}

	return it.IsEmptyError() &&
		it.Value >= min &&
		it.Value <= max
}

func (it *Result) Int() int {
	if it == nil {
		return 0
	}

	return int(it.Value)
}

func (it *Result) Byte() byte {
	if it == nil {
		return 0
	}

	if it.Value > consts.MaxUnit8AsFloat64 {
		return constants.MaxUnit8
	}

	if it.Value < 0 {
		return constants.Zero
	}

	return byte(it.Value)
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
