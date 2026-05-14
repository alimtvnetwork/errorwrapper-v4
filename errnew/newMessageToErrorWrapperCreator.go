package errnew

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type newMessageToErrorWrapperCreator struct{}

func (it newMessageToErrorWrapperCreator) Create(
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variant,
		message,
	)
}

func (it newMessageToErrorWrapperCreator) Type(
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variant,
		message,
	)
}

func (it newMessageToErrorWrapperCreator) TypeUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
	)
}

func (it newMessageToErrorWrapperCreator) CreateUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
	)
}

func (it newMessageToErrorWrapperCreator) TypeErrorMessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	joinedMessages := stringslice.PrependLineNew(err.Error(), messages)
	compiledMessage := errorwrapper.MessagesJoined(joinedMessages...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		variant,
		compiledMessage,
	)
}

func (it newMessageToErrorWrapperCreator) TypeErrorMessages(
	variant errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	joinedMessages := stringslice.PrependLineNew(err.Error(), messages)
	compiledMessage := errorwrapper.MessagesJoined(joinedMessages...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variant,
		compiledMessage,
	)
}

func (it newMessageToErrorWrapperCreator) MessagesWithRefUsingStackSkip(
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

func (it newMessageToErrorWrapperCreator) New(
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variant,
		message)
}

func (it newMessageToErrorWrapperCreator) Default(
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variant,
		message)
}

func (it newMessageToErrorWrapperCreator) NewUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message)
}

func (it newMessageToErrorWrapperCreator) Many(
	variant errtype.Variation,
	messages ...string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		defaultSkipInternal,
		variant,
		constants.CommaSpace,
		messages...)
}

func (it newMessageToErrorWrapperCreator) ManyUsingStackSkip(
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

func (it newMessageToErrorWrapperCreator) ErrorWithMany(
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

func (it newMessageToErrorWrapperCreator) ErrorWithManyUsingStackSkip(
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
