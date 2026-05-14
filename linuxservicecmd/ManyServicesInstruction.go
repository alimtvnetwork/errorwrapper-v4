package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

type ManyServicesInstruction struct {
	IsDetailedError   bool `json:"IsDetailedError,omitempty"`
	IsContinueOnError bool `json:"IsContinueOnError,omitempty"`
	CoreServicesInstruction
}

func (it *ManyServicesInstruction) IsEmpty() bool {
	return it == nil || len(it.ServicesNames) == 0
}

func (it *ManyServicesInstruction) Results(errCollection *errwrappers.Collection) *Results {
	return RunServicesResults(
		it.IsContinueOnError,
		it.IsDetailedError,
		it.IsIgnoreUnknownService,
		errCollection,
		it.Action,
		it.ServicesNames...)
}

func (it *ManyServicesInstruction) Apply(errCollection *errwrappers.Collection) (isSuccess bool) {
	if it.IsEmpty() {
		return true
	}

	return RunServices(
		it.IsDetailedError,
		it.IsContinueOnError,
		it.IsIgnoreUnknownService,
		errCollection,
		it.Action,
		it.ServicesNames...)
}

func (it *ManyServicesInstruction) Status(errCollection *errwrappers.Collection) (isSuccess bool) {
	return RunServices(
		it.IsDetailedError,
		it.IsContinueOnError,
		it.IsIgnoreUnknownService,
		errCollection,
		servicestate.Status,
		it.ServicesNames...)
}
