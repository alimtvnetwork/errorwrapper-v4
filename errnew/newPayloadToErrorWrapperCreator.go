package errnew

import (
	"github.com/alimtvnetwork/core-v9/coredata/corepayload"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

type newPayloadToErrorWrapperCreator struct{}

func (it newPayloadToErrorWrapperCreator) Create(
	payloadWrapper *corepayload.PayloadWrapper,
) *errorwrapper.Wrapper {
	if payloadWrapper.IsEmptyError() {
		return nil
	}

	return errorwrapper.NewUsingBasicErr(
		payloadWrapper.
			Attributes.
			BasicErrWrapper)
}
