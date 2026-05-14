package errjson

import (
	"bytes"
	"encoding/json"
	"strings"

	"gitlab.com/evatix-go/core/anycmp"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type Result struct {
	*corejson.Result
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Result) IsAnyNull() bool {
	return it == nil || it.Result == nil
}

func (it *Result) Length() int {
	if it == nil || it.Result == nil {
		return 0
	}

	return it.Result.Length()
}

func (it *Result) IsNull() bool {
	return it == nil || it.Result == nil
}

func (it *Result) HasAnyItem() bool {
	return it != nil &&
		it.Result != nil &&
		it.Result.HasBytes()
}

// HasSafeItems No errors and has items
func (it *Result) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *Result) IsEmpty() bool {
	return it == nil ||
		it.Result == nil ||
		it.Result.IsEmptyJsonBytes()
}

func (it *Result) HasError() bool {
	return it != nil &&
		it.ErrorWrapper.HasError() ||
		it != nil &&
			it.Result.HasError()
}

func (it *Result) SafeValuesPtr() *[]byte {
	if it.IsAnyNull() {
		return &[]byte{}
	}

	return &it.Result.Bytes
}

func (it *Result) SafeValues() []byte {
	if it.IsAnyNull() {
		return []byte{}
	}

	return it.Result.Bytes
}

func (it *Result) SafeBytes() []byte {
	return it.SafeValues()
}

func (it *Result) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it *Result) ValidValue() *corestr.ValidValue {
	if it == nil {
		return corestr.InvalidValidValueNoMessage()
	}

	return &corestr.ValidValue{
		Value:   it.SafeString(),
		IsValid: it.IsSuccess(),
		Message: it.ErrorWrapper.FullString(),
	}
}

func (it *Result) SimpleStringOnce(
	isInit bool,
) corestr.SimpleStringOnce {
	if it.IsAnyNull() {
		return corestr.Empty.SimpleStringOnce()
	}

	return corestr.New.SimpleStringOnce.Create(
		it.SafeString(), isInit)
}

func (it *Result) SplitLines() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return strings.Split(
		it.SafeString(),
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

	if !it.Result.IsEqualPtr(right.Result) {
		return false
	}

	return true
}

func (it *Result) IsEqual(term string) bool {
	if it == nil {
		return false
	}

	return it.SafeString() == term
}

func (it *Result) IsEqualIgnoreCase(term string) bool {
	if it == nil {
		return false
	}

	return strings.EqualFold(it.SafeString(), term)
}

func (it *Result) SplitLinesSimpleSlice() *corestr.SimpleSlice {
	return corestr.New.SimpleSlice.Direct(
		false,
		it.SplitLines())
}

func (it Result) String() string {
	if it.IsAnyNull() {
		return constants.EmptyString
	}

	return it.Result.JsonString()
}

func (it Result) SafeString() string {
	if it.IsAnyNull() {
		return constants.EmptyString
	}

	return it.Result.JsonString()
}

func (it Result) JsonString() string {
	if it.IsAnyNull() {
		return constants.EmptyString
	}

	return it.Result.JsonString()
}

func (it Result) PrettyJsonString() string {
	if it.IsAnyNull() {
		return constants.EmptyString
	}

	return it.Result.PrettyJsonString()
}

func (it *Result) PrettyJsonStringOrErrString() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.Result.PrettyJsonStringOrErrString()
}

func (it *Result) PrettyJsonBuffer(prefix, indent string) (*bytes.Buffer, error) {
	if it.IsAnyNull() {
		return nil, errcore.
			CannotBeNilType.
			ErrorRefOnly(coredynamic.TypeName(it))
	}

	return it.Result.PrettyJsonBuffer(prefix, indent)
}

func (it *Result) IsEmptyError() bool {
	return !it.HasError()
}

func (it *Result) IsSuccess() bool {
	return it.HasSafeItems()
}

func (it *Result) CompiledErrorWrapper() *errorwrapper.Wrapper {
	if it.HasError() && it.ErrorWrapper.HasError() {
		return it.ErrorWrapper
	}

	return errnew.Type.Error(errtype.Invalid, it.Result.Error)
}

func (it *Result) IsFailed() bool {
	return it.HasError()
}

func (it *Result) UnmarshalJsonResultTo(unmarshalToReferencePtr interface{}) *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.WithRefs("cannot unmarshal on err json nil.", it)
	}

	if it.HasError() {
		return it.CompiledErrorWrapper()
	}

	err := it.Result.Unmarshal(unmarshalToReferencePtr)

	if err == nil {
		return nil
	}

	return errnew.Type.Error(errtype.Unmarshalling, err)
}

func (it *Result) UnmarshalErrorJson(unmarshalToReferencePtr interface{}) *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.WithRefs("cannot unmarshal on err json nil.", it)
	}

	if it.HasError() {
		return it.CompiledErrorWrapper()
	}

	err := json.Unmarshal(
		it.SafeValues(),
		unmarshalToReferencePtr)

	if err == nil {
		return nil
	}

	return errnew.Type.Error(errtype.Unmarshalling, err)
}
