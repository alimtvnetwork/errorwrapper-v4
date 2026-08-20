// Frozen: prefer erranygen.Result[T] for new code; see docs/extensibility.md §6.3.
package errbyte

import (
	"strconv"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

type Result struct {
	Value        byte
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

func (it *Result) IsMax() bool {
	if it.IsAnyNull() {
		return false
	}

	return it.Value >= constants.MaxUnit8
}

func (it Result) String() string {
	return string(it.Value)
}

func (it *Result) NumberString() string {
	if it == nil {
		return "0"
	}

	return strconv.Itoa(int(it.Value))
}

func (it *Result) NumberStringSafe() string {
	return it.NumberString()
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

func (it *Result) IsValidRange(min, max byte) bool {
	if it == nil {
		return false
	}

	return it.Value >= min && it.Value <= max
}

func (it *Result) IsSafeValidRange(min, max byte) bool {
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

func (it *Result) Float32() float32 {
	if it == nil {
		return 0
	}

	return float32(it.Value)
}

func (it *Result) Float64() float64 {
	if it == nil {
		return 0
	}

	return float64(it.Value)
}

// Bool anything but zero or null will be true
func (it *Result) Bool() bool {
	if it == nil || it.Value == 0 {
		return false
	}

	return true
}

func (it *Result) ErrorWrapperInf() errorwrapper.ErrWrapper {
	if it == nil {
		return nil
	}

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
