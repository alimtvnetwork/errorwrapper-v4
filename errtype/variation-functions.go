package errtype

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/corecsv"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl/enumtype"
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/coremath"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
)

func (it Variation) TypeNameCodeMessage() string {
	currentStruct := errTypesMap[it]

	return fmt.Sprintf(
		errconsts.ErrorCodeHyphenTypeNameWithLineFormat,
		currentStruct.Variant,
		currentStruct.Name,
		currentStruct.Message)
}

func (it Variation) TypeNameCodeMessageRef(reference interface{}) string {
	currentStruct := errTypesMap[it]

	return fmt.Sprintf(
		errconsts.CombineMessageNoAdditionalMessageFormat,
		currentStruct.Name,
		currentStruct.Variant,
		currentStruct.Message,
		reference)
}

func (it Variation) String() string {
	return errTypesMap[it].Name
}

func (it Variation) Message() string {
	return errTypesMap[it].Message
}

func (it Variation) MessageToRawErrType() errcore.RawErrorType {
	return errcore.RawErrorType(errTypesMap[it].Message)
}

func (it Variation) CombineNoRefs(
	additionalMessage string,
) string {
	variantStruct := errTypesMap[it]

	return variantStruct.CombineNoRefs(
		additionalMessage,
	)
}

func (it Variation) Combine(
	additionalMessage,
	varName string,
	val interface{},
) string {
	variantStruct := errTypesMap[it]

	return variantStruct.Combine(
		additionalMessage,
		varName,
		val)
}

func (it Variation) Name() string {
	return errTypesMap[it].Name
}

func (it Variation) Error(
	additionalMessage,
	varName string,
	val interface{},
) error {
	msg := it.Combine(
		additionalMessage,
		varName,
		val)

	return errors.New(strings.ToLower(msg))
}

func (it Variation) ErrorReferences(
	additionalMessage string,
	references ...interface{},
) error {
	referencesString := converters.AnyTo.ItemsJoin(
		constants.Space,
		references...)

	msg := it.Combine(
		additionalMessage,
		"Refs",
		referencesString)

	return errors.New(strings.ToLower(msg))
}

func (it Variation) ErrorNoRefs(
	additionalMessage string,
) error {
	msg := it.CombineNoRefs(
		additionalMessage,
	)

	return errors.New(strings.ToLower(msg))
}

func (it Variation) Panic(
	additionalMessage,
	varName string,
	val interface{},
) {
	msg := it.Combine(
		additionalMessage,
		varName,
		val)

	panic(msg)
}

func (it Variation) PanicNoRefs(
	additionalMessage string,
) {
	msg := it.CombineNoRefs(
		additionalMessage,
	)

	panic(msg)
}

func (it Variation) VariantStructure() VariantStructure {
	return errTypesMap[it]
}

func (it Variation) VariantStructurePtr() *VariantStructure {
	variantStruct := errTypesMap[it]

	return &variantStruct
}

// TypeName
//
//  Refers to the reflection enum type name.
func (it Variation) TypeName() string {
	return typeName
}

// CodeWithTypeName
//
// 	errconsts.ErrorCodeHyphenTypeNameFormat  = "(#%d - %s)"
func (it Variation) CodeWithTypeName() string {
	return fmt.Sprintf(
		errconsts.ErrorCodeHyphenTypeNameFormat,
		it,
		it.Name(),
	)
}

// CodeTypeNameWithCustomMessage
//
// 	errconsts.ErrorCodeHyphenTypeNameWithLineFormat = "(#%d - %s) %s"
func (it Variation) CodeTypeNameWithCustomMessage(
	customMessage string,
) string {
	return fmt.Sprintf(
		errconsts.ErrorCodeHyphenTypeNameWithLineFormat,
		it,
		it.Name(),
		customMessage,
	)
}

func (it Variation) ReferencesCsv(
	additionalMessage string,
	references ...interface{},
) string {
	csvReferenceString := corecsv.AnyItemsToStringDefault(
		references...)
	variantStructure := it.VariantStructure()

	if additionalMessage == "" {
		return fmt.Sprintf(
			errconsts.CombineMessageNoAdditionalMessageFormat,
			variantStructure.Name,
			it,
			variantStructure.Message,
			csvReferenceString)
	}

	// has additional message
	return fmt.Sprintf(
		errconsts.CombineMessageAdditionalMessageFormat,
		variantStructure.Name,
		it,
		variantStructure.Message,
		additionalMessage,
		csvReferenceString)
}

