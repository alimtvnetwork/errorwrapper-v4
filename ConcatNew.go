package errorwrapper

import (
	"errors"
	"strings"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/internal/consts"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

type ConcatNew struct {
	errWp *Wrapper
}

func (it *ConcatNew) isAnyNull() bool {
	return it == nil || it.errWp.IsAnyNull()
}

// isEmpty
//
// Refers to no error for print or doesn't treat this as error.
//
// Conditions (true):
//  - if Wrapper nil, Or,
//  - if Wrapper is StaticEmptyPtr, Or,
//  - if Wrapper .errorType is IsNoError(), Or,
//  - if Wrapper .currentError NOT nil and Wrapper .references.isEmpty()
func (it *ConcatNew) isEmpty() bool {
	return it == nil || it.errWp.IsEmpty()
}

func (it ConcatNew) errorType() errtype.Variation {
	if it.isAnyNull() {
		return errtype.NoError
	}

	return it.errWp.errorType
}

func (it *ConcatNew) clone() *Wrapper {
	if it == nil {
		return nil
	}

	return it.errWp.ClonePtr()
}

func (it ConcatNew) NewStackSkip(stackSkip int) *Wrapper {
	if it.isEmpty() {
		return nil
	}

	return it.errWp.CloneNewStackSkipPtr(
		stackSkip + defaultSkipInternal)
}

func (it ConcatNew) ErrorInterface(
	errInf errcoreinf.BaseErrorOrCollectionWrapper,
) *Wrapper {
	if it.isEmpty() || errInf == nil || errInf.IsEmpty() {
		return it.errWp
	}

	return it.ErrorInterfaceUsingStackSkip(
		defaultSkipInternal,
		errInf)
}

func (it ConcatNew) ErrorInterfaceUsingStackSkip(
	stackSkip int,
	errInf errcoreinf.BaseErrorOrCollectionWrapper,
) *Wrapper {
	if it.isEmpty() || errInf == nil || errInf.IsEmpty() {
		return it.errWp
	}

	newErrWp := CastInterfaceToErrorWrapper(
		errInf)

	return it.WrapperUsingStackSkip(
		stackSkip+defaultSkipInternal,
		newErrWp)
}

func (it ConcatNew) BasicError(
	basicErr errcoreinf.BasicErrWrapper,
) *Wrapper {
	if it.isEmpty() || basicErr == nil || basicErr.IsEmpty() {
		return it.errWp
	}

	return it.BasicErrorUsingStackSkip(
		defaultSkipInternal,
		basicErr)
}

func (it ConcatNew) BasicErrorUsingStackSkip(
	stackSkip int,
	basicErr errcoreinf.BasicErrWrapper,
) *Wrapper {
	if it.isEmpty() || basicErr == nil || basicErr.IsEmpty() {
		return it.errWp
	}

	finalStackSkip := stackSkip + defaultSkipInternal
	newErrWp := NewUsingBasicErrStackSkip(
		finalStackSkip,
		basicErr)

	return it.WrapperUsingStackSkip(
		finalStackSkip,
		newErrWp)
}

// Message
//
// It will create new errorwrapper.Wrapper and combined
// msg with consts.DefaultErrorLineSeparator
func (it ConcatNew) Message(
	appendingMessage string,
) *Wrapper {
	return it.MsgUsingStackSkip(
		defaultSkipInternal,
		appendingMessage)
}

// Messages
//
// It will create new errorwrapper.Wrapper and combined
// msg with consts.DefaultErrorLineSeparator
func (it ConcatNew) Messages(
	errMessages ...string,
) *Wrapper {
	compiled :=
		strings.Join(
			errMessages,
			constants.CommaSpace)

	return it.MsgUsingStackSkip(
		defaultSkipInternal,
		compiled)
}

// MessagesUsingStackSkip
//
// It will create new errorwrapper.Wrapper and combined
// msg with consts.DefaultErrorLineSeparator
func (it ConcatNew) MessagesUsingStackSkip(
	skipStackIndex int,
	errMessages ...string,
) *Wrapper {
	compiled :=
		strings.Join(
			errMessages,
			constants.CommaSpace)

	return it.MsgUsingStackSkip(
		skipStackIndex+defaultSkipInternal,
		compiled)
}

func (it ConcatNew) MsgRefOne(
	skipStackIndex int,
	errMsg string,
	variableName string, valueAny interface{},
) *Wrapper {
	return it.MsgRefs(
		skipStackIndex+defaultSkipInternal,
		errMsg,
		ref.Value{
			Variable: variableName,
			Value:    valueAny,
		})
}

func (it ConcatNew) MsgRefTwo(
	skipStackIndex int,
	errMsg string,
	firstVariable string, firstValue interface{},
	secondVariable string, secondValue interface{},
) *Wrapper {
	return it.MsgRefs(
		skipStackIndex+defaultSkipInternal,
		errMsg,
		ref.Value{
			Variable: firstVariable,
			Value:    firstValue,
		},
		ref.Value{
			Variable: secondVariable,
			Value:    secondValue,
		})
}

func (it ConcatNew) MsgRefs(
	skipStackIndex int,
	errMsg string,
	references ...ref.Value,
) *Wrapper {
	if it.isAnyNull() {
		// nothing exist
		return NewMsgDisplayErrorReferencesPtr(
			skipStackIndex+defaultSkipInternal,
			errtype.Unknown,
			errMsg,
			references...)
	}

	if errMsg == "" {
		cloned2 := it.clone()

		return addReferences(references, cloned2)
	}

	hasReferences := len(references) > 0
	isEmpty := it.isEmpty()
	errType := it.errorType()

	if isEmpty && hasReferences {
		return NewRefWithMessage(
			skipStackIndex+defaultSkipInternal,
			errType,
			errMsg,
			references...)
	} else if isEmpty && !hasReferences {
		return NewMsgDisplayErrorNoReference(
			skipStackIndex+defaultSkipInternal,
			errType,
			errMsg)
	}

	// has references or existing error message
	clonedNew := it.clone()
	newErrMsg := clonedNew.ErrorString() +
		consts.DefaultErrorLineSeparator +
		errMsg

	// here we are mutating over new data.
	// exception case or else the clone code needs
	// to be copied which will break the DRY principle
	clonedNew.currentError =
		errors.New(newErrMsg)
	clonedNew.stackTraces = codestack.NewStacksDefault(
		skipStackIndex+defaultSkipInternal,
		codestack.DefaultStackCount*2)

	return addReferences(references, clonedNew)
}

func (it ConcatNew) MsgRefsOnly(
	references ...ref.Value,
) *Wrapper {
	clonedNew := it.clone()

	if len(references) == 0 {
		return clonedNew
	}

	return addReferences(
		references,
		clonedNew)
}

// Msg
//
// It will create new errorwrapper.Wrapper and combined
// msg with consts.DefaultErrorLineSeparator
func (it ConcatNew) Msg(
	errMsg string,
) *Wrapper {
	return it.MsgUsingStackSkip(
		defaultSkipInternal,
		errMsg)
}

// MsgUsingStackSkip
//
// It will create new errorwrapper.Wrapper and combined
// msg with consts.DefaultErrorLineSeparator
func (it ConcatNew) MsgUsingStackSkip(
	skipStackIndex int,
	errMsg string,
) *Wrapper {
	if it.isAnyNull() {
		return NewMsgDisplayErrorNoReference(
			skipStackIndex+defaultSkipInternal,
			errtype.Unknown,
			errMsg)
	}

	if errMsg == "" {
		return it.
			errWp.
			CloneNewStackSkipPtr(skipStackIndex)
	}

	if it.isEmpty() {
		return NewMsgDisplayErrorNoReference(
			skipStackIndex+defaultSkipInternal,
			errtype.Unknown,
			errMsg)
	}

	clonedNew := it.clone()

	newErrMsg := clonedNew.ErrorString() +
		consts.DefaultErrorLineSeparator +
		errMsg

	// here we are mutating over new data.
	// exception case or else the clone code needs
	// to be copied which will break the DRY principle
	clonedNew.currentError =
		errors.New(newErrMsg)
	clonedNew.stackTraces = codestack.NewStacksDefault(
		skipStackIndex+defaultSkipInternal,
		codestack.DefaultStackCount*2)

	return clonedNew
}

// Error
//
// It will create new errorwrapper.Wrapper and
// combined error with consts.DefaultErrorLineSeparator
func (it ConcatNew) Error(
	err error,
) *Wrapper {
	return it.ErrorUsingStackSkip(
		defaultSkipInternal,
		err)
}

// ErrorUsingStackSkip
//
// It will create new errorwrapper.Wrapper and
// combined error with consts.DefaultErrorLineSeparator
func (it ConcatNew) ErrorUsingStackSkip(
	skipStackIndex int,
	err error,
) *Wrapper {
	if err == nil {
		return it.clone()
	}

	return it.MsgUsingStackSkip(
		skipStackIndex+defaultSkipInternal,
		err.Error())
}

// Errors
//
// It will create new errorwrapper.Wrapper and
// combined errors with consts.DefaultErrorLineSeparator
func (it ConcatNew) Errors(
	errItems ...error,
) *Wrapper {
	return it.ErrorsUsingStackSkip(
		defaultSkipInternal,
		errItems...)
}

// ErrorsUsingStackSkip
//
// It will create new errorwrapper.Wrapper and
// combined errors with consts.DefaultErrorLineSeparator
func (it ConcatNew) ErrorsUsingStackSkip(
	skipStackIndex int,
	errItems ...error,
) *Wrapper {
	if len(errItems) == 0 {
		return it.errWp
	}

	newErrString := ErrorsToStringUsingJoiner(
		consts.DefaultErrorLineSeparator,
		errItems...)

	return it.MsgUsingStackSkip(
		skipStackIndex+defaultSkipInternal,
		newErrString)
}

func (it ConcatNew) NewStackTraces(
	stackTraces ...codestack.Trace,
) *Wrapper {
	tracesCollection := codestack.NewNewTraceCollectionUsing(
		false,
		stackTraces...)

	if it.isAnyNull() {
		empty := EmptyPtr()

		empty.stackTraces = *tracesCollection

		return empty
	}

	cloned := it.clone()
	cloned.stackTraces = *tracesCollection

	return cloned
}

func (it ConcatNew) CloneStackSkip(
	skipStackIndex int,
) *Wrapper {
	return it.CloneStackSkip(
		skipStackIndex + defaultSkipInternal,
	)
}

// Wrapper
//
// Warning : It will not take anything other than error message to combine with.
func (it ConcatNew) Wrapper(
	another *Wrapper,
) *Wrapper {
	return it.WrapperUsingStackSkip(
		defaultSkipInternal,
		another)
}

// WrapperUsingStackSkip
//
// Warning : It will not take anything other than error message to combine with.
func (it ConcatNew) WrapperUsingStackSkip(
	skipStackIndex int,
	right *Wrapper,
) *Wrapper {
	if right == nil || right.IsEmpty() {
		return nil
	}

	if it.isEmpty() {
		// right exist
		return NewMsgDisplayError(
			skipStackIndex+defaultSkipInternal,
			right.errorType,
			right.ErrorString(),
			right.references)
	}

	// Both exist
	clonedNew := it.clone()
	finalErrString := stringslice.NonEmptyJoin(
		[]string{
			clonedNew.ErrorString(),
			right.ErrorString(),
		},
		consts.DefaultErrorLineSeparator,
	)

	// here we are mutating over new data.
	// exception case or else the clone code needs
	// to be copied which will break the DRY principle
	clonedNew.currentError = errors.New(finalErrString)

	if clonedNew.IsReferencesEmpty() && right.HasReferences() {
		// take from right
		clonedNew.references = refs.NewUsingCollection(
			false,
			right.references)
	} else if clonedNew.HasReferences() && right.HasReferences() {
		// both has it
		clonedNew.references = clonedNew.references.
			AddCollections(right.references)

		// last case clonedNew has it,
		// and it is attached to that one.
	}

	// Give new stack traces
	clonedNew.stackTraces = codestack.NewStacksDefault(
		skipStackIndex+defaultSkipInternal,
		codestack.DefaultStackCount*2)

	rightErrType := right.errorType

	if it.errorType().IsEqualVariant(rightErrType) {
		return clonedNew
	}

	// both doesn't contain same error type
	// add the error type into reference

	if clonedNew.references == nil {
		clonedNew.references = refs.New2()
	}

	rightErrTypeVariantStruct := right.errorType.VariantStructure()
	codeWithType := rightErrTypeVariantStruct.CodeTypeName()
	message := rightErrTypeVariantStruct.Message
	clonedNew.references.Add(codeWithType, message)

	return clonedNew
}

func (it ConcatNew) AsConcatenateNewer() ConcatenateNewer {
	return it
}
