package linuxservicecmd

import "github.com/alimtvnetwork/errorwrapper-v3/errwrappers"

type StateValidateInstructions struct {
	IsContinueOnError bool
	Validations       []StateValidateInstruction `json:"Validations,omitempty"`
}

func (it *StateValidateInstructions) IsEmpty() bool {
	return it == nil || len(it.Validations) == 0
}

func (it *StateValidateInstructions) HasValidations() bool {
	return it != nil && len(it.Validations) > 0
}

func (it *StateValidateInstructions) ApplyUsingErrCollection(
	errCollection *errwrappers.Collection,
) (isSuccess bool) {
	if it.IsEmpty() {
		return true
	}

	stateTracker := errCollection.StateTracker()

	for _, validation := range it.Validations {
		err := validation.Apply()
		errCollection.AddWrapperPtr(err)

		if !it.IsContinueOnError && err.HasError() {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