func (it Variation) ReferencesLines(
	additionalMessage string,
	referencesLines ...interface{},
) string {
	variantStructure := it.VariantStructure()
	line := converters.AnyTo.ItemsJoin(constants.CommaSpace, referencesLines...)

	if additionalMessage == "" {
		return fmt.Sprintf(
			errconsts.CombineMessageNoAdditionalMessageFormat,
			variantStructure.Name,
			it,
			variantStructure.Message,
			line)
	}

	// has additional message
	return fmt.Sprintf(
		errconsts.CombineMessageAdditionalMessageFormat,
		variantStructure.Name,
		it,
		variantStructure.Message,
		additionalMessage,
		line)
}

func (it Variation) ReferencesLinesError(
	additionalMessage string,
	referencesLines ...interface{},
) error {
	compiledLine := it.ReferencesLines(
		additionalMessage,
		referencesLines...)

	return errors.New(compiledLine)
}

func (it Variation) ReferencesCsvError(
	additionalMessage string,
	references ...interface{},
) error {
	csvReferenceString := it.ReferencesCsv(
		additionalMessage,
		references...)

	return errors.New(csvReferenceString)
}

func (it Variation) ShortReferencesCsv(
	references ...interface{},
) string {
	if len(references) == 0 {
		return ""
	}

	csvReferenceString := corecsv.AnyItemsToStringDefault(
		references...)

	return it.CodeTypeNameWithReference(csvReferenceString)
}

func (it Variation) ShortReferencesCsvError(
	references ...interface{},
) error {
	if len(references) == 0 {
		return nil
	}

	csvReferenceString := corecsv.AnyItemsToStringDefault(
		references...)

	finalLine := it.CodeTypeNameWithReference(csvReferenceString)

	return errors.New(finalLine)
}

// ExplicitCodeValueName
//
// 	errconsts.ErrorCodeWithTypeNameFormat = "(Code - #%d) : %s"
func (it Variation) ExplicitCodeValueName() string {
	return fmt.Sprintf(
		errconsts.ErrorCodeWithTypeNameFormat,
		it,
		it.Name())
}

// CodeTypeNameWithReference
//
// 	errconsts.ErrorCodeHyphenTypeNameWithReferencesFormat = "(#%d - %s - {%v})"
func (it Variation) CodeTypeNameWithReference(
	referenceLine string,
) string {
	return fmt.Sprintf(
		errconsts.TypeNameWithReferencesFormat,
		it.Name(),
		referenceLine,
	)
}

// CodeTypeNameWithReferences
//
// 	errconsts.ErrorCodeHyphenTypeNameWithReferencesFormat = "(#%d - %s - {%v})"
func (it Variation) CodeTypeNameWithReferences(
	references ...string,
) string {
	return fmt.Sprintf(
		errconsts.ErrorCodeHyphenTypeNameWithReferencesFormat,
		it,
		it.Name(),
		references,
	)
}

func (it Variation) ValueInt16() int16 {
	return int16(it)
}

func (it Variation) ValueInt() int {
	return int(it)
}

func (it Variation) ValueUInt() uint {
	return uint(it)
}

func (it Variation) CategoryName() string {
	return it.Name()
}

func (it Variation) Serialize() ([]byte, error) {
	return it.MarshalJSON()
}

func (it Variation) SerializeMust() (jsonBytes []byte) {
	jsonBytes, err := it.MarshalJSON()
	errcore.MustBeEmpty(err)

	return jsonBytes
}

func (it Variation) IsErrorTyperEqual(
	errTyper errcoreinf.BaseErrorTyper,
) bool {
	if errTyper == nil {
		return false
	}

	return it.Name() == errTyper.Name() &&
		it.Value() == errTyper.Value() &&
		it.TypeName() == errTyper.TypeName()
}

// TypenameString
//
//  Refers to Name()
func (it Variation) TypenameString() string {
	return it.Name()
}

func (it Variation) TypeMessage() string {
	return errTypesMap[it].Message
}

func (it Variation) BaseErrorTyper() errcoreinf.BaseErrorTyper {
	return &it
}

