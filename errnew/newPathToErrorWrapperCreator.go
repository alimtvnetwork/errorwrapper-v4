package errnew

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type newPathToErrorWrapperCreator struct{}

func (it newPathToErrorWrapperCreator) Dir(
	err error,
	filePath string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewPath(
		defaultSkipInternal,
		errtype.DirIssue,
		err,
		filePath)
}

func (it newPathToErrorWrapperCreator) NotDir(
	filePath string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewPathMsg(
		defaultSkipInternal,
		errtype.NotValidDirectory,
		"Not a valid dir! (Either not exist, permission issue or a file, expected to be a dir)",
		filePath)
}

func (it newPathToErrorWrapperCreator) NotFile(
	filePath string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewPathMsg(
		defaultSkipInternal,
		errtype.NotValidFile,
		"Not a valid file! (Either not exist, permission issue or a dir, expected to be a file)",
		filePath)
}

func (it newPathToErrorWrapperCreator) Invalid(
	filePath string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewPathMsg(
		defaultSkipInternal,
		errtype.InvalidPath,
		"Path not exist, permission issue or has error!",
		filePath)
}

func (it newPathToErrorWrapperCreator) InvalidUsingStackSkip(
	stackSkipIndex int,
	filePath string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewPathMsg(
		stackSkipIndex+defaultSkipInternal,
		errtype.InvalidPath,
		"Path not exist, permission issue or has error!",
		filePath)
}

func (it newPathToErrorWrapperCreator) InvalidManyUsingStackSkip(
	stackSkipIndex int,
	filePaths ...string,
) *errorwrapper.Wrapper {
	if len(filePaths) == 0 {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errtype.InvalidPath,
		"Path not exist, permission issue or has error!",
		ref.Value{
			Variable: "Path(s)",
			Value:    strings.Join(filePaths, constants.SpaceCommaSpace),
		})
}

func (it newPathToErrorWrapperCreator) InvalidMany(
	filePaths ...string,
) *errorwrapper.Wrapper {
	if len(filePaths) == 0 {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Invalid,
		"Path not exist, permission issue or has error!",
		ref.Value{
			Variable: "Path(s)",
			Value:    strings.Join(filePaths, constants.SpaceCommaSpace),
		})
}

func (it newPathToErrorWrapperCreator) Empty() *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.EmptyPath,
		"Empty path given!")
}

func (it newPathToErrorWrapperCreator) File(
	err error,
	filePath string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewPath(
		defaultSkipInternal,
		errtype.FileIssue,
		err,
		filePath)
}

func (it newPathToErrorWrapperCreator) FileContentIssue(
	msg string,
	filePath string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewPathMsg(
		defaultSkipInternal,
		errtype.ContentValidationFailed,
		msg,
		filePath)
}

func (it newPathToErrorWrapperCreator) Type(
	errType errtype.Variation,
	filePath string,
) *errorwrapper.Wrapper {
	reference := ref.Value{
		Variable: pathRefKeyword,
		Value:    filePath,
	}

	return errorwrapper.NewOnlyRefs(
		defaultSkipInternal,
		errType,
		reference)
}

func (it newPathToErrorWrapperCreator) TypeMsg(
	errType errtype.Variation,
	msg,
	filePath string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewPathMsg(
		defaultSkipInternal,
		errType,
		msg,
		filePath)
}

func (it newPathToErrorWrapperCreator) TypeMsgManyPaths(
	errType errtype.Variation,
	msg string,
	filePaths ...string,
) *errorwrapper.Wrapper {
	if len(filePaths) == 0 {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		msg,
		ref.Value{
			Variable: "Path(s)",
			Value:    strings.Join(filePaths, constants.SpaceCommaSpace),
		})
}

