package errnew

import (
	"os/exec"

	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type newNotFoundErrCreator struct{}

func (it newNotFoundErrCreator) Type() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.NotFound)
}

func (it newNotFoundErrCreator) Error(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.NotFound,
		err.Error(),
	)
}

func (it newNotFoundErrCreator) Reference(
	reference interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFound,
		"Element not found!",
		ref.Value{
			Variable: "reference",
			Value:    reference,
		})
}

func (it newNotFoundErrCreator) MessageRef(
	message string,
	reference interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFound,
		message,
		ref.Value{
			Variable: "reference",
			Value:    reference,
		})
}

func (it newNotFoundErrCreator) MessageRefName(
	message string,
	referenceVariable string,
	references ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFound,
		message,
		ref.Value{
			Variable: referenceVariable,
			Value:    references,
		})
}

func (it newNotFoundErrCreator) Missing(
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Missing,
		message,
		ref.Value{
			Variable: "References",
			Value:    references,
		})
}

func (it newNotFoundErrCreator) Invalid(
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Invalid,
		message,
		ref.Value{
			Variable: "References",
			Value:    references,
		})
}

func (it newNotFoundErrCreator) InvalidData(
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Invalid,
		message,
		ref.Value{
			Variable: "Data",
			Value:    references,
		})
}

func (it newNotFoundErrCreator) InvalidStatus(
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Invalid,
		message,
		ref.Value{
			Variable: "Status",
			Value:    references,
		})
}

func (it newNotFoundErrCreator) InvalidBytes(
	message string,
	allBytes []byte,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Invalid,
		message,
		ref.Value{
			Variable: "InvalidBytes",
			Value:    string(allBytes),
		})
}

func (it newNotFoundErrCreator) All(
	errType errtype.Variation,
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		message,
		ref.Value{
			Variable: "References",
			Value:    references,
		})
}

func (it newNotFoundErrCreator) Message(
	message string,
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.NotFound,
		message,
	)
}

func (it newNotFoundErrCreator) Simple(
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.NotFound,
		message,
	)
}

func (it newNotFoundErrCreator) MessageError(
	message string,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.NotFound,
		message+err.Error(),
	)
}

func (it newNotFoundErrCreator) MessageReference(
	message string,
	reference interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFound,
		message,
		ref.Value{
			Variable: "Not Found",
			Value:    reference,
		})
}

func (it newNotFoundErrCreator) ErrorReference(
	err error,
	reference interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFound,
		err.Error(),
		ref.Value{
			Variable: "Not Found",
			Value:    reference,
		})
}

func (it newNotFoundErrCreator) MessageErrorReference(
	message string,
	err error,
	reference interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFound,
		message+err.Error(),
		ref.Value{
			Variable: "Not Found",
			Value:    reference,
		})
}

func (it newNotFoundErrCreator) ProcessError(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ProcessNotFound,
		err.Error(),
	)
}

func (it newNotFoundErrCreator) ProcessWholeCommand(
	wholeCommand string,
) *errorwrapper.Wrapper {
	if wholeCommand == "" {
		return errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.ProcessNotFound,
			"whole process command is empty!",
		)
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ProcessNotFound,
		wholeCommand,
	)
}

func (it newNotFoundErrCreator) CmdWholeCommand(
	cmd *exec.Cmd,
	wholeCommand string,
) *errorwrapper.Wrapper {
	if cmd != nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ProcessNotFound,
		wholeCommand,
	)
}

func (it newNotFoundErrCreator) Cmd(
	cmd *exec.Cmd,
) *errorwrapper.Wrapper {
	if cmd != nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ProcessNotFound,
		"cmd null or nil!",
	)
}

func (it newNotFoundErrCreator) ProcessMessageError(
	message string,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ProcessNotFound,
		message+err.Error(),
	)
}

func (it newNotFoundErrCreator) Record() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.RecordNotFound,
	)
}

func (it newNotFoundErrCreator) Process() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.NotFoundProcess,
	)
}

func (it newNotFoundErrCreator) User() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.UserNotFound,
	)
}

func (it newNotFoundErrCreator) RedisCmd() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.RedisNullCmd,
	)
}

func (it newNotFoundErrCreator) Id() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.IdNotFound,
	)
}

func (it newNotFoundErrCreator) Name() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.NotFoundName,
	)
}

func (it newNotFoundErrCreator) Data() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.NotFoundData,
	)
}

func (it newNotFoundErrCreator) Group() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.NotFoundGroup,
	)
}

func (it newNotFoundErrCreator) Block() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.BlockNotFound,
	)
}

func (it newNotFoundErrCreator) BlockReference(
	reference interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.BlockNotFound,
		"Block invalid or not found",
		ref.Value{
			Variable: "Block",
			Value:    reference,
		})
}

func (it newNotFoundErrCreator) BlockMessage(
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.BlockNotFound,
		message,
	)
}

func (it newNotFoundErrCreator) Key() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.KeyNotFound,
	)
}

func (it newNotFoundErrCreator) KeyReference(
	keyReference interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.KeyNotFound,
		"Key invalid or not found",
		ref.Value{
			Variable: "Key",
			Value:    keyReference,
		})
}

