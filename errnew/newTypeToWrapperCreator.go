package errnew

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/corecsv"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

type newTypeToWrapperCreator struct{}

func (it newErrorToWrapperCreator) Many(
	mainErrType errtype.Variation,
	errTypes ...errtype.Variation,
) *errorwrapper.Wrapper {
	if len(errTypes) == 0 {
		return nil
	}

	compiledStrings := corestr.New.SimpleSlice.Cap(len(errTypes))
	for _, errType := range errTypes {
		compiledStrings.AddIf(errType.HasError(), errType.String())
	}

	return errorwrapper.NewMessagesUsingJoiner(
		defaultSkipInternal,
		mainErrType,
		errorwrapper.MessagesJoiner,
		compiledStrings.Items...)
}

func (it newErrorToWrapperCreator) ManyUsingStackSkip(
	skipStartStackIndex int,
	mainErrType errtype.Variation,
	errTypes ...errtype.Variation,
) *errorwrapper.Wrapper {
	if len(errTypes) == 0 {
		return nil
	}

	compiledStrings := corestr.New.SimpleSlice.Cap(len(errTypes))
	for _, errType := range errTypes {
		compiledStrings.AddIf(errType.HasError(), errType.String())
	}

	return errorwrapper.NewMessagesUsingJoiner(
		skipStartStackIndex+defaultSkipInternal,
		mainErrType,
		errorwrapper.MessagesJoiner,
		compiledStrings.Items...)
}

func (it *newTypeToWrapperCreator) Messages(
	variant errtype.Variation,
	messages ...string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		defaultSkipInternal,
		variant,
		errorwrapper.MessagesJoiner,
		messages...)
}

func (it *newTypeToWrapperCreator) MessagesUsingStackSkip(
	skipStartStackIndex int,
	variant errtype.Variation,
	messages ...string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		skipStartStackIndex+defaultSkipInternal,
		variant,
		errorwrapper.MessagesJoiner,
		messages...)
}

func (it *newTypeToWrapperCreator) ErrorWithMessages(
	variant errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	joinedStrings := stringslice.PrependLineNew(err.Error(), messages)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variant,
		errorwrapper.MessagesJoined(joinedStrings...))
}

func (it *newTypeToWrapperCreator) ErrorWithMessagesUsingStackSkip(
	skipStartStackIndex int,
	variant errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	joinedStrings := stringslice.PrependLineNew(err.Error(), messages)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		skipStartStackIndex+defaultSkipInternal,
		variant,
		errorwrapper.MessagesJoined(joinedStrings...))
}

func (it newErrorToWrapperCreator) NoType(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingError(
		defaultSkipInternal,
		errtype.Generic,
		err)
}

func (it newErrorToWrapperCreator) Default(
	errType errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		err.Error())
}

func (it newErrorToWrapperCreator) NoTypeUsingStackSkip(
	skipStartStackIndex int,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingError(
		skipStartStackIndex+defaultSkipInternal,
		errtype.Generic,
		err)
}

func (it *newTypeToWrapperCreator) Marshal(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingError(
		defaultSkipInternal,
		errtype.Marshalling,
		err)
}

func (it *newTypeToWrapperCreator) Unmarshal(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingError(
		defaultSkipInternal,
		errtype.Unmarshalling,
		err)
}

func (it *newTypeToWrapperCreator) Error(
	variant errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingError(
		defaultSkipInternal,
		variant,
		err)
}

func (it *newTypeToWrapperCreator) New(
	variant errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingError(
		defaultSkipInternal,
		variant,
		err)
}

func (it *newTypeToWrapperCreator) ErrorWithMessage(
	variant errtype.Variation,
	err error,
	message string,
	collection refs.Collection,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewErrorPlusMsgUsingAllParamsPtr(
		defaultSkipInternal,
		variant,
		true,
		err,
		message,
		&collection)
}

func (it *newTypeToWrapperCreator) ErrorWithMessageUsingStackSkip(
	skipStartStackIndex int,
	variant errtype.Variation,
	err error,
	message string,
	collection refs.Collection,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewErrorPlusMsgUsingAllParamsPtr(
		skipStartStackIndex+defaultSkipInternal,
		variant,
		true,
		err,
		message,
		&collection)
}

func (it *newTypeToWrapperCreator) Default(
	variant errtype.Variation,
) *errorwrapper.Wrapper {
	if variant.IsNoError() {
		return nil
	}

	return errorwrapper.NewPtrUsingStackSkip(
		defaultSkipInternal,
		variant)
}

func (it *newTypeToWrapperCreator) DefaultUsingStackSkip(
	skipStartStackIndex int,
	variant errtype.Variation,
) *errorwrapper.Wrapper {
	if variant.IsNoError() {
		return nil
	}

	return errorwrapper.NewPtrUsingStackSkip(
		skipStartStackIndex+defaultSkipInternal,
		variant)
}

func (it *newTypeToWrapperCreator) Create(
	variant errtype.Variation,
) *errorwrapper.Wrapper {
	if variant.IsNoError() {
		return nil
	}

	return errorwrapper.NewPtrUsingStackSkip(
		defaultSkipInternal,
		variant)
}

func (it *newTypeToWrapperCreator) UsingStackSkip(
	skipStartStackIndex int,
	variant errtype.Variation,
) *errorwrapper.Wrapper {
	if variant.IsNoError() {
		return nil
	}

	return errorwrapper.NewPtrUsingStackSkip(
		skipStartStackIndex+defaultSkipInternal,
		variant)
}

func (it *newTypeToWrapperCreator) ErrorUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingError(
		stackSkipIndex+defaultSkipInternal,
		variant,
		err)
}

func (it *newTypeToWrapperCreator) Refs(
	errType errtype.Variation,
	refValues ...ref.Value,
) *errorwrapper.Wrapper {
	reference := refs.
		New(len(refValues)).
		Adds(refValues...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		"",
		reference)
}

func (it *newTypeToWrapperCreator) References(
	errType errtype.Variation,
	refValues ...ref.Value,
) *errorwrapper.Wrapper {
	reference := refs.
		New(len(refValues)).
		Adds(refValues...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		"",
		reference)
}

func (it *newTypeToWrapperCreator) DirectRefs(
	errType errtype.Variation,
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		message,
		ref.Value{
			Variable: "References",
			Value:    corecsv.AnyItemsToStringDefault(references...),
		})
}

func (it *newRefToErrorWrapperCreator) RefsUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	refValues ...ref.Value,
) *errorwrapper.Wrapper {
	reference := refs.
		New(len(refValues)).
		Adds(refValues...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		"",
		reference)
}

func (it *newRefToErrorWrapperCreator) ErrorWithRefs(
	errType errtype.Variation,
	error error,
	refValues ...ref.Value,
) *errorwrapper.Wrapper {
	if error == nil {
		return nil
	}

	return it.ManyWithMsgUsingStackSkip(
		defaultSkipInternal,
		errType,
		error.Error(),
		refValues...)
}

func (it *newRefToErrorWrapperCreator) ErrorWithRef(
	errType errtype.Variation,
	error error,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	if error == nil {
		return nil
	}

	referenceCollection := refs.
		NewWithItem(
			constants.One,
			varName,
			value)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		error.Error(),
		referenceCollection)
}

func (it *newRefToErrorWrapperCreator) ErrorWithRefUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	error error,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	if error == nil {
		return nil
	}

	referenceCollection := refs.
		NewWithItem(
			constants.One,
			varName,
			value)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		error.Error(),
		referenceCollection)
}
