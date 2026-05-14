package errnew

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
	"gitlab.com/evatix-go/errorwrapper/refs"
)

type newRefToErrorWrapperCreator struct{}

func (it newRefToErrorWrapperCreator) ErrorOne(
	errType errtype.Variation,
	err error,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		NewWithItem(
			constants.One,
			varName,
			value)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		err.Error(),
		reference)
}

func (it newRefToErrorWrapperCreator) ErrorWithMessage(
	errType errtype.Variation,
	err error,
	message string,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		NewWithItem(
			constants.One,
			varName,
			value)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		err.Error()+errorwrapper.MessagesJoiner+message,
		reference)
}

func (it newRefToErrorWrapperCreator) ErrorWithMessageDefault(
	errType errtype.Variation,
	err error,
	message string,
	value interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		NewWithItem(
			constants.One,
			"Ref",
			value)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		err.Error()+errorwrapper.MessagesJoiner+message,
		reference)
}

func (it newRefToErrorWrapperCreator) ErrorOneUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		NewWithItem(
			constants.One,
			varName,
			value)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error(),
		reference)
}

func (it newRefToErrorWrapperCreator) MsgWithMany(
	errType errtype.Variation,
	msg string,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if msg == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		msg,
		references...)
}

func (it newRefToErrorWrapperCreator) ErrorWithOne(
	errType errtype.Variation,
	err error,
	reference ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		err.Error(),
		reference)
}

func (it newRefToErrorWrapperCreator) ErrorWithOneUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	reference ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error(),
		reference)
}

func (it newRefToErrorWrapperCreator) MsgWithOne(
	errType errtype.Variation,
	msg string,
	reference ref.Value,
) *errorwrapper.Wrapper {
	if msg == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		msg,
		reference)
}

func (it newRefToErrorWrapperCreator) MsgWithOneUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	msg string,
	reference ref.Value,
) *errorwrapper.Wrapper {
	if msg == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errType,
		msg,
		reference)
}

func (it newRefToErrorWrapperCreator) ErrorWithMany(
	errType errtype.Variation,
	err error,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		err.Error(),
		references...)
}

func (it newRefToErrorWrapperCreator) ErrorWithManyUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error(),
		references...)
}

func (it newRefToErrorWrapperCreator) OnlyOne(
	errType errtype.Variation,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewRefOne(
		defaultSkipInternal,
		errType,
		varName,
		value)
}

func (it newRefToErrorWrapperCreator) OnlyOneUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewRefOne(
		stackSkipIndex+defaultSkipInternal,
		errType,
		varName,
		value)
}

func (it newRefToErrorWrapperCreator) MsgOne(
	errType errtype.Variation,
	msg string,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewRef1Msg(
		defaultSkipInternal,
		errType,
		msg,
		varName,
		value)
}

func (it newRefToErrorWrapperCreator) MsgOneUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	msg string,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewRef1Msg(
		stackSkipIndex+defaultSkipInternal,
		errType,
		msg,
		varName,
		value)
}

func (it newRefToErrorWrapperCreator) TwoWithMsg(
	errType errtype.Variation,
	msg string,
	varName1 string,
	value1 interface{},
	varName2 string,
	value2 interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewRef2Msg(
		defaultSkipInternal,
		errType,
		msg,
		varName1,
		value1,
		varName2,
		value2,
	)
}

func (it newRefToErrorWrapperCreator) TwoWithMsgUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	msg string,
	varName1 string,
	value1 interface{},
	varName2 string,
	value2 interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewRef2Msg(
		stackSkipIndex+defaultSkipInternal,
		errType,
		msg,
		varName1,
		value1,
		varName2,
		value2,
	)
}

func (it newRefToErrorWrapperCreator) ManyWithMsg(
	errType errtype.Variation,
	message string,
	refValues ...ref.Value,
) *errorwrapper.Wrapper {
	reference := refs.
		New(len(refValues)).
		Adds(refValues...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		message,
		reference)
}

func (it newRefToErrorWrapperCreator) ManyWithMsgUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	message string,
	refValues ...ref.Value,
) *errorwrapper.Wrapper {
	reference := refs.
		New(len(refValues)).
		Adds(refValues...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		message,
		reference)
}

func (it newRefToErrorWrapperCreator) Many(
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

func (it newRefToErrorWrapperCreator) ManyUsingStackSkip(
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

func (it newRefToErrorWrapperCreator) ManyWithError(
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

func (it newRefToErrorWrapperCreator) ManyWithErrorUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	error error,
	refValues ...ref.Value,
) *errorwrapper.Wrapper {
	if error == nil {
		return nil
	}

	return it.ManyWithMsgUsingStackSkip(
		stackSkipIndex+defaultSkipInternal,
		errType,
		error.Error(),
		refValues...)
}

func (it *newReferencesToErrorWrapperCreator) SingleErrorWithMessages(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	referenceOne ref.Value,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	compiledMessage := errorwrapper.MessagesJoined(
		messages...)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
		referenceOne)
}

func (it *newReferencesToErrorWrapperCreator) SingleErrorWithMessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	referenceOne ref.Value,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	compiledMessage := errorwrapper.MessagesJoined(
		messages...)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
		referenceOne)
}

