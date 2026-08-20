package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

type StateValidateInstruction struct {
	ServiceName      string
	ExpectedExitCode linuxservicestate.ExitCode
}

func (it *StateValidateInstruction) IsEmpty() bool {
	return it == nil
}

func (it *StateValidateInstruction) Apply() *errorwrapper.Wrapper {
	if it.IsEmpty() {
		return nil
	}

	return VerifyExitCode(it.ServiceName, it.ExpectedExitCode)
}

func (it *StateValidateInstruction) ApplyUsingErrCollection(
	errCollection *errwrappers.Collection,
) (isSuccess bool) {
	if it.IsEmpty() {
		return true
	}

	err := it.Apply()
	errCollection.AddWrapperPtr(err)

	return err.IsSuccess()
}
