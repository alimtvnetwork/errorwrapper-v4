package errnew

import (
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coreinstruction"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type newSourceDestinationToErrorWrapperCreator struct{}

func (it newSourceDestinationToErrorWrapperCreator) FromTo(
	variant errtype.Variation,
	from, to interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		"",
		ref.Value{
			Variable: "SetFromTo",
			Value: coreinstruction.FromTo{
				From: converters.AnyTo.ToValueString(from),
				To:   converters.AnyTo.ToValueString(to),
			},
		})
}

func (it newSourceDestinationToErrorWrapperCreator) Create(
	variant errtype.Variation,
	source, destination interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		"",
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}

func (it newSourceDestinationToErrorWrapperCreator) CreateUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	source, destination interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		"",
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}

func (it newSourceDestinationToErrorWrapperCreator) Error(
	variant errtype.Variation,
	err error,
	source, destination interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewRefWithMessage(
		defaultSkipInternal,
		variant,
		err.Error(),
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}

func (it newSourceDestinationToErrorWrapperCreator) ErrorUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	source, destination interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		variant,
		err.Error(),
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}

func (it newSourceDestinationToErrorWrapperCreator) ErrorWithMessage(
	variant errtype.Variation,
	err error,
	msg string,
	source, destination interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewRefWithMessage(
		defaultSkipInternal,
		variant,
		errorwrapper.MessagesJoined(err.Error(), msg),
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}

func (it newSourceDestinationToErrorWrapperCreator) Message(
	variant errtype.Variation,
	source, destination interface{},
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewRefWithMessage(
		defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}

func (it newSourceDestinationToErrorWrapperCreator) Messages(
	variant errtype.Variation,
	source, destination interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	message := errorwrapper.MessagesJoined(
		messages...)

	return errorwrapper.NewRefWithMessage(
		defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}

func (it newSourceDestinationToErrorWrapperCreator) MessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	source, destination interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	message := errorwrapper.MessagesJoined(
		messages...)

	return errorwrapper.NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "Source",
			Value:    source,
		},
		ref.Value{
			Variable: "Destination",
			Value:    destination,
		})
}