// FromToError
//
//	from, to - refers to newPathToErrorWrapperCreator
func (it newPathToErrorWrapperCreator) FromToError(
	errType errtype.Variation,
	err error,
	from, to string, // refers to from newPathToErrorWrapperCreator
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	references := []ref.Value{
		{
			Variable: "From - Path",
			Value:    from,
		},
		{
			Variable: "To - Path",
			Value:    to,
		},
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		err.Error(),
		references...)
}

func (it newPathToErrorWrapperCreator) FromToMessage(
	errType errtype.Variation,
	message,
	from, to string, // refers to from newPathToErrorWrapperCreator
) *errorwrapper.Wrapper {
	references := []ref.Value{
		{
			Variable: "From - Path",
			Value:    from,
		},
		{
			Variable: "To - Path",
			Value:    to,
		},
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		message,
		references...)
}

func (it newPathToErrorWrapperCreator) FromToErrorMessage(
	errType errtype.Variation,
	err error,
	message,
	from, to string, // refers to from newPathToErrorWrapperCreator
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	references := []ref.Value{
		{
			Variable: "From - Path",
			Value:    from,
		},
		{
			Variable: "To - Path",
			Value:    to,
		},
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoined(message, err.Error()),
		references...)
}

func (it newPathToErrorWrapperCreator) TypeUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	filePath string,
) *errorwrapper.Wrapper {
	reference := ref.Value{
		Variable: pathRefKeyword,
		Value:    filePath,
	}

	return errorwrapper.NewOnlyRefs(
		stackSkipIndex+defaultSkipInternal,
		errType,
		reference)
}

func (it newPathToErrorWrapperCreator) Error(
	variant errtype.Variation,
	err error,
	filePath string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewPath(
		defaultSkipInternal,
		variant,
		err,
		filePath)
}

func (it newPathToErrorWrapperCreator) Messages(
	variant errtype.Variation,
	filePath string,
	messages ...string,
) *errorwrapper.Wrapper {
	if len(messages) == 0 {
		return nil
	}

	return errorwrapper.NewPathMessages(
		defaultSkipInternal,
		variant,
		filePath,
		messages...)
}

func (it newPathToErrorWrapperCreator) ErrorMessages(
	variant errtype.Variation,
	err error,
	filePath string,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil && len(messages) == 0 {
		return nil
	}

	if err != nil {
		messages = stringslice.PrependLineNew(
			err.Error(),
			messages)
	}

	return errorwrapper.NewPathMessages(
		defaultSkipInternal,
		variant,
		filePath,
		messages...)
}

func (it newPathToErrorWrapperCreator) EmptyContent(
	location string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.EmptyContent,
		"empty content received!",
		ref.Value{
			Variable: "Path",
			Value:    location,
		})
}

func (it newPathToErrorWrapperCreator) Marshal(
	err error,
	location string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewPath(
		defaultSkipInternal,
		errtype.Marshalling,
		err,
		location)
}

func (it newPathToErrorWrapperCreator) Unmarshal(
	err error,
	location string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewPath(
		defaultSkipInternal,
		errtype.Unmarshalling,
		err,
		location)
}

func (it newPathToErrorWrapperCreator) ErrorUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	filePath string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewPath(
		stackSkipIndex+defaultSkipInternal,
		variant,
		err,
		filePath)
}

func (it newPathToErrorWrapperCreator) MessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	filePath string,
	messages ...string,
) *errorwrapper.Wrapper {
	if len(messages) == 0 {
		return nil
	}

	return errorwrapper.NewPathMessages(
		stackSkipIndex+defaultSkipInternal,
		variant,
		filePath,
		messages...)
}

func (it newPathToErrorWrapperCreator) ErrorMessagesUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	err error,
	filePath string,
	messages ...string,
) *errorwrapper.Wrapper {
	if err == nil && len(messages) == 0 {
		return nil
	}

	if err != nil {
		messages = stringslice.PrependLineNew(
			err.Error(),
			messages)
	}

	return errorwrapper.NewPathMessages(
		stackSkipIndex+defaultSkipInternal,
		variant,
		filePath,
		messages...)
}
