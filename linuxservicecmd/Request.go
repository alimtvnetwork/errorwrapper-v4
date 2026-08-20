package linuxservicecmd

import (
	"time"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type Request struct {
	ServiceName string
	Action      servicestate.Action
}

// ExecDefault
//
// Returns error for detailed action,
// use ExecSimple or Exec for simple error
func (it Request) ExecDefault() (*Result, *errorwrapper.Wrapper) {
	return RunAction(
		true,
		false,
		it.Action,
		it.ServiceName)
}

func (it Request) Exec(
	isDetailedError,
	isIgnoreUnknownService bool,
) (*Result, *errorwrapper.Wrapper) {
	return RunAction(
		isDetailedError,
		isIgnoreUnknownService,
		it.Action,
		it.ServiceName)
}

func (it Request) ExecSimple() (*Result, *errorwrapper.Wrapper) {
	return RunAction(
		false,
		false,
		it.Action,
		it.ServiceName)
}

func (it Request) ExecError(isDetailedError bool) *errorwrapper.Wrapper {
	_, errWp := RunAction(
		isDetailedError,
		false,
		it.Action,
		it.ServiceName)

	return errWp
}

func (it Request) actionExecError(
	isDetailedError bool,
	action servicestate.Action,
) *errorwrapper.Wrapper {
	_, errWp := RunAction(
		isDetailedError,
		false,
		action,
		it.ServiceName)

	return errWp
}

func (it Request) ExecSimpleError() *errorwrapper.Wrapper {
	_, errWp := RunAction(
		false,
		false,
		it.Action,
		it.ServiceName)

	return errWp
}

func (it Request) ExecDetailed() (*Result, *errorwrapper.Wrapper) {
	return RunAction(
		true,
		false,
		it.Action,
		it.ServiceName)
}

func (it Request) ExecDetailedError() *errorwrapper.Wrapper {
	_, errWp := RunAction(
		true,
		false,
		it.Action,
		it.ServiceName)

	return errWp
}

func (it Request) StartError(isDetailed bool) *errorwrapper.Wrapper {
	return it.actionExecError(isDetailed, servicestate.Start)
}

func (it Request) RestartError(isDetailed bool) *errorwrapper.Wrapper {
	return it.actionExecError(isDetailed, servicestate.Restart)
}

func (it Request) ReloadError(isDetailed bool) *errorwrapper.Wrapper {
	return it.actionExecError(isDetailed, servicestate.Reload)
}

func (it Request) StopError(isDetailed bool) *errorwrapper.Wrapper {
	return it.actionExecError(isDetailed, servicestate.Reload)
}

func (it Request) Start() *Result {
	return Run(servicestate.Start, it.ServiceName)
}

func (it Request) Stop() *Result {
	return Run(servicestate.Stop, it.ServiceName)
}

func (it Request) Status() *Result {
	return Run(servicestate.Status, it.ServiceName)
}

func (it Request) ReferenceValues() []ref.Value {
	return []ref.Value{
		{
			Variable: "Service Name",
			Value:    it.ServiceName,
		},
		{
			Variable: "Applying Action",
			Value:    it.Action.Name(),
		},
	}
}

func (it Request) Enable() *Result {
	return Run(servicestate.Enable, it.ServiceName)
}

func (it Request) Disable() *Result {
	return Run(servicestate.Disable, it.ServiceName)
}

func (it Request) Restart() *Result {
	return Run(servicestate.Restart, it.ServiceName)
}

func (it Request) Reload() *Result {
	return Run(servicestate.Reload, it.ServiceName)
}

func (it Request) IsRunning() bool {
	return IsRunning(it.ServiceName)
}

func (it Request) IsUnknownService() bool {
	return GetStatus(it.ServiceName).IsUnknownService()
}

func (it Request) IsServiceExist() bool {
	return IsServiceExist(it.ServiceName)
}

func (it Request) StopSleepStart(
	sleep time.Duration,
) (stop, start *Result, isRunSuccess bool) {
	stop = it.Stop()

	time.Sleep(sleep)

	start = it.Start()

	isRunSuccess = stop.IsSuccess() &&
		start.IsSuccess()

	return stop, start, isRunSuccess
}

func (it Request) Run() *Result {
	return Run(it.Action, it.ServiceName)
}

func (it Request) SimplifiedErrorWrapper() *errorwrapper.Wrapper {
	return it.Run().SimplifiedError()
}

func (it Request) VerifyExitCode(exitCode linuxservicestate.ExitCode) *errorwrapper.Wrapper {
	return it.Run().VerifyExitCode(exitCode)
}

func (it Request) CompiledErrorWrapper() *errorwrapper.Wrapper {
	return it.Run().CompiledErrorWrapper()
}

func (it *Request) Dispose() {
	if it == nil {
		return
	}

	it.ServiceName = constants.EmptyString
	it.Action = servicestate.Status
}
