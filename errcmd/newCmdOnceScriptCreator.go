package errcmd

import (
	"bytes"
	"strings"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

type newCmdOnceScriptCreator struct{}

// Lines
//
// Join Examples:
//   - Args are joined using space (example ls /temp).
//   - ScriptLines are joined using
//     SingleLineScriptsJoiner or ScriptsMultiLineJoiner (example "cd /temp" && "ls .")
//
// Refers to
//
//	hasOutput true
//	hasSecureData false
func (it *newCmdOnceScriptCreator) Lines(
	hasOutput,
	hasSecureData bool,
	scriptType scripttype.Variant,
	scriptLines ...string,
) *CmdOnce {
	scriptDefault := scriptType.ScriptDefault()

	if !scriptDefault.IsImplemented {
		return getNotImplementedCmdOnceForScript(scriptDefault)
	}

	process, scriptLinesCompiled := it.ProcessScriptsFormat(
		scriptType,
		scriptLines...)

	return New.Create(
		hasOutput,
		hasSecureData,
		process,
		scriptLinesCompiled...,
	)
}

func (it *newCmdOnceScriptCreator) ProcessScriptsFormat(
	scriptType scripttype.Variant,
	scriptLines ...string,
) (format string, scriptLinesCompiled []string) {
	scriptDefault := scriptType.ScriptDefault()

	if !scriptDefault.IsImplemented {
		return "", nil
	}

	argsSingleLine := strings.Join(
		scriptLines,
		SingleLineScriptsJoiner)
	mergedSlice := stringslice.MergeNew(
		scriptDefault.DefaultArguments,
		argsSingleLine)

	return scriptDefault.ProcessName, mergedSlice
}

func (it *newCmdOnceScriptCreator) LinesWithStdIns(
	hasOutput,
	hasSecureData bool,
	stdIn, stdOut, stdErr *bytes.Buffer,
	scriptType scripttype.Variant,
	scriptLines ...string,
) *CmdOnce {
	scriptDefault := scriptType.ScriptDefault()

	if !scriptDefault.IsImplemented {
		return getNotImplementedCmdOnceForScript(scriptDefault)
	}

	process, scriptLinesCompiled := it.ProcessScriptsFormat(
		scriptType,
		scriptLines...)

	return New.CreateUsingStdIns(
		hasOutput,
		hasSecureData,
		stdIn, stdOut, stdErr,
		process,
		scriptLinesCompiled...,
	)
}

// LinesDefault
//
// Join Examples:
//   - Args are joined using space (example ls /temp).
//   - ScriptLines are joined using
//     SingleLineScriptsJoiner or ScriptsMultiLineJoiner (example "cd /temp" && "ls .")
//
// Refers to
//
//	hasOutput true
//	hasSecureData false
func (it *newCmdOnceScriptCreator) LinesDefault(
	scriptType scripttype.Variant,
	scriptLines ...string,
) *CmdOnce {
	return it.Lines(
		true,
		false,
		scriptType,
		scriptLines...)
}

func (it *newCmdOnceScriptCreator) ProcessArgsDefault(
	scriptType scripttype.Variant,
	process string,
	args ...string,
) *CmdOnce {
	return it.LinesDefault(
		scriptType,
		ProcessArgsJoinAppend(process, args...))
}

// ArgsDefault
//
// Join Examples:
//   - Args are joined using space (example ls /temp).
//   - ScriptLines are joined using
//     SingleLineScriptsJoiner or ScriptsMultiLineJoiner (example "cd /temp" && "ls .")
//
// Refers to
//
//	hasOutput true
//	hasSecureData false
func (it *newCmdOnceScriptCreator) ArgsDefault(
	scriptType scripttype.Variant,
	args ...string,
) *CmdOnce {
	return it.Args(
		true,
		false,
		scriptType,
		args...)
}

func (it *newCmdOnceScriptCreator) PowershellArgsDefault(
	args ...string,
) *CmdOnce {
	return it.Args(
		true,
		false,
		scripttype.Powershell,
		args...)
}

// ArgsDefaultResult
//
// Join Examples:
//   - Args are joined using space (example ls /temp).
//   - ScriptLines are joined using
//     SingleLineScriptsJoiner or ScriptsMultiLineJoiner (example "cd /temp" && "ls .")
//
// Refers to
//
//	hasOutput true
//	hasSecureData false
func (it *newCmdOnceScriptCreator) ArgsDefaultResult(
	scriptType scripttype.Variant,
	args ...string,
) *Result {
	cmdOnce := New.Script.ArgsDefault(
		scriptType,
		args...)

	result := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()
	cmdOnce = nil

	return result
}

func (it *newCmdOnceScriptCreator) ArgsDefaultLines(
	scriptType scripttype.Variant,
	args ...string,
) ([]string, *errorwrapper.Wrapper) {
	result := it.ArgsDefaultResult(scriptType, args...)

	return result.CompiledOutputLines(),
		result.CompiledFullErrorWrapper()
}

func (it *newCmdOnceScriptCreator) ArgsDefaultLinesMust(
	scriptType scripttype.Variant,
	args ...string,
) []string {
	result := it.ArgsDefaultResult(scriptType, args...)
	result.HandleError()

	return result.CompiledOutputLines()
}

func (it *newCmdOnceScriptCreator) ArgsDefaultErrorLinesMust(
	scriptType scripttype.Variant,
	args ...string,
) []string {
	result := it.ArgsDefaultResult(scriptType, args...)
	result.HandleError()

	return result.CompiledErrorLines()
}

func (it *newCmdOnceScriptCreator) ArgsDefaultOutputMust(
	scriptType scripttype.Variant,
	args ...string,
) string {
	result := it.ArgsDefaultResult(scriptType, args...)
	result.HandleError()

	return result.OutputString()
}

func (it *newCmdOnceScriptCreator) ArgsDefaultOutputBytesMust(
	scriptType scripttype.Variant,
	args ...string,
) []byte {
	result := it.ArgsDefaultResult(scriptType, args...)
	result.HandleError()

	return result.OutputBytes()
}

func (it *newCmdOnceScriptCreator) ArgsDefaultCombinedOutput(
	scriptType scripttype.Variant,
	args ...string,
) ([]byte, error) {
	scriptBuilder := New.ScriptBuilder.Default(scriptType)

	return scriptBuilder.Args(args...).CombinedOutput()
}

func (it *newCmdOnceScriptCreator) ArgsDefaultOutputWithErrorWrapper(
	scriptType scripttype.Variant,
	args ...string,
) ([]byte, *errorwrapper.Wrapper) {
	allBytes, err := it.ArgsDefaultCombinedOutput(
		scriptType,
		args...)

	return allBytes, errnew.
		Error.Type(
		errtype.FailedProcess,
		err)
}

// Args
//
// Join Examples:
//   - Args are joined using space (example ls /temp).
//   - ScriptLines are joined using
//     SingleLineScriptsJoiner or ScriptsMultiLineJoiner (example "cd /temp" && "ls .")
func (it *newCmdOnceScriptCreator) Args(
	hasOutput,
	hasSecureData bool,
	scriptType scripttype.Variant,
	args ...string,
) *CmdOnce {
	scriptDefault := scriptType.ScriptDefault()

	if !scriptDefault.IsImplemented {
		return getNotImplementedCmdOnceForScript(scriptDefault)
	}

	argsSingleLine := strings.Join(args, constants.Space)
	compiledArguments := stringslice.MergeNew(
		scriptDefault.DefaultArguments,
		argsSingleLine)

	return New.Create(
		hasOutput,
		hasSecureData,
		scriptDefault.ProcessName,
		compiledArguments...,
	)
}

// BashArgsDefault
//
// Join Examples:
//   - Args are joined using space (example ls /temp).
//   - ScriptLines are joined using
//     SingleLineScriptsJoiner or ScriptsMultiLineJoiner (example "cd /temp" && "ls .")
func (it *newCmdOnceScriptCreator) BashArgsDefault(
	args ...string,
) *CmdOnce {
	return it.ArgsDefault(
		scripttype.Bash,
		args...)
}

// ChainSuccess
//
// if any script is failed to run then don't continue others
//
// Returns true if scriptLines length 0
func (it *newCmdOnceScriptCreator) ChainSuccess(
	scriptType scripttype.Variant,
	scriptLines ...string,
) (errorWrapper *errorwrapper.Wrapper, isAllSuccess bool) {
	if len(scriptLines) == 0 {
		return nil, true
	}

	for _, scriptLine := range scriptLines {
		cmdOnce := it.Lines(
			false,
			false,
			scriptType,
			scriptLine)

		result := cmdOnce.CompiledResult()

		if result.HasAnyError() {
			return cmdOnce.CompiledResult().CompiledFullErrorWrapper(),
				false
		}

		go ClearDispose(cmdOnce)
	}

	return nil, true
}

func (it *newCmdOnceScriptCreator) IsSuccessArgs(
	scriptType scripttype.Variant,
	args ...string,
) bool {
	cmdOnce := it.Args(
		false,
		true,
		scriptType,
		args...)

	isSuccess := cmdOnce.
		IsSuccessfullyExecuted()

	go cmdOnce.Dispose()

	return isSuccess
}

func (it *newCmdOnceScriptCreator) IsSuccessLines(
	scriptType scripttype.Variant,
	scriptLines ...string,
) bool {
	cmdOnce := it.Lines(
		false,
		true,
		scriptType,
		scriptLines...)

	isSuccess := cmdOnce.
		IsSuccessfullyExecuted()
	go cmdOnce.Dispose()

	return isSuccess
}

func (it *newCmdOnceScriptCreator) IsFailedArgs(
	scriptType scripttype.Variant,
	args ...string,
) bool {
	return !it.IsSuccessArgs(
		scriptType,
		args...,
	)
}

func (it *newCmdOnceScriptCreator) Error(
	scriptType scripttype.Variant,
	args ...string,
) *errorwrapper.Wrapper {
	cmdOnce := it.Args(
		false,
		false,
		scriptType,
		args...)

	isSuccess := cmdOnce.
		IsSuccessfullyExecuted()

	if isSuccess {
		go ClearDispose(cmdOnce)

		return nil
	}

	compiledErr := cmdOnce.CompiledFullErrorWrapper()

	if compiledErr.IsEmpty() {
		go ClearDispose(cmdOnce)

		return nil
	}

	newErr := compiledErr.CloneNewStackSkipPtr(
		codestack.Skip1)
	go ClearDispose(cmdOnce)

	return newErr
}

func (it *newCmdOnceScriptCreator) BufferErrorString(
	scriptType scripttype.Variant,
	args ...string,
) string {
	cmdOnce := it.Args(
		false,
		false,
		scriptType,
		args...)

	compiledResult := cmdOnce.CompiledResult()
	errorLine := compiledResult.ErrorString()
	compiledResult.Dispose()
	go ClearDispose(cmdOnce)

	return errorLine
}

func (it *newCmdOnceScriptCreator) BufferErrorLines(
	scriptType scripttype.Variant,
	args ...string,
) []string {
	cmdOnce := it.Args(
		false,
		false,
		scriptType,
		args...)

	compiledResult := cmdOnce.CompiledResult()
	errorLines := compiledResult.CompiledErrorLines()
	copiedLines := stringslice.Clone(errorLines)
	compiledResult.Dispose()
	go ClearDispose(cmdOnce)

	return copiedLines
}

func (it *newCmdOnceScriptCreator) BufferTrimmedLines(
	scriptType scripttype.Variant,
	args ...string,
) []string {
	cmdOnce := it.Args(
		true,
		false,
		scriptType,
		args...)

	compiledResult := cmdOnce.CompiledResult()
	outputLines := compiledResult.CompiledTrimmedOutputLines()
	copiedLines := stringslice.Clone(outputLines)
	compiledResult.Dispose()
	go ClearDispose(cmdOnce)

	return copiedLines
}

func (it *newCmdOnceScriptCreator) ExitCode(
	scriptType scripttype.Variant,
	args ...string,
) int {
	cmdOnce := it.Args(
		false,
		false,
		scriptType,
		args...)

	compiledResult := cmdOnce.CompiledResult()
	exitCode := compiledResult.ExitCode
	compiledResult.Dispose()
	go ClearDispose(cmdOnce)

	return exitCode
}

func (it *newCmdOnceScriptCreator) ExitCodeByte(
	scriptType scripttype.Variant,
	args ...string,
) byte {
	cmdOnce := it.Args(
		false,
		false,
		scriptType,
		args...)

	compiledResult := cmdOnce.CompiledResult()
	exitCode := compiledResult.ExitCodeByte()
	compiledResult.Dispose()
	go ClearDispose(cmdOnce)

	return exitCode
}
