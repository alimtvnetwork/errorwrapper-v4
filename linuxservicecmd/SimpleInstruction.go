package linuxservicecmd

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

type SimpleInstruction struct {
	ServiceName     string
	Action          servicestate.Action
	IsDetailedError bool
	lazyResult      *Result
	lazyStatus      linuxservicestate.ExitCode
}

func (it *SimpleInstruction) LazyStatus() linuxservicestate.ExitCode {
	if it.lazyStatus.IsDefined() {
		return it.lazyStatus
	}

	exitCode := it.Status()
	it.lazyStatus = exitCode

	return it.lazyStatus
}

func (it *SimpleInstruction) CommandLine() string {
	return it.LazyRun().CmdOnce.CommandLine()
}

func (it *SimpleInstruction) DetailedOutput() string {
	result := it.LazyRun()
	cmdOnce := result.CmdOnce
	slice := corestr.New.SimpleSlice.Cap(constants.Capacity8)
	slice.Add("CommandLine : " + cmdOnce.CommandLine())
	slice.Add("Exit Code or Status : " + result.ExitCode.Name())
	slice.Add("Output : ")
	slice.Add(cmdOnce.CompiledResult().DetailedOutput())

	return slice.String()
}

func (it *SimpleInstruction) OutputLines() []string {
	return it.LazyRun().CmdOnce.CompiledOutputLines()
}

func (it *SimpleInstruction) LazyRun() *Result {
	if it.lazyResult != nil {
		return it.lazyResult
	}

	it.lazyResult = Run(it.Action, it.ServiceName)

	return it.lazyResult
}

func (it *SimpleInstruction) Run() *Result {
	return Run(it.Action, it.ServiceName)
}

func (it *SimpleInstruction) Status() linuxservicestate.ExitCode {
	return GetStatus(it.ServiceName)
}

func (it *SimpleInstruction) Apply() *errorwrapper.Wrapper {
	return it.LazyRun().CompiledErrorWrapper()
}

func (it *SimpleInstruction) CompiledError() *errorwrapper.Wrapper {
	return it.LazyRun().CompiledErrorWrapper()
}

func (it *SimpleInstruction) ApplyOnErrorCollection(
	errCollection *errwrappers.Collection,
) (isSuccess bool) {
	err := it.Apply()

	errCollection.AddWrapperPtr(err)

	return err.IsSuccess()
}
