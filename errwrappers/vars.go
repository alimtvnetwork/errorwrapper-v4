package errwrappers

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

var (
	defaultErrorType       = errtype.Unknown
	Deserialize            = newDeserializeCreator{} // Deserialize from payload to error collection
	ErrorInterface         = newErrorInterfaceToCollection{}
	usingBytesDeserializer = corejson.Deserialize.UsingBytes
	serializer             = corejson.Serialize.ToBytesErr
)