func (it newNotFoundErrCreator) KeyMessage(
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.KeyNotFound,
		message,
	)
}

func (it newNotFoundErrCreator) KeyError(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.KeyNotFound,
		err.Error(),
	)
}

func (it newNotFoundErrCreator) KeyMessageRef(
	message string,
	key interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.KeyNotFound,
		message,
		ref.Value{
			Variable: "key",
			Value:    key,
		})
}

func (it newNotFoundErrCreator) GroupMessageRef(
	message string,
	group interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFoundGroup,
		message,
		ref.Value{
			Variable: "group",
			Value:    group,
		})
}

func (it newNotFoundErrCreator) GroupReference(group interface{}) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFoundGroup,
		"Group not found!",
		ref.Value{
			Variable: "group",
			Value:    group,
		})
}

func (it newNotFoundErrCreator) NameMessageRef(
	message string,
	name interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFoundName,
		message,
		ref.Value{
			Variable: "name",
			Value:    name,
		})
}

func (it newNotFoundErrCreator) NameMessage(
	message string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.NotFoundName,
		message)
}

func (it newNotFoundErrCreator) NameReference(
	name interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.NotFoundName,
		"",
		ref.Value{
			Variable: "name",
			Value:    name,
		})
}

func (it newNotFoundErrCreator) IdMessage(message string) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.IdNotFound,
		message)
}

func (it newNotFoundErrCreator) IdReference(
	id interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.IdNotFound,
		"Id invalid or not found!",
		ref.Value{
			Variable: "Id",
			Value:    id,
		})
}

func (it newNotFoundErrCreator) IdErrorRef(
	err error,
	id interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.IdNotFound,
		err.Error(),
		ref.Value{
			Variable: "Id",
			Value:    id,
		})
}

func (it newNotFoundErrCreator) IdMessageRef(
	message string,
	id interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.IdNotFound,
		message,
		ref.Value{
			Variable: "Id",
			Value:    id,
		})
}

func (it newNotFoundErrCreator) RecordMessageRef(
	message string,
	record interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.RecordNotFound,
		message,
		ref.Value{
			Variable: "record",
			Value:    record,
		})
}

func (it newNotFoundErrCreator) RecordReference(
	record interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.RecordNotFound,
		"Record invalid or missing or not found!",
		ref.Value{
			Variable: "record",
			Value:    record,
		})
}

func (it newNotFoundErrCreator) RecordError(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.RecordNotFound,
		err.Error(),
	)
}

func (it newNotFoundErrCreator) PathRef(
	path string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.PathNotFound,
		"Path not found or have access issues!",
		ref.Value{
			Variable: "Path",
			Value:    path,
		})
}

func (it newNotFoundErrCreator) PathErr(
	err error,
	path string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.PathNotFound,
		err.Error(),
		ref.Value{
			Variable: "Path",
			Value:    path,
		})
}

func (it newNotFoundErrCreator) PathMessage(
	message string,
	path string,
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.PathNotFound,
		message,
		ref.Value{
			Variable: "Path",
			Value:    path,
		})
}

func (it newNotFoundErrCreator) FileErr(
	err error,
	path string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.FileNotFound,
		err.Error(),
		ref.Value{
			Variable: "FilePath",
			Value:    path,
		})
}

func (it newNotFoundErrCreator) EmptyPath() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.EmptyPath,
	)
}

func (it newNotFoundErrCreator) EmptyString() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.EmptyString,
	)
}

func (it newNotFoundErrCreator) EmptyStatus() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.EmptyStatus,
	)
}

func (it newNotFoundErrCreator) EmptyStates() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.EmptyStates,
	)
}

func (it newNotFoundErrCreator) FileMessage(
	message string,
	path string,
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.FileNotFound,
		message,
		ref.Value{
			Variable: "FilePath",
			Value:    path,
		})
}

func (it newNotFoundErrCreator) DirErr(
	err error,
	path string,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.DirNotFound,
		err.Error(),
		ref.Value{
			Variable: "DirPath",
			Value:    path,
		})
}

func (it newNotFoundErrCreator) DirMessage(
	message string,
	path string,
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.DirNotFound,
		message,
		ref.Value{
			Variable: "DirPath",
			Value:    path,
		})
}

func (it newNotFoundErrCreator) Bytes() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.BytesNotFound,
	)
}

func (it newNotFoundErrCreator) Payload() *errorwrapper.Wrapper {
	return errorwrapper.NewTypeUsingStackSkip(
		defaultSkipInternal,
		errtype.PayloadNotFound,
	)
}

func (it newNotFoundErrCreator) PayloadMessage(
	message string,
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.PayloadNotFound,
		message,
	)
}

func (it newNotFoundErrCreator) PayloadReference(
	references interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.PayloadNotFound,
		"Payload invalid!",
		ref.Value{
			Variable: "Payload",
			Value:    references,
		})
}

func (it newNotFoundErrCreator) PayloadErr(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.PayloadNotFound,
		err.Error(),
	)
}

func (it newNotFoundErrCreator) PayloadErrRef(
	err error,
	references interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.PayloadNotFound,
		err.Error(),
		ref.Value{
			Variable: "Payload",
			Value:    references,
		})
}
