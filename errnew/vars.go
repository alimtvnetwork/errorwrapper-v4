package errnew

import (
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

var (
	Messages                      = newMessagesToErrorWrapperCreator{}
	Message                       = newMessageToErrorWrapperCreator{}
	Enum                          = newEnumToErrorWrapperCreator{}
	Ref                           = newRefToErrorWrapperCreator{}
	Refs                          = newReferencesToErrorWrapperCreator{}
	Path                          = newPathToErrorWrapperCreator{}
	Merge                         = newMergeToErrorWrapperCreator{}
	FromTo                        = newFromToErrorWrapperCreator{}
	Null                          = newNullToErrorWrapperCreator{}
	Type                          = newTypeToWrapperCreator{}
	Error                         = newErrorToWrapperCreator{}
	Range                         = newRangeWrapperCreator{}
	Unmarshal                     = newUnmarshalWrapperCreator{} // refers to unmarshal related quick errors
	Fmt                           = newFormatterCreator{}
	MessageWithRef                = newMessageWithRefCreator{}
	Json                          = newJsonToWrapperCreator{}
	Payload                       = newPayloadToErrorWrapperCreator{}
	Reflect                       = newReflectErrToWrapperCreator{}
	DeserializeTo                 = newDeserializeToWrapperCreator{} // actually helps to deserialize to something
	ErrInterface                  = newErrorInterfaceToWrapperCreator{}
	SrcDst                        = newSourceDestinationToErrorWrapperCreator{}
	MappingFailed                 = Messages.Single(errtype.MappingFailed, "Cannot map the given data to expected data format from the dictionary or map.")
	InvalidSystemUser             = Messages.Single(errtype.SysUserInvalid, "System user not found or id has issues.")
	InvalidSystemGroup            = Messages.Single(errtype.SysGroupInvalid, "System group not found or id has issues.")
	FinalizedResourceCannotAccess = Type.Default(errtype.FinalizedResourceCannotAccess)
	NotFound                      = newNotFoundErrCreator{}
	EmptyString                   = Type.Default(errtype.EmptyString)
	Unexpected                    = Type.Default(errtype.Unexpected)
	InvalidInput                  = Type.Default(errtype.InvalidInput)
	Invalid                       = Type.Default(errtype.Invalid)
	InvalidOption                 = Type.Default(errtype.InvalidOption)
)