func (it Variation) ErrTypeDetailDefiner() errcoreinf.ErrTypeDetailDefiner {
	return &it
}

func (it Variation) Value() uint16 {
	return uint16(it)
}

func (it Variation) ValueByte() byte {
	val := it.RawValue()

	if coremath.IsRangeWithin.Integer64.ToByte(int64(val)) {
		return byte(val)
	}

	return constants.Zero
}

func (it Variation) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.Name() == name {
			return true
		}
	}

	return false
}

func (it Variation) ToNumberString() string {
	return it.ValueString()
}

func (it Variation) ValueInt8() int8 {
	val := it.RawValue()

	if val <= math.MaxInt8 {
		return int8(val)
	}

	return constants.InvalidValue
}

func (it Variation) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variation) ValueInt32() int32 {
	return int32(it)
}

func (it Variation) ValueString() string {
	return strconv.Itoa(int(it.ValueUInt16()))
}

func (it Variation) RangeNamesCsv() string {
	return rangesCsvNameStringOnce.String()
}

func (it Variation) JsonModel() JsonModel {
	return JsonModel{
		Id:       it.ValueUInt16(),
		Category: it.Name(),
	}
}

func (it Variation) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it Variation) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModelAny())
}

func (it *Variation) UnmarshalJSON(data []byte) error {
	var model JsonModel
	err := corejson.Deserialize.UsingBytes(
		data, &model)

	if err == nil {
		*it = Variation(model.Id)
	}

	return err
}

// Format
//
//  Outputs name and
//  value by given format.
//
// sample-format :
//  - "Enum of {type-name} - {name} - {value}"
//
// sample-format-output :
//  - "Enum of EnumFullName - Invalid - 0"
//
// Key-Meaning :
//  - {type-name} : represents type-name string
//  - {name}      : represents name string
//  - {value}     : represents value string
func (it Variation) Format(format string) (compiled string) {
	return enumimpl.FormatUsingFmt(
		it,
		format)
}

func (it Variation) MinMaxAny() (min, max interface{}) {
	return NoError, MaxError
}

func (it Variation) MinValueString() string {
	return strconv.Itoa(minValue)
}

func (it Variation) MaxValueString() string {
	return strconv.Itoa(maxValue)
}

func (it Variation) MaxInt() int {
	return maxValue
}

func (it Variation) MinInt() int {
	return minValue
}

func (it Variation) RangesDynamicMap() map[string]interface{} {
	slice := make(map[string]interface{}, maxValue+1)

	for i := 0; i <= maxValue; i++ {
		slice[it.Name()] = i
	}

	return slice
}

func (it Variation) AllNameValues() []string {
	return allNameWithValuesOnce.List()
}

func (it Variation) OnlySupportedErr(names ...string) error {
	if it.IsAnyNamesOf(names...) {
		return nil
	}

	return errcore.RangesOnlySupportedType.ErrorRefOnly(names)
}

func (it Variation) OnlySupportedMsgErr(message string, names ...string) error {
	if it.IsAnyNamesOf(names...) {
		return nil
	}

	return errcore.RangesOnlySupportedType.Error(message, names)
}

func (it Variation) IntegerEnumRanges() []int {
	slice := make([]int, maxValue+1)

	for i := 0; i <= maxValue; i++ {
		slice[i] = i
	}

	return slice
}

func (it Variation) EnumType() enuminf.EnumTyper {
	return enumtype.UnsignedInteger16
}

func (it Variation) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variation) IsEqualVariant(errType Variation) bool {
	return it == errType
}

func (it Variation) NameValue() string {
	return enumimpl.NameWithValue(it)
}

// IsValid
//
//  For error type IsValid refers to NoError type
func (it Variation) IsValid() bool {
	return it == NoError
}

// IsInvalid
//
//  For error type Invalid refers to any error but NoError type
func (it Variation) IsInvalid() bool {
	return it != NoError
}

func (it Variation) IsRawValue(rawValue uint16) bool {
	return it.RawValue() == rawValue
}

func (it Variation) IsEmptyError() bool {
	return it == NoError
}

func (it Variation) RawValue() uint16 {
	panic("implement me")
}

func (it Variation) ErrorTypeAsBasicEnum() enuminf.BasicEnumer {
	return &it
}

func (it Variation) AsBasicErrorTyper() errcoreinf.BasicErrorTyper {
	return &it
}
