package errcmd

import (
	"bytes"
	"os/exec"
	"strings"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type newCreator struct {
	ScriptBuilder       newCmdOnceScriptBuilder
	OsTypeScriptBuilder newOsTypeScriptBuilderCreator
	Script              newCmdOnceScriptCreator
	ShellScript         *newCmdOnceTypedScriptsCreator
	BashScript          *newCmdOnceTypedScriptsCreator
	OutputGetter        newOutputGetter
	NoOutput            newCmdOnceNoOutputCreator
	Cmd                 newCmdCreator
}

// Output (alias for Default)
//
// Create cmd once for
//
//  - hasOutput true
//  - hasSecureData false
func (it newCreator) Output(
	process string,
	arguments ...string,
) *CmdOnce {
	return it.Create(
		true,
		false,
		process,
		arguments...)
}

// Default (alias for Output)
//
// Create cmd once for
//
//  - hasOutput true
//  - hasSecureData false
func (it newCreator) Default(
	process string,
	arguments ...string,
) *CmdOnce {
	return it.Create(
		true,
		false,
		process,
		arguments...)
}

func (it newCreator) UsingCmd(
	hasOutput,
	hasSecureData bool,
	cmd *exec.Cmd,
) *CmdOnce {
	isEmptyCmd := cmd == nil

	if isEmptyCmd {
		return &CmdOnce{
			baseCmdWrapper: BaseCmdWrapper{
				baseBufferStdOutError: &baseBufferStdOutError{},
				initializeErrorWrapper: errnew.Messages.SingleUsingStackSkip(
					codestack.Skip1,
					errtype.CommandExecutionNotFound,
					errcore.FailedToCreateCmdType.String()),
			},
			Cmd:            cmd,
			hasCmd:         false,
			hasOutput:      hasOutput,
			hasSecureData:  hasSecureData,
			compiledResult: nil,
		}
	}

	argumentsSingle := strings.Join(
		cmd.Args,
		constants.Space)

	hasCmd := !isEmptyCmd
	var errWrapper *errorwrapper.Wrapper
	var stdOut, stdErr *bytes.Buffer
	process := cmd.Path
	wholeCmdLine := process +
		constants.Space +
		argumentsSingle

	if isEmptyCmd && hasSecureData {
		errWrapper = errnew.
			Messages.
			SingleUsingStackSkip(
				codestack.Skip1,
				errtype.InvalidProcess,
				process)
	} else if isEmptyCmd && !hasSecureData {
		errWrapper = errnew.Messages.
			SingleUsingStackSkip(
				codestack.Skip1,
				errtype.InvalidProcess,
				wholeCmdLine)
	}

	if hasCmd && !hasSecureData {
		stdErr = &bytes.Buffer{}
		cmd.Stderr = stdErr
	}

	if hasCmd && hasOutput {
		stdOut = &bytes.Buffer{}
		stdOut.Grow(constants.Capacity1G)
		cmd.Stdout = stdOut
	}

	return &CmdOnce{
		baseCmdWrapper: BaseCmdWrapper{
			baseBufferStdOutError: &baseBufferStdOutError{
				stdErr: stdErr,
				stdOut: stdOut,
			},
			initializeErrorWrapper: errWrapper,
			process:                process,
			wholeCommand:           wholeCmdLine,
			argumentsSingleLine:    argumentsSingle,
			arguments:              cmd.Args,
		},
		Cmd:           cmd,
		hasCmd:        hasCmd,
		hasOutput:     hasOutput,
		hasSecureData: hasSecureData,
	}
}

func (it newCreator) Create(
	hasOutput,
	hasSecureData bool,
	process string,
	arguments ...string,
) *CmdOnce {
	cmd := exec.Command(process, arguments...)
	argumentsSingle := strings.Join(
		arguments,
		constants.Space)
	isEmptyCmd := cmd == nil
	hasCmd := !isEmptyCmd
	var errWrapper *errorwrapper.Wrapper
	var stdOut, stdErr *bytes.Buffer
	wholeCmdLine := process +
		constants.Space +
		argumentsSingle

	if isEmptyCmd && hasSecureData {
		errWrapper = errnew.
			Messages.
			SingleUsingStackSkip(
				codestack.Skip1,
				errtype.InvalidProcess,
				process)
	} else if isEmptyCmd && !hasSecureData {
		errWrapper = errnew.
			Messages.
			SingleUsingStackSkip(
				codestack.Skip1,
				errtype.InvalidProcess,
				wholeCmdLine)
	}

	if hasCmd && !hasSecureData {
		stdErr = &bytes.Buffer{}
		cmd.Stderr = stdErr
	}

	if hasCmd && hasOutput {
		stdOut = &bytes.Buffer{}
		stdOut.Grow(constants.Capacity1G)
		cmd.Stdout = stdOut
	}

	return &CmdOnce{
		baseCmdWrapper: BaseCmdWrapper{
			baseBufferStdOutError: &baseBufferStdOutError{
				stdErr: stdErr,
				stdOut: stdOut,
			},
			initializeErrorWrapper: errWrapper,
			process:                process,
			wholeCommand:           wholeCmdLine,
			argumentsSingleLine:    argumentsSingle,
			arguments:              arguments,
		},
		Cmd:           cmd,
		hasCmd:        hasCmd,
		hasOutput:     hasOutput,
		hasSecureData: hasSecureData,
	}
}

// CreateUsingStdIns
//
//  stdIn, stdOut, stdErr will override if given, nill will be ignored.
func (it newCreator) CreateUsingStdIns(
	hasOutput,
	hasSecureData bool,
	stdIn, stdOut, stdErr *bytes.Buffer,
	process string,
	arguments ...string,
) *CmdOnce {
	cmd := exec.Command(
		process,
		arguments...)
	argumentsSingle := strings.Join(
		arguments,
		constants.Space)
	isEmptyCmd := cmd == nil
	hasCmd := !isEmptyCmd
	var errWrapper *errorwrapper.Wrapper
	wholeCmdLine := process +
		constants.Space +
		argumentsSingle

	if isEmptyCmd && hasSecureData {
		errWrapper = errnew.
			Messages.
			SingleUsingStackSkip(
				codestack.Skip1,
				errtype.InvalidProcess,
				process)
	} else if isEmptyCmd && !hasSecureData {
		errWrapper = errnew.
			Messages.
			SingleUsingStackSkip(
				codestack.Skip1,
				errtype.InvalidProcess,
				wholeCmdLine)
	}

	if stdErr != nil {
		cmd.Stderr = stdErr
	}

	if stdErr == nil {
		stdErr = &bytes.Buffer{}
	}

	hasOutput = hasOutput || stdOut != nil

	if stdOut != nil {
		cmd.Stdout = stdOut
	} else if stdOut == nil && hasCmd && hasOutput {
		writer := &bytes.Buffer{}
		writer.Grow(constants.Capacity1G)
		stdOut = writer
		cmd.Stdout = stdOut
	}

	if hasCmd && !hasSecureData {
		cmd.Stderr = stdErr
	}

	if stdIn != nil {
		cmd.Stdin = stdIn
	}

	return &CmdOnce{
		baseCmdWrapper: BaseCmdWrapper{
			baseBufferStdOutError: &baseBufferStdOutError{
				stdErr: stdErr,
				stdOut: stdOut,
			},
			initializeErrorWrapper: errWrapper,
			process:                process,
			wholeCommand:           wholeCmdLine,
			argumentsSingleLine:    argumentsSingle,
			arguments:              arguments,
		},
		Cmd:           cmd,
		hasCmd:        hasCmd,
		hasOutput:     hasOutput,
		hasSecureData: hasSecureData,
	}
}

func (it newCreator) Empty() *CmdOnce {
	return &CmdOnce{
		baseCmdWrapper: BaseCmdWrapper{
			initializeErrorWrapper: cmdNilErr,
		},
		hasOutput:     true,
		hasSecureData: true,
	}
}

func (it newCreator) Result(
	process string,
	args ...string,
) *Result {
	cmdOnce := New.Default(
		process,
		args...)

	result := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()
	cmdOnce = nil

	return result
}

func (it newCreator) CombinedOutput(
	process string,
	args ...string,
) ([]byte, error) {
	cmd := exec.Command(
		process,
		args...)

	if cmd == nil {
		return nil, errtype.CommandExecutionNotFound.ReferencesCsvError(
			"cmd not found",
			process,
			ArgsJoin(args...))
	}

	return cmd.CombinedOutput()
}

func (it newCreator) CombinedOutputRaw(
	process string,
	args ...string,
) []byte {
	cmd := exec.Command(
		process,
		args...)

	if cmd == nil {
		return []byte("no cmd found for `" + process + "` " + ArgsJoin(args...))
	}

	rawBytes, err := cmd.CombinedOutput()
	hasBytes := len(rawBytes) > 0

	if hasBytes && err != nil {
		return []byte(string(rawBytes) + err.Error())
	} else if !hasBytes && err != nil {
		return []byte(err.Error())
	}

	return rawBytes
}

func (it newCreator) CombinedOutputString(
	process string,
	args ...string,
) string {
	rawBytes := it.CombinedOutputRaw(process, args...)

	return string(rawBytes)
}

func (it newCreator) CompiledError(
	process string,
	args ...string,
) *errorwrapper.Wrapper {
	cmdOnce := New.Default(
		process,
		args...)

	result := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()
	cmdOnce = nil

	return result.CompiledFullErrorWrapper()
}

func (it newCreator) ErrorLines(
	process string,
	args ...string,
) []string {
	cmdOnce := New.Default(
		process,
		args...)

	result := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()
	cmdOnce = nil

	return result.CompiledErrorLines()
}

func (it newCreator) TrimmedErrorLines(
	process string,
	args ...string,
) []string {
	cmdOnce := New.Default(
		process,
		args...)

	result := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()
	cmdOnce = nil

	return result.CompiledTrimmedErrorLines()
}

func (it newCreator) TrimmedOutputLines(
	process string,
	args ...string,
) []string {
	cmdOnce := New.Default(
		process,
		args...)

	result := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()
	cmdOnce = nil

	return result.CompiledTrimmedOutputLines()
}

func (it newCreator) IsSuccess(
	process string,
	args ...string,
) bool {
	return it.
		NoOutput.
		Default(process, args...).
		IsSuccessfullyExecuted()
}

func (it newCreator) IsFailed(
	process string,
	args ...string,
) bool {
	return !it.
		NoOutput.
		Default(process, args...).
		IsSuccessfullyExecuted()
}

func (it newCreator) ExitCode(
	process string,
	args ...string,
) int {
	return it.NoOutput.ExitCode(
		process,
		args...)
}

func (it newCreator) ExitCodeAsByte(
	process string,
	args ...string,
) byte {
	return it.NoOutput.ExitCodeAsByte(
		process,
		args...)
}
