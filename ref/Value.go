package ref

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/defaulterr"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
)

type Value struct {
	Variable            string
	Value               interface{}
	compiled            corestr.SimpleStringOnce
	compiledWithoutType corestr.SimpleStringOnce
}

func New(varName string, varValue interface{}) Value {
	return Value{
		Variable: varName,
		Value:    varValue,
	}
}

func NewUsingReferencer(
	ref errcoreinf.Referencer,
) Value {
	if ref == nil {
		return Value{}
	}

	return Value{
		Variable: ref.VarName(),
		Value:    ref.ValueDynamic(),
	}
}

func NewUsingKeyAnyVal(
	ref coreinterface.KeyAnyValueDefiner,
) Value {
	if ref == nil {
		return Value{}
	}

	return Value{
		Variable: ref.KeyName(),
		Value:    ref.ValueAny(),
	}
}

func NewPtr(varName string, varValue interface{}) *Value {
	return &Value{
		Variable: varName,
		Value:    varValue,
	}
}

func NewFromDataModelPtr(
	model *ValueDataModel,
) *Value {
	if model == nil {
		return &Value{}
	}

	return &Value{
		Variable: model.VariableName,
		Value:    model.ValueString,
	}
}

func (it Value) ToPtr() *Value {
	return &it
}

func (it Value) ToNonPtr() Value {
	return it
}

func (it *Value) ReflectSetTo(
	toPointer interface{},
) error {
	return coredynamic.ReflectSetFromTo(it, toPointer)
}

func (it *Value) KeyName() string {
	return it.Variable
}

func (it *Value) ValueAny() interface{} {
	return it.Value
}

func (it *Value) IsVariableNameEqual(name string) bool {
	return it.Variable == name
}

func (it *Value) IsAnyValueEqual(right interface{}) bool {
	isRightNull := isany.Null(right)

	if it == nil && isRightNull {
		return true
	}

	if it == nil || isRightNull {
		return false
	}

	if it.Value == right {
		return true
	}

	if isany.NullBoth(it.Value, right) {
		return true
	}

	if reflect.DeepEqual(it.Value, right) {
		return true
	}

	return false
}

func (it *Value) IsEqualKeyAnyValueDefiner(
	right coreinterface.KeyAnyValueDefiner,
) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	return it.VarName() == right.VariableName() &&
		it.IsAnyValueEqual(right.ValueAny())
}

func (it Value) VarName() string {
	return it.Variable
}

func (it Value) ValueDynamic() interface{} {
	return it.Value
}

func (it Value) VariableValueString() (varName, value string) {
	return it.Variable, converters.AnyTo.ToValueString(it.Value)
}

func (it Value) VariableValueDynamic() (varName string, value interface{}) {
	return it.Variable, it.Value
}

func (it Value) Compile() string {
	return it.FullString()
}

func (it Value) Serialize() ([]byte, error) {
	return it.JsonPtr().Raw()
}

func (it Value) SerializeMust() (jsonBytes []byte) {
	return it.JsonPtr().RawMust()
}

func (it *Value) IsEqualReferencer(ref errcoreinf.Referencer) bool {
	if it == nil && ref == nil {
		return true
	}

	if it == nil || ref == nil {
		return false
	}

	return it.VarName() == ref.VarName() &&
		it.IsAnyValueEqual(ref.ValueDynamic())
}

func (it Value) ValueString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.Value)
}

func (it Value) VariableName() string {
	return it.Variable
}

func (it Value) IsEqual(another Value) bool {
	return it.Variable == another.Variable &&
		it.FullString() == another.FullString()
}

func (it *Value) IsEqualPtr(another *Value) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return true
	}

	return it.Variable == another.Variable &&
		it.IsAnyValueEqual(another.Value)
}

func (it Value) String() string {
	return it.FullString()
}

func (it Value) StringWithoutType() string {
	if it.compiledWithoutType.IsInitialized() {
		return it.compiledWithoutType.String()
	}

	msg := it.Variable +
		constants.SpaceColonSpace +
		fmt.Sprintf(constants.SprintValueFormat, it.Value)

	return it.
		compiledWithoutType.
		GetPlusSetOnUninitialized(msg)
}

func (it Value) FullString() string {
	if it.compiled.IsInitialized() {
		return it.compiled.String()
	}

	compiled := fmt.Sprintf(
		errconsts.SingleReferenceCompile,
		it.Variable,
		it.Value,
		it.Value,
	)

	return it.
		compiled.
		GetPlusSetOnUninitialized(compiled)
}

func (it Value) Clone() Value {
	return Value{
		Variable: it.Variable,
		Value:    it.ValueString(),
	}
}

func (it *Value) ClonePtr() *Value {
	if it == nil {
		return nil
	}

	return &Value{
		Variable: it.Variable,
		Value:    it.ValueString(),
	}
}

func (it *Value) ToDataModel() ValueDataModel {
	return NewDataModel(it)
}

func (it *Value) JsonModel() ValueDataModel {
	return NewDataModel(it)
}

func (it Value) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it Value) MarshalJSON() ([]byte, error) {
	return corejson.Serialize.ToBytesErr(it.JsonModel())
}

func (it *Value) UnmarshalJSON(data []byte) error {
	var dataModel ValueDataModel
	err := corejson.
		Deserialize.
		UsingBytes(data, &dataModel)

	if err == nil {
		it.Value = dataModel.ValueString
		it.Variable = dataModel.VariableName
	}

	return err
}

func (it Value) Json() corejson.Result {
	return corejson.New(it)
}

func (it Value) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Value) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Value, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return it, defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(
		jsonResult.Bytes,
		&it)

	if err != nil {
		return it, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *Value) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Value {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Value) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it Value) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Value) AsJsoner() corejson.Jsoner {
	return &it
}

func (it Value) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return &it
}

func (it Value) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it Value) AsReferencer() errcoreinf.Referencer {
	return &it
}

func (it Value) AsKeyAnyValueDefinerBinder() coreinterface.KeyAnyValueDefinerBinder {
	return &it
}
