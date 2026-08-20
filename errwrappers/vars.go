package errwrappers

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

var (
	defaultErrorType       = errtype.Unknown
	Deserialize            = newDeserializeCreator{} // Deserialize from payload to error collection
	ErrorInterface         = newErrorInterfaceToCollection{}
	usingBytesDeserializer = corejson.Deserialize.UsingBytes
	serializer             = corejson.Serialize.ToBytesErr
)
