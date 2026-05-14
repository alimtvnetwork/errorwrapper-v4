package errnew

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

type newNullToErrorWrapperCreator struct{}

func (it newNullToErrorWrapperCreator) OrWrapper(
	errWp *errorwrapper.Wrapper,
	objectNull interface{},
) *errorwrapper.Wrapper {
	if errWp.HasError() {
		return errWp
	}

	return it.UsingStackSkip(
		defaultSkipInternal,
		"",
		objectNull,
	)
}

func (it newNullToErrorWrapperCreator) OrError(
	errType errtype.Variation,
	err error,
	objectNull interface{},
) *errorwrapper.Wrapper {
	if err != nil {
		return Type.ErrorUsingStackSkip(
			defaultSkipInternal,
			errType,
			err)
	}

	return it.UsingStackSkip(
		defaultSkipInternal,
		"",
		objectNull,
	)
}

func (it newNullToErrorWrapperCreator) OrErrorFunc(
	objectNull interface{},
	errType errtype.Variation,
	errFunc func() error,
) *errorwrapper.Wrapper {
	if errFunc == nil {
		return it.UsingStackSkip(
			defaultSkipInternal,
			"",
			objectNull)
	}

	err := errFunc()
	if err != nil {
		return Type.ErrorUsingStackSkip(
			defaultSkipInternal,
			errType,
			err)
	}

	return nil
}

func (it newNullToErrorWrapperCreator) OrErrorWrapperFunc(
	objectNull interface{},
	executor func() *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if executor == nil {
		return it.UsingStackSkip(
			defaultSkipInternal,
			"",
			objectNull)
	}

	return executor()
}

func (it newNullToErrorWrapperCreator) OrWrapperOrError(
	objectNull interface{},
	errWp *errorwrapper.Wrapper,
	errType errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	if errWp.HasError() {
		return errWp
	}

	if err != nil {
		return Type.ErrorUsingStackSkip(
			defaultSkipInternal,
			errType,
			err)
	}

	return it.UsingStackSkip(
		defaultSkipInternal,
		"",
		objectNull,
	)
}

func (it newNullToErrorWrapperCreator) WithError(
	errType errtype.Variation,
	err error,
	objectNull interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return it.UsingStackSkip(
			defaultSkipInternal,
			"",
			objectNull,
		)
	}

	compiledErrWp := Type.ErrorUsingStackSkip(
		defaultSkipInternal,
		errType,
		err)
	nullErr := it.UsingStackSkip(
		defaultSkipInternal,
		"",
		objectNull,
	)

	return compiledErrWp.ConcatNew().Wrapper(nullErr)
}

func (it newNullToErrorWrapperCreator) Simple(
	objectNull interface{},
) *errorwrapper.Wrapper {
	return it.UsingStackSkip(
		defaultSkipInternal,
		"",
		objectNull,
	)
}

// ManyWithMessage
//
// All the types will be collected
// and returned with details
func (it newNullToErrorWrapperCreator) ManyWithMessage(
	message string,
	objectsNull ...interface{},
) *errorwrapper.Wrapper {
	if len(objectsNull) == 0 {
		return nil
	}

	slice := make([]string, len(objectsNull))
	for i, currentItem := range objectsNull {
		slice[i] = fmt.Sprintf(
			constants.SprintTypeFormat,
			currentItem)
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Null,
		message,
		ref.Value{
			Variable: "With Refs Types",
			Value:    strings.Join(slice, constants.CommaSpace),
		})
}

// ManyByChecking
//
// Only null givens will be added to the list
// if no nul items then nil returns
func (it newNullToErrorWrapperCreator) ManyByChecking(
	objectsNull ...interface{},
) *errorwrapper.Wrapper {
	if len(objectsNull) == 0 {
		return nil
	}

	return it.ManyByCheckingUsingStackSkip(
		defaultSkipInternal,
		objectsNull...)
}

// ManyByCheckingUsingStackSkip
//
// Only null givens will be added to the list
// if no nul items then nil returns
func (it newNullToErrorWrapperCreator) ManyByCheckingUsingStackSkip(
	stackSkipIndex int,
	objectsNull ...interface{},
) *errorwrapper.Wrapper {
	if len(objectsNull) == 0 {
		return nil
	}

	slice := make([]string, 0, len(objectsNull))
	for _, currentItem := range objectsNull {
		if isany.Null(currentItem) {
			slice = append(
				slice,
				fmt.Sprintf(constants.SprintTypeFormat, currentItem))
		}
	}

	if len(slice) == 0 {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errtype.Null,
		"Any of the items are null.",
		ref.Value{
			Variable: "With Refs Types",
			Value:    strings.Join(slice, constants.CommaSpace),
		})
}

func (it newNullToErrorWrapperCreator) WithRefs(
	message string,
	objectNull interface{},
	references ...ref.Value,
) *errorwrapper.Wrapper {
	return it.UsingStackSkip(
		defaultSkipInternal,
		message,
		objectNull,
		references...)
}

func (it newNullToErrorWrapperCreator) Message(
	message string,
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.Null,
		message,
	)
}

func (it newNullToErrorWrapperCreator) Error(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.Null,
		err.Error(),
	)
}

func (it newNullToErrorWrapperCreator) ErrorWithMessage(
	message string,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	finalMessage := message + err.Error()

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.Null,
		finalMessage,
	)
}

func (it newNullToErrorWrapperCreator) WithMessage(
	message string,
	objectNull interface{},
) *errorwrapper.Wrapper {
	return it.UsingStackSkip(
		defaultSkipInternal,
		message,
		objectNull)
}

func (it newNullToErrorWrapperCreator) UsingStackSkip(
	stackSkipIndex int,
	message string,
	objectNull interface{},
	references ...ref.Value,
) *errorwrapper.Wrapper {
	if len(references) == 0 && isany.NotNull(objectNull) {
		return nil
	}

	typeName := coredynamic.SafeTypeName(objectNull)

	if typeName == "" {
		typeName = "interface{}.(nil)"
	}

	mergedReference := refs.PrependReferences(
		false,
		references,
		ref.Value{
			Variable: "Type",
			Value:    typeName,
		})

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errtype.Null,
		message,
		mergedReference...)
}
