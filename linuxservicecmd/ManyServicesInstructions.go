package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

type ManyServicesInstructions struct {
	IsDetailedError   bool                      `json:"IsDetailedError,omitempty"`
	IsContinueOnError bool                      `json:"IsContinueOnError,omitempty"`
	Instructions      []CoreServicesInstruction `json:"Instructions,omitempty"`
}

func (it *ManyServicesInstructions) IsEmpty() bool {
	return it == nil || len(it.Instructions) == 0
}

func (it *ManyServicesInstructions) Apply(
	errCollection *errwrappers.Collection,
) (isSuccess bool) {
	if it.IsEmpty() {
		return true
	}

	stateTracker := errCollection.StateTracker()

	for _, coreServicesInstruction := range it.Instructions {
		isSuccessInner := RunServices(
			it.IsDetailedError,
			it.IsContinueOnError,
			coreServicesInstruction.IsIgnoreUnknownService,
			errCollection,
			coreServicesInstruction.Action,
			coreServicesInstruction.ServicesNames...)

		if !isSuccessInner && !it.IsContinueOnError {
			return false
		}
	}

	return stateTracker.IsSuccess()
}

func (it *ManyServicesInstructions) Status(
	errCollection *errwrappers.Collection,
) (isSuccess bool) {
	return it.Run(errCollection, servicestate.Status)
}

func (it *ManyServicesInstructions) Run(
	errCollection *errwrappers.Collection,
	action servicestate.Action,
) (isSuccess bool) {
	if it.IsEmpty() {
		return true
	}

	stateTracker := errCollection.StateTracker()

	for _, coreServicesInstruction := range it.Instructions {
		isSuccessInner := RunServices(
			it.IsDetailedError,
			it.IsContinueOnError,
			coreServicesInstruction.IsIgnoreUnknownService,
			errCollection,
			action,
			coreServicesInstruction.ServicesNames...)

		if !isSuccessInner && !it.IsContinueOnError {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
