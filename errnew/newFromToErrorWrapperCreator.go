package errnew

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

type newFromToErrorWrapperCreator struct{}

func (it newFromToErrorWrapperCreator) Message(
	variant errtype.Variation,
	message string,
	from, to interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		})
}

func (it newFromToErrorWrapperCreator) MessageStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	message string,
	from, to interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		})
}

func (it newFromToErrorWrapperCreator) Create(
	variant errtype.Variation,
	from, to interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variant,
		"",
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		})
}

func (it newFromToErrorWrapperCreator) CreateUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	from, to interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		variant,
		"",
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		})
}

func (it newFromToErrorWrapperCreator) MessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	isIncludeType bool,
	from, to interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	message := errorwrapper.MessagesJoined(
		messages...)

	if isIncludeType {
		return errorwrapper.NewRefWithMessage(
			stackSkipIndex+defaultSkipInternal,
			variant,
			message,
			ref.Value{
				Variable: "From",
				Value:    from,
			},
			ref.Value{
				Variable: "FromType",
				Value:    coredynamic.TypeName(from),
			},
			ref.Value{
				Variable: "To",
				Value:    to,
			},
			ref.Value{
				Variable: "ToType",
				Value:    coredynamic.TypeName(to),
			})
	}

	return errorwrapper.NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		})
}

func (it newFromToErrorWrapperCreator) Messages(
	variant errtype.Variation,
	isIncludeType bool,
	from, to interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	message := errorwrapper.MessagesJoined(
		messages...)

	if isIncludeType {
		return errorwrapper.NewRefWithMessage(
			defaultSkipInternal,
			variant,
			message,
			ref.Value{
				Variable: "From",
				Value:    from,
			},
			ref.Value{
				Variable: "FromType",
				Value:    coredynamic.TypeName(from),
			},
			ref.Value{
				Variable: "To",
				Value:    to,
			},
			ref.Value{
				Variable: "ToType",
				Value:    coredynamic.TypeName(to),
			})
	}

	return errorwrapper.NewRefWithMessage(
		defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		})
}

func (it newFromToErrorWrapperCreator) WithMetaMessages(
	variant errtype.Variation,
	isIncludeType bool,
	from, fromMeta,
	to, toMeta interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	message := errorwrapper.MessagesJoined(
		messages...)

	if isIncludeType {
		return errorwrapper.NewRefWithMessage(
			defaultSkipInternal,
			variant,
			message,
			ref.Value{
				Variable: "From",
				Value:    from,
			},
			ref.Value{
				Variable: "FromType",
				Value:    coredynamic.TypeName(from),
			},
			ref.Value{
				Variable: "FromMeta",
				Value:    fromMeta,
			},
			ref.Value{
				Variable: "To",
				Value:    to,
			},
			ref.Value{
				Variable: "ToType",
				Value:    coredynamic.TypeName(to),
			},
			ref.Value{
				Variable: "ToMeta",
				Value:    toMeta,
			})
	}

	return errorwrapper.NewRefWithMessage(
		defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "FromMeta",
			Value:    fromMeta,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		},
		ref.Value{
			Variable: "ToMeta",
			Value:    toMeta,
		})
}

func (it newFromToErrorWrapperCreator) WithMetaMessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	isIncludeType bool,
	from, fromMeta,
	to, toMeta interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	message := errorwrapper.MessagesJoined(
		messages...)

	if isIncludeType {
		return it.withMetaMessagesUsingStackSkipIncludingType(
			stackSkipIndex,
			variant,
			message,
			from,
			fromMeta,
			to,
			toMeta)
	}

	return errorwrapper.NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "FromMeta",
			Value:    fromMeta,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		},
		ref.Value{
			Variable: "ToMeta",
			Value:    toMeta,
		})
}

func (it newFromToErrorWrapperCreator) withMetaMessagesUsingStackSkipIncludingType(
	stackSkipIndex int,
	variant errtype.Variation,
	message string,
	from interface{},
	fromMeta interface{},
	to interface{},
	toMeta interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		variant,
		message,
		ref.Value{
			Variable: "From",
			Value:    from,
		},
		ref.Value{
			Variable: "FromType",
			Value:    coredynamic.TypeName(from),
		},
		ref.Value{
			Variable: "FromMeta",
			Value:    fromMeta,
		},
		ref.Value{
			Variable: "To",
			Value:    to,
		},
		ref.Value{
			Variable: "ToType",
			Value:    coredynamic.TypeName(to),
		},
		ref.Value{
			Variable: "ToMeta",
			Value:    toMeta,
		})
}
