package errnew

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type newDeserializeToWrapperCreator struct{}

// JsonResultToAny
//
//	Warning : on nil json result may yield error so be sure to check out
//	JsonResultToAnySkipOnNull, JsonResultToAnyOption
//
//	json result deserializes to ptr on fail creates error wrapper
func (it newDeserializeToWrapperCreator) JsonResultToAny(
	jsonResult *corejson.Result,
	toPtr interface{},
) *errorwrapper.Wrapper {
	err := jsonResult.Deserialize(toPtr)

	if err == nil {
		return nil
	}

	references := Unmarshal.references(
		jsonResult.SafeBytes(),
		toPtr)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Deserialize,
		err.Error(),
		references...)
}

// JsonResultToAnyOption
//
//	isSkipOnNull : Skip on nil json result gives nil error wrapper
//
//	json result deserializes to ptr on fail creates error wrapper
func (it newDeserializeToWrapperCreator) JsonResultToAnyOption(
	isSkipOnNull bool,
	jsonResult *corejson.Result,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if isSkipOnNull {
		return it.JsonResultToAnySkipOnNull(
			jsonResult,
			toPtr)
	}

	return it.JsonResultToAny(
		jsonResult,
		toPtr)
}

// JsonResultToAnySkipOnNull
//
//	Skip on nil json result gives nil error wrapper
//
//	json result deserializes to ptr on fail creates error wrapper
func (it newDeserializeToWrapperCreator) JsonResultToAnySkipOnNull(
	jsonResult *corejson.Result,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if jsonResult == nil {
		return nil
	}

	if jsonResult.IsEmpty() && !jsonResult.IsEmptyError() {
		return nil
	}

	err := jsonResult.Deserialize(toPtr)

	if err == nil {
		return nil
	}

	references := Unmarshal.references(
		jsonResult.SafeBytes(),
		toPtr)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Deserialize,
		err.Error(),
		references...)
}

// JsonResultToAnyOnErrAddMsg
//
//	deserializes jsonResult to toPtr if fails creates error with message
func (it newDeserializeToWrapperCreator) JsonResultToAnyOnErrAddMsg(
	onFailErrorMessage string,
	jsonResult *corejson.Result,
	toPtr interface{},
) *errorwrapper.Wrapper {
	err := jsonResult.Deserialize(toPtr)

	if err == nil {
		return nil
	}

	references := Unmarshal.references(
		jsonResult.SafeBytes(),
		toPtr)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Deserialize,
		onFailErrorMessage+err.Error(),
		references...)
}

func (it newDeserializeToWrapperCreator) JsonErrToWrapper(
	errTextInJson error,
) (convErrWp *errorwrapper.Wrapper, parsedErrorWp *errorwrapper.Wrapper) {
	if errTextInJson == nil {
		return nil, nil
	}

	jsonString := errTextInJson.Error()

	return it.JsonStringToWrapperUsingStackSkip(
		defaultSkipInternal,
		false,
		jsonString)
}

func (it newDeserializeToWrapperCreator) JsonResultErrToWrapper(
	errAsJsonResultJson error,
) (convErrWp *errorwrapper.Wrapper, parsedErrorWp *errorwrapper.Wrapper) {
	if errAsJsonResultJson == nil {
		return nil, nil
	}

	jsonString := errAsJsonResultJson.Error()
	jsonResult := corejson.NewResult.EmptyPtr()
	err := corejson.Deserialize.UsingString(
		jsonString,
		jsonResult)

	if err == nil {
		return it.JsonResultToWrapperUsingStackSkip(
			defaultSkipInternal,
			jsonResult)
	}

	// failed to convert to jsonResult
	return nil, errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.Deserialize,
		err.Error(),
		ref.Value{
			Variable: "Error String",
			Value:    jsonString,
		},
		ref.Value{
			Variable: "ConvertFailed",
			Value:    "JsonResultErrToWrapper",
		})
}

// BytesToWrapper
//
//	On empty bytes returns nil.
func (it newDeserializeToWrapperCreator) BytesToWrapper(
	allBytes []byte,
) (errWp *errorwrapper.Wrapper, parsedErrWp *errorwrapper.Wrapper) {
	if len(allBytes) == 0 {
		return nil, nil
	}

	return it.BytesToWrapperUsingStackSkip(
		defaultSkipInternal,
		allBytes)
}

