package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

type ServicesInstruction struct {
	IsDetailedError      bool                       `json:"IsDetailedError,omitempty"`
	IsContinueOnError    bool                       `json:"IsContinueOnError,omitempty"`
	StopServicesNames    []string                   `json:"StopServicesNames,omitempty"`
	StartServicesNames   []string                   `json:"StartServicesNames,omitempty"`
	RestartServicesNames []string                   `json:"RestartServicesNames,omitempty"`
	DisableServicesNames []string                   `json:"DisableServicesNames,omitempty"`
	EnableServicesNames  []string                   `json:"EnableServicesNames,omitempty"`
	StatusServicesNames  []string                   `json:"StatusServicesNames,omitempty"`
	Services             ManyServicesInstructions   `json:"Instructions,omitempty"`
	Validations          *StateValidateInstructions `json:"Validations,omitempty"`
}

func (it *ServicesInstruction) HasStopServicesNames() bool {
	return it != nil && len(it.StopServicesNames) > 0
}

func (it *ServicesInstruction) HasStartServicesNames() bool {
	return it != nil && len(it.StartServicesNames) > 0
}

func (it *ServicesInstruction) HasRestartServicesNames() bool {
	return it != nil && len(it.RestartServicesNames) > 0
}

func (it *ServicesInstruction) HasEnableServicesNames() bool {
	return it != nil && len(it.EnableServicesNames) > 0
}

func (it *ServicesInstruction) HasDisableServicesNames() bool {
	return it != nil && len(it.DisableServicesNames) > 0
}

func (it *ServicesInstruction) HasStatusServicesNames() bool {
	return it != nil && len(it.StatusServicesNames) > 0
}

func (it *ServicesInstruction) HasServices() bool {
	return it != nil && !it.Services.IsEmpty()
}

func (it *ServicesInstruction) HasValidations() bool {
	return it != nil && it.Validations.HasValidations()
}

func (it *ServicesInstruction) applyByActions(
	errCollection *errwrappers.Collection,
	action servicestate.Action,
	names []string,
) {
	RunServices(
		it.IsDetailedError,
		it.IsContinueOnError,
		false,
		errCollection,
		action,
		names...)
}

func (it *ServicesInstruction) Apply(errCollection *errwrappers.Collection) (isSuccess bool) {
	if it == nil {
		return true
	}

	stateTracker := errCollection.StateTracker()
	isImmediateExit := !it.IsContinueOnError

	// stop
	if it.HasStopServicesNames() {
		it.applyByActions(
			errCollection,
			servicestate.Stop,
			it.StopServicesNames)
	}

	if isImmediateExit && stateTracker.IsFailed() {
		return false
	}

	// start
	if it.HasStartServicesNames() {
		it.applyByActions(
			errCollection,
			servicestate.Start,
			it.StartServicesNames)
	}

	if isImmediateExit && stateTracker.IsFailed() {
		return false
	}

	// restart
	if it.HasRestartServicesNames() {
		it.applyByActions(
			errCollection,
			servicestate.Restart,
			it.RestartServicesNames)
	}

	if isImmediateExit && stateTracker.IsFailed() {
		return false
	}

	// disable
	if it.HasDisableServicesNames() {
		it.applyByActions(
			errCollection,
			servicestate.Disable,
			it.DisableServicesNames)
	}

	if isImmediateExit && stateTracker.IsFailed() {
		return false
	}

	// enable
	if it.HasEnableServicesNames() {
		it.applyByActions(
			errCollection,
			servicestate.Enable,
			it.EnableServicesNames)
	}

	if isImmediateExit && stateTracker.IsFailed() {
		return false
	}

	// status
	if it.HasStatusServicesNames() {
		it.applyByActions(
			errCollection,
			servicestate.Status,
			it.StatusServicesNames)
	}

	if isImmediateExit && stateTracker.IsFailed() {
		return false
	}

	// services
	if it.HasServices() {
		it.Services.IsDetailedError = it.IsDetailedError || it.Services.IsDetailedError
		it.Services.IsContinueOnError = it.IsContinueOnError || it.Services.IsContinueOnError
		it.Services.Apply(errCollection)
	}

	if isImmediateExit && stateTracker.IsFailed() {
		return false
	}

	// validations
	if it.HasValidations() {
		it.Validations.IsContinueOnError = it.IsContinueOnError ||
			it.Validations.IsContinueOnError
		it.Validations.
			ApplyUsingErrCollection(errCollection)
	}

	return stateTracker.IsSuccess()
}
