package errcmd

import (
	"os/exec"

	"github.com/alimtvnetwork/core-v9/coredata/corestr"
)

type cmdCompiledOutput struct {
	Cmd         *exec.Cmd
	Output      corestr.SimpleStringOnce
	Error       error
	ProcessName string
	Arguments   []string
}

func InvalidCmdCompiledOutput(
	err error,
	process string,
	args []string,
) *cmdCompiledOutput {
	return &cmdCompiledOutput{
		Error:       err,
		ProcessName: process,
		Arguments:   args,
	}
}

func (it *cmdCompiledOutput) CloneCmd() *exec.Cmd {
	return exec.Command(it.ProcessName, it.Arguments...)
}

func (it cmdCompiledOutput) Clone() cmdCompiledOutput {
	return *CombinedOutputError(
		nil,
		it.ProcessName,
		it.Arguments...)
}

func (it *cmdCompiledOutput) ClonePtr() *cmdCompiledOutput {
	if it == nil {
		return nil
	}

	return CombinedOutputError(
		nil,
		it.ProcessName,
		it.Arguments...)
}

func (it *cmdCompiledOutput) Dispose() {
	if it == nil {
		return
	}

	it.Cmd.Args = nil
	it.Cmd = nil
	it.Output.Dispose()
	it.Error = nil
}
