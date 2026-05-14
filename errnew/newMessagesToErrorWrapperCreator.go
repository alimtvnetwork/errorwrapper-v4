package errnew

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type newMessagesToErrorWrapperCreator struct{}

func (it newMessagesToErrorWrapperCreator) ErrorWithRef(
	variant errtype.Variation,
	err error,
	reference ref.Value,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	compiledMessages := stringslice.PrependLineNew(err.Error(), messages)
	compiledMessage := errorwrapper.MessagesJoined(compiledMessages...)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		compiledMessage,
		reference)
}

func (it newMessagesToErrorWrapperCreator) ErrorWithRefUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	reference ref.Value,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	compiledMessages := stringslice.PrependLineNew(err.Error(), messages)
	compiledMessage := errorwrapper.MessagesJoined(compiledMessages...)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
		reference)
}

func (it newMessagesToErrorWrapperCreator) WithRef(
	variant errtype.Variation,
	reference ref.Value,
	messages ...string,
) *errorwrapper.Wrapper {
	compiledMessage := errorwrapper.MessagesJoined(messages...)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		compiledMessage,
		reference)
}

func (it newMessagesToErrorWrapperCreator) WithRefUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	reference ref.Value,
	messages ...string,
) *errorwrapper.Wrapper {
	compiledMessage := errorwrapper.MessagesJoined(messages...)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
		reference)
}

func (it newMessagesToErrorWrapperCreator) WithOnlyRefs(
	variant errtype.Variation,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		constants.EmptyString,
		references...)
}

func (it newMessagesToErrorWrapperCreator) Single(
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variant,
		message)
}

func (it newMessagesToErrorWrapperCreator) SingleUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message)
}

func (it newMessagesToErrorWrapperCreator) Many(
	variant errtype.Variation,
	messages ...string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		defaultSkipInternal,
		variant,
		constants.CommaSpace,
		messages...)
}

func (it newMessagesToErrorWrapperCreator) Create(
	variant errtype.Variation,
	messages ...string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		defaultSkipInternal,
		variant,
		constants.CommaSpace,
		messages...)
}

func (it newMessagesToErrorWrapperCreator) ManyUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	messages ...string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		stackSkipIndex+defaultSkipInternal,
		variant,
		constants.CommaSpace,
		messages...)
}

func (it newMessagesToErrorWrapperCreator) ErrorWithMany(
	variant errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	appendSlice :=
		stringslice.AppendLineNew(
			messages,
			err.Error())

	return errorwrapper.NewMessagesUsingJoiner(
		defaultSkipInternal,
		variant,
		constants.CommaSpace,
		appendSlice...)
}

func (it newMessagesToErrorWrapperCreator) ErrorWithManyUsingStackSkip(
	stackStartIndex int,
	variant errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	appendSlice :=
		stringslice.AppendLineNew(
			messages,
			err.Error())

	return errorwrapper.NewMessagesUsingJoiner(
		stackStartIndex+defaultSkipInternal,
		variant,
		constants.CommaSpace,
		appendSlice...)
}
