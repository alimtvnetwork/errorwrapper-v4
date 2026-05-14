package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/enum-v10/servicestate"
)

type Results struct {
	Items []*Result
}

func NewResults(capacity int) *Results {
	rs := make([]*Result, 0, capacity)

	return &Results{Items: rs}
}

func EmptyResults() *Results {
	return NewResults(0)
}

func (it *Results) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *Results) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Results) Add(action servicestate.Action, serviceName string) linuxservicestate.ExitCode {
	rs := Run(action, serviceName)
	it.Items = append(it.Items, rs)

	return rs.ExitCode
}

func (it *Results) AddServices(
	isContinueOnError bool,
	action servicestate.Action,
	servicesNames ...string,
) *Results {
	for _, s := range servicesNames {
		exitCode := it.Add(action, s)

		if !isContinueOnError && exitCode.IsFailed() {
			return it
		}
	}

	return it
}

func (it *Results) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Results) HasAnyFailed() bool {
	for _, item := range it.Items {
		if !item.ExitCode.IsActiveRunning() {
			return true
		}
	}

	return false
}

func (it *Results) IsAllSuccess() bool {
	return !it.HasAnyFailed()
}