func (it newDeserializeToWrapperCreator) BytesToUnmarshal(
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if len(allBytes) == 0 {
		return nil
	}

	err := corejson.
		Deserialize.
		UsingBytes(allBytes, toPtr)

	if err == nil {
		return nil
	}

	return Unmarshal.Error(err)
}

func (it newDeserializeToWrapperCreator) BytesToAnyPtr(
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if len(allBytes) == 0 {
		return nil
	}

	err := corejson.
		Deserialize.
		UsingBytes(allBytes, toPtr)

	if err == nil {
		return nil
	}

	return Unmarshal.Error(err)
}

func (it newDeserializeToWrapperCreator) JsonResultToWrapper(
	jsonResult *corejson.Result,
) (errWp *errorwrapper.Wrapper, parsedErrWp *errorwrapper.Wrapper) {
	return it.JsonResultToWrapperUsingStackSkip(
		defaultSkipInternal,
		jsonResult)
}

func (it newDeserializeToWrapperCreator) JsonResultToWrapperUsingStackSkip(
	stackSkip int,
	jsonResult *corejson.Result,
) (errWp *errorwrapper.Wrapper, parsedErrWp *errorwrapper.Wrapper) {
	if jsonResult == nil {
		return nil, nil
	}

	emptyPtr := &errorwrapper.Wrapper{}
	err := jsonResult.Deserialize(emptyPtr)

	if err == nil {
		// success in deserializing
		return emptyPtr, nil
	}

	// deserializing failed
	return emptyPtr, errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkip+defaultSkipInternal,
		errtype.Deserialize,
		err.Error(),
		ref.Value{
			Variable: "JsonResultBytes",
			Value:    jsonResult.JsonString(),
		},
		ref.Value{
			Variable: "JsonResultError",
			Value:    jsonResult.MeaningfulErrorMessage(),
		},
		ref.Value{
			Variable: "ConvertFailed",
			Value:    "JsonResultToErrorWrapper",
		})
}

// BytesToWrapperUsingStackSkip
//
//	On empty bytes returns nil.
func (it newDeserializeToWrapperCreator) BytesToWrapperUsingStackSkip(
	stackSkip int,
	allBytes []byte,
) (convErrWp *errorwrapper.Wrapper, parsedErrorWp *errorwrapper.Wrapper) {
	if len(allBytes) == 0 {
		return nil, nil
	}

	emptyPtr := &errorwrapper.Wrapper{}
	err := corejson.Deserialize.UsingBytes(
		allBytes,
		emptyPtr)

	if err == nil {
		// success in deserializing
		return emptyPtr, nil
	}

	// deserializing failed
	return emptyPtr, errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkip+defaultSkipInternal,
		errtype.Deserialize,
		err.Error(),
		ref.Value{
			Variable: "BytesConvertToErrWrapper",
			Value:    allBytes,
		},
		ref.Value{
			Variable: "ConvertFailed",
			Value:    "BytesToErrorWrapper",
		})
}

func (it newDeserializeToWrapperCreator) JsonStringToWrapper(
	isSkipEmptyString bool,
	jsonString string,
) (convErrWp *errorwrapper.Wrapper, parsedErrorWp *errorwrapper.Wrapper) {
	return it.JsonStringToWrapperUsingStackSkip(
		defaultSkipInternal,
		isSkipEmptyString,
		jsonString)
}

func (it newDeserializeToWrapperCreator) JsonStringToWrapperUsingStackSkip(
	stackSkip int,
	isSkipEmptyString bool,
	jsonString string,
) (convErrWp *errorwrapper.Wrapper, parsedErrorWp *errorwrapper.Wrapper) {
	if !isSkipEmptyString && jsonString == "" {
		return nil, errorwrapper.NewMsgDisplayErrorReferencesPtr(
			stackSkip+defaultSkipInternal,
			errtype.Deserialize,
			"Empty string to convert to error wrapper!",
			ref.Value{
				Variable: "JsonString",
				Value:    jsonString,
			},
			ref.Value{
				Variable: "ConvertFailed",
				Value:    "BytesToErrorWrapper",
			})
	}

	return it.BytesToWrapperUsingStackSkip(
		stackSkip+defaultSkipInternal,
		[]byte(jsonString))
}
