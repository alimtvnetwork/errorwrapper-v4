package errnew

import (
	"reflect"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

type newUnmarshalWrapperCreator struct{}

func (it newUnmarshalWrapperCreator) Error(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.Unmarshalling,
		err.Error())
}

func (it newUnmarshalWrapperCreator) MessageError(
	message string,
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.Unmarshalling,
		message+err.Error())
}

func (it newUnmarshalWrapperCreator) Message(
	message string,
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.Unmarshalling,
		message)
}

func (it newUnmarshalWrapperCreator) MessageRef(
	message string,
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return it.MessageRefUsingStackSkip(
		defaultSkipInternal,
		message,
		allBytes,
		toPtr)
}

func (it newUnmarshalWrapperCreator) ErrorRef(
	err error,
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return it.MessageRefUsingStackSkip(
		defaultSkipInternal,
		err.Error(),
		allBytes,
		toPtr)
}

func (it newUnmarshalWrapperCreator) MessageErrorRef(
	message string,
	err error,
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return it.MessageRefUsingStackSkip(
		defaultSkipInternal,
		message+err.Error(),
		allBytes,
		toPtr)
}

func (it newUnmarshalWrapperCreator) Reference(
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	return it.ReferenceUsingStackSkip(
		defaultSkipInternal,
		allBytes,
		toPtr)
}

func (it newUnmarshalWrapperCreator) ReferenceUsingStackSkip(
	stackSkip int,
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	references := it.references(
		allBytes,
		toPtr)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkip+defaultSkipInternal,
		errtype.Unmarshalling,
		"Unmarshalling or deserializing failed.",
		references...)
}

func (it newUnmarshalWrapperCreator) MessageRefUsingStackSkip(
	stackSkip int,
	message string,
	allBytes []byte,
	toPtr interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	references := it.references(
		allBytes,
		toPtr)

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		stackSkip+defaultSkipInternal,
		errtype.Unmarshalling,
		message,
		references...)
}

func (it newUnmarshalWrapperCreator) ptrStatus(
	toPtr interface{},
) (typeName string, isNull, isPtr bool) {
	typeName = "null pointer, cannot retrieve type"
	currentType := reflect.TypeOf(toPtr)

	if isany.Defined(currentType) {
		typeName = currentType.String()
		isPtr = currentType.Kind() == reflect.Ptr
		isNull = isany.Null(toPtr)
	} else {
		isNull = true
	}

	return typeName, isNull, isPtr
}

func (it newUnmarshalWrapperCreator) references(
	allBytes []byte,
	toPtr interface{},
) []ref.Value {
	var bytesString string
	if len(allBytes) > 0 {
		bytesString = string(allBytes)
	}

	typeName, isNull, isPtr := it.ptrStatus(toPtr)

	return []ref.Value{
		{
			Variable: "From Bytes",
			Value:    bytesString,
		},
		{
			Variable: "ToPtr Actual",
			Value:    toPtr,
		},
		{
			Variable: "To Type",
			Value:    typeName,
		},
		{
			Variable: "IsPointerType",
			Value:    isPtr,
		},
		{
			Variable: "IsNull",
			Value:    isNull,
		},
	}
}

func (it newUnmarshalWrapperCreator) BytesToDeserializeTo(
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

func (it newUnmarshalWrapperCreator) JsonResultToDeserializeTo(
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
		errtype.Deserialize,
		err.Error(),
		references...)
}
