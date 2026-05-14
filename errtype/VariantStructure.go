package errtype

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
)

type VariantStructure struct {
	Name    string
	Message string
	Variant Variation
}

// String
//
// 	errconsts.VariantStructStringFormat = "%s (Code - %d) : %s"
func (it VariantStructure) String() string {
	return it.TypeNameCodeMessage()
}

func (it VariantStructure) MessageToRawType() errcore.RawErrorType {
	return errcore.RawErrorType(it.Message)
}

// TypeNameCodeMessage
//
// 	errconsts.VariantStructStringFormat = "%s (Code - %d) : %s"
func (it VariantStructure) TypeNameCodeMessage() string {
	return fmt.Sprintf(
		errconsts.VariantStructStringFormat,
		it.Name,
		it.Variant,
		it.Message)
}

// CodeTypeName
//
// 	errconsts.ErrorCodeHyphenTypeNameFormat  = "(#%d - %s)"
func (it VariantStructure) CodeTypeName() string {
	return it.Variant.CodeWithTypeName()
}

// CodeTypeNameWithCustomMessage
//
// 	errconsts.ErrorCodeHyphenTypeNameWithLineFormat = "(#%d - %s) %s"
func (it *VariantStructure) CodeTypeNameWithCustomMessage(
	customMessage string,
) string {
	return fmt.Sprintf(
		errconsts.ErrorCodeHyphenTypeNameWithLineFormat,
		it.Variant,
		it.Name,
		customMessage,
	)
}

func (it *VariantStructure) CodeTypeNameWithReference(
	referenceLine string,
) string {
	return fmt.Sprintf(
		errconsts.ErrorCodeHyphenTypeNameWithReferencesFormat,
		it.Variant,
		it.Name,
		referenceLine,
	)
}

func (it *VariantStructure) CodeTypeNameWithReferences(
	references ...string,
) string {
	return it.Variant.CodeTypeNameWithReferences(references...)
}

func (it VariantStructure) MsgReferenceValues(msg, referenceValue string) string {
	str := fmt.Sprintf(errconsts.VariantStructMsgReferenceValuesFormat,
		it.String(),
		msg,
		referenceValue)

	return str
}

func (it VariantStructure) ReferenceValues(referenceValue string) string {
	str := fmt.Sprintf(errconsts.VariantStructReferenceValuesFormat,
		it.String(),
		referenceValue)

	return str
}

func (it VariantStructure) CombineNoRefs(
	additionalMessage string,
) string {
	if additionalMessage == "" {
		message := fmt.Sprintf(
			errconsts.CombineMessageNoReferencesNoAdditionalMessageFormat,
			it.Name,
			it.Variant,
			it.Message,
		)

		return message
	}

	message := fmt.Sprintf(
		errconsts.CombineMessageNoReferencesAdditionalMessageFormat,
		it.Name,
		it.Variant,
		it.Message,
		additionalMessage,
	)

	return message
}

// CombineRefs Conditional checks if refs compile is empty then ignore it and prints using CombineNoRefs
func (it VariantStructure) CombineRefs(
	additionalMessage string,
	refsCompiled string,
) string {
	if refsCompiled == "" {
		return it.CombineNoRefs(additionalMessage)
	}

	if additionalMessage == "" {
		message := fmt.Sprintf(
			errconsts.CombineMessageNoAdditionalMessageFormat,
			it.Name,
			it.Variant,
			it.Message,
			refsCompiled,
		)

		return message
	}

	message := fmt.Sprintf(
		errconsts.CombineMessageAdditionalMessageFormat,
		it.Name,
		it.Variant,
		it.Message,
		additionalMessage,
		refsCompiled,
	)

	return message
}

func (it VariantStructure) Combine(
	additionalMessage,
	varName string,
	val interface{},
) string {
	valTypeName := fmt.Sprintf(constants.SprintTypeFormat, val)
	reference := fmt.Sprintf(
		errconsts.ReferenceWithTypeFormat,
		varName,
		valTypeName,
		val,
	)

	if additionalMessage == "" {
		return fmt.Sprintf(
			errconsts.CombineMessageNoAdditionalMessageFormat,
			it.Name,
			it.Variant,
			it.Message,
			reference,
		)
	}

	message := fmt.Sprintf(
		errconsts.CombineMessageAdditionalMessageFormat,
		it.Name,
		it.Variant,
		it.Message,
		additionalMessage,
		reference,
	)

	return message
}

func (it VariantStructure) Error(
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

func (it VariantStructure) ErrorNoRefs(
	additionalMessage string,
) error {
	msg := it.CombineNoRefs(
		additionalMessage,
	)

	return errors.New(strings.ToLower(msg))
}

func (it VariantStructure) Panic(
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

func (it VariantStructure) PanicNoRefs(
	additionalMessage string,
) {
	msg := it.CombineNoRefs(
		additionalMessage,
	)

	panic(msg)
}
