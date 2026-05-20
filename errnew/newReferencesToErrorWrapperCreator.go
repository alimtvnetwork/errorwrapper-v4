package errnew

import (
	"strings"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

type newReferencesToErrorWrapperCreator struct{}

func (it newReferencesToErrorWrapperCreator) Error(
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

func (it newReferencesToErrorWrapperCreator) ErrorUsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) ErrorWithMessage(
	errType errtype.Variation,
	err error,
	message string,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(err.Error(), message),
		references...)
}

func (it newReferencesToErrorWrapperCreator) ErrorWithMessageUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	message string,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(err.Error(), message),
		references...)
}

func (it newReferencesToErrorWrapperCreator) Msg(
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

func (it newReferencesToErrorWrapperCreator) Quick(
	errType errtype.Variation,
	msg string,
	locations ...string,
) *errorwrapper.Wrapper {
	refValues := make([]ref.Value, 0, len(locations))
	for i, loc := range locations {
		refValues = append(refValues, ref.Value{
			Location: loc,
			Name:     "ref" + strings.TrimSpace(strings.Repeat(" ", 0)) + itoaQuick(i),
		})
	}
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		msg,
		refValues...)
}

func itoaQuick(i int) string {
	if i == 0 {
		return "0"
	}
	digits := ""
	neg := i < 0
	if neg {
		i = -i
	}
	for i > 0 {
		digits = string(rune('0'+i%10)) + digits
		i /= 10
	}
	if neg {
		return "-" + digits
	}
	return digits
}

func (it newReferencesToErrorWrapperCreator) MsgUsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) Type(
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

func (it newReferencesToErrorWrapperCreator) UsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) Many(
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

func (it newReferencesToErrorWrapperCreator) ManyUsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) ErrorWithOne(
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

func (it newReferencesToErrorWrapperCreator) ErrorWithOneUsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) MsgWithOne(
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

func (it newReferencesToErrorWrapperCreator) MsgWithOneUsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) OnlyOne(
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

func (it newReferencesToErrorWrapperCreator) OnlyOneUsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) MergeWrapper(
	currentWrapper *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if currentWrapper == nil || currentWrapper.IsEmpty() {
		return nil
	}

	references := refs.
		NewExistingCollectionPlusAddition(
			currentWrapper.References(),
			additionalReferences...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		currentWrapper.Type(),
		currentWrapper.Error().Error(),
		references)
}

func (it newReferencesToErrorWrapperCreator) MergeWrapperUsingStackSkip(
	stackSkipIndex int,
	currentWrapper *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if currentWrapper == nil || currentWrapper.IsEmpty() {
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

func (it newReferencesToErrorWrapperCreator) Messages(
	variant errtype.Variation,
	references *refs.Collection,
	messages ...string,
) *errorwrapper.Wrapper {
	compiledMessage := errorwrapper.MessagesJoined(messages...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		variant,
		compiledMessage,
		references)
}

func (it newReferencesToErrorWrapperCreator) ErrorMessages(
	variant errtype.Variation,
	err error,
	references *refs.Collection,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	compiledMessage := errorwrapper.MessagesJoined(
		messages...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		variant,
		compiledMessage,
		references)
}

func (it newReferencesToErrorWrapperCreator) ErrorMessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	references *refs.Collection,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	compiledMessage := errorwrapper.MessagesJoined(messages...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
		references)
}

func (it newReferencesToErrorWrapperCreator) MessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	references *refs.Collection,
	messages ...string,
) *errorwrapper.Wrapper {
	compiledMessage := errorwrapper.MessagesJoined(messages...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
		references)
}

func (it newReferencesToErrorWrapperCreator) OneErrorMessages(
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

func (it newReferencesToErrorWrapperCreator) OneErrorMessagesUsingStackSkip(
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

func (it newReferencesToErrorWrapperCreator) MessagesUsingJoiner(
	variant errtype.Variation,
	references *refs.Collection,
	joiner string,
	messages ...string,
) *errorwrapper.Wrapper {
	compiledMessage := strings.Join(messages, joiner)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		variant,
		compiledMessage,
		references)
}

func (it newReferencesToErrorWrapperCreator) MessagesUsingJoinerStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	references *refs.Collection,
	joiner string,
	messages ...string,
) *errorwrapper.Wrapper {
	compiledMessage := strings.Join(messages, joiner)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
		references)
}

// TypeQuick - errorTypeName - (...., items)...
func (it newReferencesToErrorWrapperCreator) TypeQuick(
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
func (it newReferencesToErrorWrapperCreator) TypeQuickUsingStackSkip(
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
