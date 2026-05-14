package errnew

import (
	"gitlab.com/evatix-go/core/coredata/corepayload"
	"gitlab.com/evatix-go/errorwrapper"
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
