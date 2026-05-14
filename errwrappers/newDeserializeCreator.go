package errwrappers

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/serializerinf"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

type newDeserializeCreator struct{}

func (it newDeserializeCreator) UsingBytes(
	jsonBytes []byte,
) (errCollection *Collection, parsedErrorWrapper *errorwrapper.Wrapper) {
	if len(jsonBytes) == 0 {
		return nil, nil
	}

	empty := Empty()
	errWp := errnew.
		DeserializeTo.
		BytesToAnyPtr(jsonBytes, empty)

	return empty, errWp
}

func (it newDeserializeCreator) UsingJsonResult(
	jsonResult *corejson.Result,
) (errCollection *Collection, parsedErrorWrapper *errorwrapper.Wrapper) {
	empty := Empty()
	errWp := errnew.
		DeserializeTo.
		JsonResultToAnySkipOnNull(
			jsonResult, empty)

	return empty, errWp
}

func (it newDeserializeCreator) UsingString(
	jsonString string,
) (*Collection, *errorwrapper.Wrapper) {
	empty := Empty()

	if jsonString == "" {
		return empty, nil
	}

	errWp := errnew.
		DeserializeTo.
		BytesToAnyPtr([]byte(jsonString), empty)

	return empty, errWp
}

// UsingError
//
//  here the error actually contains the json payload.
//  which will be unmarshalled and created to error collection.
func (it newDeserializeCreator) UsingError(
	errorAsJsonString error,
) (*Collection, *errorwrapper.Wrapper) {
	empty := Empty()

	if errorAsJsonString == nil {
		return empty, nil
	}

	toString := errcore.ToString(
		errorAsJsonString)

	errWp := errnew.
		DeserializeTo.
		BytesToAnyPtr([]byte(toString), empty)

	return empty, errWp
}

func (it newDeserializeCreator) UsingSerializer(
	serializer serializerinf.Serializer,
) (*Collection, *errorwrapper.Wrapper) {
	empty := Empty()

	if serializer == nil {
		return empty, nil
	}

	allBytes, err := serializer.Serialize()
	jsonResult := corejson.
		NewResult.Ptr(
		allBytes,
		err,
		coredynamic.TypeName(empty))

	return it.UsingJsonResult(jsonResult)
}
