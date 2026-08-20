package errnew

import (
	"github.com/alimtvnetwork/core-v9/coredata"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type newJsonToWrapperCreator struct{}

func (it newJsonToWrapperCreator) Jsoner(
	jsoner corejson.Jsoner,
) (*corejson.Result, *errorwrapper.Wrapper) {
	if jsoner == nil {
		return nil, errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			"cannot convert to json if actual object is null!")
	}

	result := jsoner.JsonPtr()

	if result.HasError() {
		return result, errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			result.MeaningfulErrorMessage())
	}

	return result, nil
}

func (it newJsonToWrapperCreator) JsonerToBytes(
	jsoner corejson.Jsoner,
) ([]byte, *errorwrapper.Wrapper) {
	if jsoner == nil {
		return nil, errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			"cannot convert to json if actual object is null!")
	}

	result := jsoner.JsonPtr()

	if result.HasError() {
		return result.SafeValues(), errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			result.MeaningfulErrorMessage())
	}

	return result.SafeValues(), nil
}

func (it newJsonToWrapperCreator) JsonerToDeserialize(
	jsoner corejson.Jsoner,
	toPointer interface{},
) *errorwrapper.Wrapper {
	if jsoner == nil {
		return errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			"cannot convert to json if actual object is null!")
	}

	result := jsoner.JsonPtr()
	err := result.Deserialize(toPointer)

	if err == nil {
		return nil
	}

	references := Unmarshal.references(
		result.SafeBytes(),
		toPointer)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Unmarshalling,
		err.Error(),
		references...)
}

func (it newJsonToWrapperCreator) BytesToDeserializeTo(
	allBytes []byte,
	toPointer interface{},
) *errorwrapper.Wrapper {
	err := corejson.Deserialize.UsingBytes(
		allBytes,
		toPointer)

	if err == nil {
		return nil
	}

	references := Unmarshal.references(
		allBytes,
		toPointer)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Unmarshalling,
		err.Error(),
		references...)
}

func (it newJsonToWrapperCreator) ResultDeserializeTo(
	jsonResult *corejson.Result,
	toPointer interface{},
) *errorwrapper.Wrapper {
	err := jsonResult.Deserialize(toPointer)

	if err == nil {
		return nil
	}

	references := Unmarshal.references(
		jsonResult.SafeBytes(),
		toPointer)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.JsonError,
		err.Error(),
		references...)
}

func (it newJsonToWrapperCreator) JsonerToJsonString(
	jsoner corejson.Jsoner,
) (jsonString string, errWp *errorwrapper.Wrapper) {
	if jsoner == nil {
		return "", errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			"cannot convert to json if actual object is null!")
	}

	result := jsoner.JsonPtr()

	if result.HasError() {
		return result.JsonString(), errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			result.MeaningfulErrorMessage())
	}

	return result.JsonString(), nil
}

func (it newJsonToWrapperCreator) Result(
	jsonResult *corejson.Result,
) *errorwrapper.Wrapper {
	if jsonResult.HasIssuesOrEmpty() {
		return errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			jsonResult.MeaningfulErrorMessage())
	}

	return nil
}

func (it newJsonToWrapperCreator) BytesWithError(
	bytesWithErr *coredata.BytesError,
) *errorwrapper.Wrapper {
	if bytesWithErr.HasError() {
		return errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			bytesWithErr.Error.Error())
	}

	return nil
}

func (it newJsonToWrapperCreator) BytesErrorFunc(
	executor func() ([]byte, error),
) (allBytes []byte, errWp *errorwrapper.Wrapper) {
	if executor == nil {
		return nil, errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.Null,
			"Null executor given to convert to bytes and error wrapper.")
	}

	allBytes, err := executor()

	if err == nil {
		return allBytes, nil
	}

	return allBytes, errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.ExecuteFailed,
		err.Error())
}

func (it newJsonToWrapperCreator) ResultSkipNullIssues(
	jsonResult *corejson.Result,
) *errorwrapper.Wrapper {
	if jsonResult.HasError() {
		return errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.JsonError,
			jsonResult.MeaningfulErrorMessage())
	}

	return nil
}

func (it newJsonToWrapperCreator) MessageReferenceJson(
	errType errtype.Variation,
	message string,
	referencesToJsonItems ...interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	refString := corejson.NewPtr(
		referencesToJsonItems).
		SafeString()

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		message,
		ref.Value{
			Variable: "References (JSON)",
			Value:    refString,
		})
}

func (it newJsonToWrapperCreator) ErrorReferenceJson(
	errType errtype.Variation,
	err error,
	referencesToJsonItems ...interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	refString := corejson.NewPtr(referencesToJsonItems).
		SafeString()

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		err.Error(),
		ref.Value{
			Variable: "References (JSON)",
			Value:    refString,
		})
}