func (it newRefToErrorWrapperCreator) TwoWithError(
	errType errtype.Variation,
	error error,
	varName string,
	val interface{},
	varName2 string,
	val2 interface{},
) *errorwrapper.Wrapper {
	if error == nil {
		return nil
	}

	reference := refs.
		NewWithItem(constants.Capacity2, varName, val).
		Add(varName2, val2)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		error.Error(),
		reference)
}

func (it newRefToErrorWrapperCreator) TwoWithErrorUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	error error,
	varName string,
	val interface{},
	varName2 string,
	val2 interface{},
) *errorwrapper.Wrapper {
	if error == nil {
		return nil
	}

	references := refs.
		NewWithItem(constants.Capacity2, varName, val).
		Add(varName2, val2)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		error.Error(),
		references)
}

func (it newRefToErrorWrapperCreator) ManyUsingWrapper(
	currentWrapper *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if currentWrapper == nil || currentWrapper.IsEmpty() && currentWrapper.IsReferencesEmpty() {
		return nil
	}

	references := currentWrapper.MergeNewReferences(
		additionalReferences...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		currentWrapper.Type(),
		currentWrapper.Error().Error(),
		references)
}

func (it newRefToErrorWrapperCreator) ManyUsingWrapperUsingStackSkip(
	stackSkipIndex int,
	currentWrapper *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if currentWrapper == nil || currentWrapper.IsEmpty() && currentWrapper.IsReferencesEmpty() {
		return nil
	}

	references := currentWrapper.MergeNewReferences(
		additionalReferences...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		currentWrapper.Type(),
		currentWrapper.Error().Error(),
		references)
}

func (it newRefToErrorWrapperCreator) Message(
	variant errtype.Variation,
	varName string,
	value interface{},
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: varName,
			Value:    value,
		})
}

func (it newRefToErrorWrapperCreator) MessageUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	varName string,
	value interface{},
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: varName,
			Value:    value,
		})
}

func (it newRefToErrorWrapperCreator) Messages(
	variant errtype.Variation,
	varName string,
	value interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	return it.MessagesUsingJoinerStackSkip(
		defaultSkipInternal,
		variant,
		varName,
		value,
		constants.Space,
		messages...)
}

func (it newRefToErrorWrapperCreator) MessagesUsingJoiner(
	variant errtype.Variation,
	varName string, value interface{},
	joiner string,
	messages ...string,
) *errorwrapper.Wrapper {
	refsCollection := refs.
		New(1).
		Add(varName, value)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		variant,
		strings.Join(messages, joiner),
		refsCollection)
}

func (it newRefToErrorWrapperCreator) MessagesUsingJoinerStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	varName string,
	value interface{},
	joiner string,
	messages ...string,
) *errorwrapper.Wrapper {
	refsCollection := refs.
		New(1).
		Add(varName, value)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		variant,
		strings.Join(messages, joiner),
		refsCollection)
}

// TypeQuick - errorTypeName - (...., items)...
func (it newRefToErrorWrapperCreator) TypeQuick(
	errType errtype.Variation,
	referencesValues ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.TypeReferenceQuick(
		defaultSkipInternal,
		errType,
		referencesValues...,
	)
}

// TypeQuickUsingStackSkip - errorTypeName - (...., items)...
func (it newRefToErrorWrapperCreator) TypeQuickUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	referencesValues ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.TypeReferenceQuick(
		stackSkipIndex+defaultSkipInternal,
		errType,
		referencesValues...,
	)
}

func (it newRefToErrorWrapperCreator) WithMessagesJoiner(
	variant errtype.Variation,
	varName string, val interface{},
	joiner string,
	messages ...interface{},
) *errorwrapper.Wrapper {
	finalMessage := converters.AnyItemsJoin(
		joiner,
		messages...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		variant,
		finalMessage,
		refs.NewWithItem(1, varName, val))
}

func (it newRefToErrorWrapperCreator) WithMessagesJoinerStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	varName string, val interface{},
	joiner string,
	messages ...interface{},
) *errorwrapper.Wrapper {
	finalMessage := converters.AnyItemsJoin(
		joiner,
		messages...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		variant,
		finalMessage,
		refs.NewWithItem(1, varName, val))
}

func (it newRefToErrorWrapperCreator) CollectionWithMessagesJoiner(
	variant errtype.Variation,
	collection *refs.Collection,
	joiner string,
	messages ...interface{},
) *errorwrapper.Wrapper {
	finalMessage := converters.AnyItemsJoin(
		joiner,
		messages...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		variant,
		finalMessage,
		collection)
}

func (it newRefToErrorWrapperCreator) CollectionWithMessagesJoinerUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	collection *refs.Collection,
	joiner string,
	messages ...interface{},
) *errorwrapper.Wrapper {
	finalMessage := converters.AnyItemsJoin(
		joiner,
		messages...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		variant,
		finalMessage,
		collection)
}

func (it newRefToErrorWrapperCreator) Default(
	variation errtype.Variation,
	reference interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variation,
		"Reference",
		ref.Value{
			Variable: "Reference",
			Value:    reference,
		},
		ref.Value{
			Variable: "Type",
			Value:    coredynamic.TypeName(reference),
		})
}
