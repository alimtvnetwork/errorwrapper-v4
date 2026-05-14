package errwrappers

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface/serializerinf"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
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
		coredynamic.SafeTypeName(empty))

	return it.UsingJsonResult(jsonResult)
}
