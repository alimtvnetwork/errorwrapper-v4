package errnew

import (
	"reflect"

	"github.com/alimtvnetwork/core-v9/anycmp"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

type newReflectErrToWrapperCreator struct{}

// SetFromTo
//
// Set any object from to toPointer object
//
// Valid Inputs or Supported (https://t.ly/1Lpt):
//  - From, To: (null, null)                          -- do nothing
//  - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//  - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//  - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
//  - From, To: (otherType, *[]byte)                  -- try marshal, reflect
//
// Validations:
//  - Check null, if both null no error return quickly.
//  - NotSupported returns as error.
//      - NotSupported: (from, to) - (..., not pointer)
//      - NotSupported: (from, to) - (null, notNull)
//      - NotSupported: (from, to) - (notNull, null)
//      - NotSupported: (from, to) - not same type and not bytes on any
//  - `From` null or nil is not supported and will return error.
//
// Reference:
//  - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
//  - Method document screenshot    : https://prnt.sc/26dmf5g
func (it newReflectErrToWrapperCreator) SetFromTo(
	from, toPtr interface{},
) *errorwrapper.Wrapper {
	err := coredynamic.ReflectSetFromTo(
		from,
		toPtr)

	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ReflectSetTo,
		err.Error())
}

func (it newReflectErrToWrapperCreator) TypeMismatch(
	left, right interface{},
) *errorwrapper.Wrapper {
	sameStatus := coredynamic.TypeSameStatus(
		left,
		right)

	err := sameStatus.NotEqualSrcDestinationErr()

	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ExpectationMismatch,
		err.Error())
}

func (it newReflectErrToWrapperCreator) ValueMismatchOption(
	isRegardlessOfType bool,
	message string,
	actual, expected interface{},
) *errorwrapper.Wrapper {
	cmp := anycmp.Cmp(actual, expected)

	if cmp.IsEqual() {
		return nil
	}

	if cmp.IsNotEqual() {
		return errorwrapper.NewMsgDisplayErrorReferencesPtr(
			defaultSkipInternal,
			errtype.ExpectationMismatch,
			message,
			it.references(actual, expected)...,
		)
	}

	if isRegardlessOfType {
		return it.ValueMismatchRegardless(
			message,
			actual,
			expected)
	}

	if reflect.DeepEqual(actual, expected) {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.ExpectationMismatch,
		message,
		it.references(actual, expected)...,
	)
}

func (it newReflectErrToWrapperCreator) ValueMismatchRegardless(
	message string,
	actual interface{},
	expected interface{},
) *errorwrapper.Wrapper {
	if actual == expected {
		return nil
	}

	left := converters.AnyToValueString(actual)
	right := converters.AnyToValueString(expected)

	if left == right {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.ExpectationMismatch,
		message,
		it.references(actual, expected)...,
	)
}

func (it newReflectErrToWrapperCreator) references(
	actual interface{},
	expected interface{},
) []ref.Value {
	return []ref.Value{
		{
			Variable: "Actual",
			Value:    actual,
		},
		{
			Variable: "Expected",
			Value:    expected,
		},
		{
			Variable: "ActualType",
			Value:    coredynamic.TypeName(actual),
		},
		{
			Variable: "ExpectedType",
			Value:    coredynamic.TypeName(expected),
		},
	}
}
