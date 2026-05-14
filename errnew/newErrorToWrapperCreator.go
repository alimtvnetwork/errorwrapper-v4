package errnew

import (
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
	"gitlab.com/evatix-go/errorwrapper/refs"
)

type newErrorToWrapperCreator struct{}

func (it newErrorToWrapperCreator) TypeOnly(
	errType errtype.Variation,
) *errorwrapper.Wrapper {
	if errType.IsNoError() {
		return nil
	}

	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errType)
}

func (it newErrorToWrapperCreator) TypeFunc(
	errType errtype.Variation,
	executor func() error,
) *errorwrapper.Wrapper {
	return it.TypeFuncStackSkip(
		defaultSkipInternal,
		errType,
		executor)
}

// TypeFuncStackSkip
//
//  Skips nil executors
func (it newErrorToWrapperCreator) TypeFuncStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	executor func() error,
) *errorwrapper.Wrapper {
	if errType.IsNoError() || executor == nil {
		return nil
	}

	err := executor()

	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error())
}

// TypeAnyFunctions
//
//  Halts execution after one has error
//
//  Skips nil executors
func (it newErrorToWrapperCreator) TypeAnyFunctions(
	errType errtype.Variation,
	executors ...func() error,
) *errorwrapper.Wrapper {
	return it.TypeAnyFunctionsStackSkip(
		defaultSkipInternal,
		errType,
		executors...)
}

// TypeAnyFunctionsStackSkip
//
//  Halts execution after one has error
//
//  Skips nil executors
func (it newErrorToWrapperCreator) TypeAnyFunctionsStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	executors ...func() error,
) *errorwrapper.Wrapper {
	if errType.IsNoError() || len(executors) == 0 {
		return nil
	}

	for _, executor := range executors {
		errWp := it.TypeFuncStackSkip(
			stackSkipIndex+defaultSkipInternal,
			errType,
			executor)

		if errWp.HasError() {
			return errWp
		}
	}

	// no issues
	return nil
}

// TypeAllFunctions
//
//  Continues on error and collects all error together.
//
//  Skips nil executors
func (it newErrorToWrapperCreator) TypeAllFunctions(
	errType errtype.Variation,
	executors ...func() error,
) *errorwrapper.Wrapper {
	return it.TypeAllFunctionsStackSkip(
		defaultSkipInternal,
		errType,
		executors...)
}

// TypeAllFunctionsStackSkip
//
//  Continues on error and collects all error together.
//
//  Skips nil executors
func (it newErrorToWrapperCreator) TypeAllFunctionsStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	executors ...func() error,
) *errorwrapper.Wrapper {
	if errType.IsNoError() || len(executors) == 0 {
		return nil
	}

	rawErrCollection := errcore.RawErrCollection{}

	for _, executor := range executors {
		if executor == nil {
			continue
		}

		rawErrCollection.Add(executor())
	}

	if rawErrCollection.IsEmpty() {
		return nil
	}

	compiledErr := rawErrCollection.CompiledError()
	go rawErrCollection.Dispose()

	return it.TypeUsingStackSkip(
		stackSkipIndex+defaultSkipInternal,
		errType,
		compiledErr)
}

func (it newErrorToWrapperCreator) TypeOnlyUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
) *errorwrapper.Wrapper {
	if errType.IsNoError() {
		return nil
	}

	return errorwrapper.NewTypeUsingStackSkip(
		stackSkipIndex+defaultSkipInternal,
		errType)
}

func (it newErrorToWrapperCreator) Error(
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

func (it newErrorToWrapperCreator) ErrorWithMsg(
	err error,
	message string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewUsingErrorAndMessage(
		defaultSkipInternal,
		err,
		message)
}

func (it newErrorToWrapperCreator) TypeMsg(
	errType errtype.Variation,
	err error,
	message string,
) *errorwrapper.Wrapper {
	if err == nil || errType.IsNoError() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(err.Error(), message),
	)
}

func (it newErrorToWrapperCreator) TypeMsgRef(
	errType errtype.Variation,
	err error,
	message string,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	if err == nil || errType.IsNoError() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(err.Error(), message),
		ref.Value{
			Variable: varName,
			Value:    value,
		})
}

func (it newErrorToWrapperCreator) TypeMsgRefUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	message string,
	varName string,
	value interface{},
) *errorwrapper.Wrapper {
	if err == nil || errType.IsNoError() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(err.Error(), message),
		ref.Value{
			Variable: varName,
			Value:    value,
		})
}

func (it newErrorToWrapperCreator) TypeMsgUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	message string,
) *errorwrapper.Wrapper {
	if err == nil || errType.IsNoError() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(err.Error(), message),
	)
}

func (it newErrorToWrapperCreator) TypeMessages(
	errType errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil || errType.IsNoError() {
		return nil
	}

	joinedMessages := stringslice.PrependLineNew(err.Error(), messages)
	compiledMessage := errorwrapper.MessagesJoined(joinedMessages...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		compiledMessage,
	)
}

func (it newErrorToWrapperCreator) TypeMessagesUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil || errType.IsNoError() {
		return nil
	}

	joinedMessages := stringslice.PrependLineNew(err.Error(), messages)
	compiledMessage := errorwrapper.MessagesJoined(joinedMessages...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		compiledMessage,
	)
}

func (it newErrorToWrapperCreator) UsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error())
}

func (it newErrorToWrapperCreator) DefaultUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error())
}

func (it newErrorToWrapperCreator) Create(
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

func (it newErrorToWrapperCreator) Type(
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

func (it newErrorToWrapperCreator) TypeUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error())
}

func (it newErrorToWrapperCreator) TypeWithMessageRefs(
	errType errtype.Variation,
	err error,
	msg string,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		New(len(references)).
		Adds(references...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(err.Error(), msg),
		reference)
}

func (it newErrorToWrapperCreator) TypeWithRefs(
	errType errtype.Variation,
	err error,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		New(len(references)).
		Adds(references...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		err.Error(),
		reference)
}

func (it newErrorToWrapperCreator) TypeWithRefsUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		New(len(references)).
		Adds(references...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		err.Error(),
		reference)
}

func (it newErrorToWrapperCreator) TypeWithMessagesUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	messagesCompiled := stringslice.PrependLineNew(
		err.Error(),
		messages)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(messagesCompiled...),
	)
}
