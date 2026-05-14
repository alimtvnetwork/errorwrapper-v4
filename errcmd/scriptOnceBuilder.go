package errcmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type scriptOnceBuilder struct {
	scriptType                                       scripttype.Variant
	scriptLines                                      *corestr.SimpleSlice
	envVars                                          *corestr.KeyValueCollection
	isSecure, hasOutput, isSetCurrentEnv, isDisposed bool
	stdIn, stdOut, stderr                            *bytes.Buffer
}

func (it *scriptOnceBuilder) IsNull() bool {
	return it == nil
}

func (it *scriptOnceBuilder) IsNotNull() bool {
	return it != nil
}

func (it *scriptOnceBuilder) IsAnyNull() bool {
	return it == nil || it.scriptLines == nil
}

func (it *scriptOnceBuilder) IsValidWithoutItems() bool {
	return it != nil && it.scriptLines.IsEmpty()
}

func (it *scriptOnceBuilder) initializeEnvVars() *corestr.KeyValueCollection {
	if it == nil {
		panic("cannot initialize on null, script builder")
	}

	if it.envVars != nil {
		return it.envVars
	}

	it.envVars = corestr.New.KeyValues.Cap(constants.Capacity5)

	return it.envVars
}

func (it *scriptOnceBuilder) SetCurrentEnvPlus(
	envValues ...corestr.KeyValuePair,
) ScriptOnceBuilder {
	if it == nil {
		return nil
	}

	it.initializeEnvVars().Adds(envValues...)

	return it
}

func (it *scriptOnceBuilder) SetCurrentEnv() ScriptOnceBuilder {
	it.isSetCurrentEnv = true

	return it
}

func (it *scriptOnceBuilder) SetCurrentEnvPlusIf(
	isCondition bool,
	envValues ...corestr.KeyValuePair,
) ScriptOnceBuilder {
	if it == nil || !isCondition {
		return nil
	}

	it.initializeEnvVars().Adds(envValues...)

	return it
}

func (it *scriptOnceBuilder) SetCurrentEnvPlusMap(
	envMap map[string]string,
) ScriptOnceBuilder {
	if it == nil {
		return nil
	}

	it.initializeEnvVars().AddMap(envMap)

	return it
}

func (it *scriptOnceBuilder) EnvVars() *corestr.KeyValueCollection {
	return it.envVars
}

func (it *scriptOnceBuilder) HasAnyEnvVars() bool {
	return it != nil && it.envVars.HasAnyItem()
}

func (it *scriptOnceBuilder) IsCurrentEnvSet() bool {
	return it != nil && it.isSetCurrentEnv
}

func (it *scriptOnceBuilder) NewBuilderWithArgs(
	args ...string,
) ScriptOnceBuilder {
	cloned := it.Clone()

	return cloned.Args(args...)
}

func (it *scriptOnceBuilder) Clone() ScriptOnceBuilder {
	if it == nil {
		return nil
	}

	return &scriptOnceBuilder{
		scriptType:  it.scriptType,
		scriptLines: it.scriptLines.ClonePtr(true),
		isSecure:    it.isSecure,
		hasOutput:   it.hasOutput,
		isDisposed:  it.isDisposed,
		stdIn:       BufferClone(it.stdIn),
		stdOut:      BufferClone(it.stdOut),
		stderr:      BufferClone(it.stderr),
	}
}

func (it *scriptOnceBuilder) OutputLinesSimpleSlice() *corestr.SimpleSlice {
	return corestr.New.SimpleSlice.Create(it.OutputLines())
}

func (it *scriptOnceBuilder) HasPlainOutput() bool {
	return it.hasOutput && !it.isSecure
}

func (it *scriptOnceBuilder) HasInsecureOutput() bool {
	return !it.isSecure
}

func (it *scriptOnceBuilder) AsyncStartErrorWrapper() (*exec.Cmd, *errorwrapper.Wrapper) {
	cmd, err := it.AsyncStart()
	var combinedErr string

	if err != nil && it.HasInsecureOutput() {
		combinedErr = CmdToScriptLine(cmd)
	}

	return cmd, errnew.Error.TypeMsg(
		errtype.FailedProcess,
		err,
		combinedErr)
}

func (it *scriptOnceBuilder) BuildCmdClear() *exec.Cmd {
	cmd := it.BuildCmd()
	it.Clear()

	return cmd
}

// IsDefined
//
//  return it != nil &&
//		it.scriptType.IsValid() &&
//		!it.HasError() &&
//		it.scriptLines.HasAnyItem()
func (it *scriptOnceBuilder) IsDefined() bool {
	return it.IsValid()
}

func (it *scriptOnceBuilder) HasStdIn() bool {
	return it != nil && it.stdIn != nil
}

func (it *scriptOnceBuilder) HasStdOut() bool {
	return it != nil && it.stdOut != nil
}

func (it *scriptOnceBuilder) HasStdErr() bool {
	return it != nil && it.stderr != nil
}

// AsyncStart
//
//  Starts cmd in async manner,
//  have to wait to be waited on
//
//  Check out exec.Cmd Start()
func (it *scriptOnceBuilder) AsyncStart() (*exec.Cmd, error) {
	cmd := it.StandardOutputCmd()
	err := cmd.Start()

	return cmd, err
}

func (it *scriptOnceBuilder) CombinedOutput() ([]byte, error) {
	cmd := it.StandardOutputCmd()
	cmd.Stderr = nil
	cmd.Stdout = nil

	return cmd.CombinedOutput()
}

func (it *scriptOnceBuilder) CombinedOutputErrorWrapper() ([]byte, *errorwrapper.Wrapper) {
	cmd := it.BuildCmd()
	cmd.Stderr = nil
	cmd.Stdout = nil

	allBytes, err := cmd.CombinedOutput()

	var combinedErr string

	if err != nil && it.HasInsecureOutput() {
		combinedErr = CmdToScriptLine(cmd)
	}

	return allBytes, errnew.Error.TypeMsg(
		errtype.FailedProcess,
		err,
		combinedErr)
}

func (it *scriptOnceBuilder) ScriptLines() corestr.SimpleSlice {
	return it.scriptLines.NonPtr()
}

func (it *scriptOnceBuilder) Append(
	appendItems ...ScriptOnceBuilder,
) ScriptOnceBuilder {
	if len(appendItems) == 0 {
		return it
	}

	for _, scriptBuilder := range appendItems {
		if scriptBuilder == nil || scriptBuilder.IsEmpty() {
			continue
		}

		it.scriptLines.Adds(
			([]string)(scriptBuilder.ScriptLines())...)
	}

	return it
}

// AppendCmd
//
//  Empty or has issues will be ignored.
func (it *scriptOnceBuilder) AppendCmd(appendItems ...*exec.Cmd) ScriptOnceBuilder {
	if len(appendItems) == 0 {
		return it
	}

	for _, cmdExe := range appendItems {
		if cmdExe == nil {
			continue
		}

		scriptLine := CmdToScriptLine(cmdExe)

		if scriptLine == "" {
			continue
		}

		it.scriptLines.Adds(scriptLine)
	}

	return it
}

// AppendCmdOnce
//
//  Empty or has issues will be ignored.
func (it *scriptOnceBuilder) AppendCmdOnce(
	appendItems ...*CmdOnce,
) ScriptOnceBuilder {
	if len(appendItems) == 0 {
		return it
	}

	for _, cmdOnce := range appendItems {
		if cmdOnce == nil {
			continue
		}

		scriptLine := cmdOnce.WholeCommandLine()
		if scriptLine == "" {
			continue
		}

		it.scriptLines.Adds(scriptLine)
	}

	return it
}

func (it *scriptOnceBuilder) OutputMust() string {
	result := it.Result()
	errWp := result.CompiledFullErrorWrapper()
	errWp.MustBeEmpty()

	return result.OutputString()
}

func (it *scriptOnceBuilder) Log() {
	result := it.Result()

	fmt.Println(result.DetailedOutput())
}

func (it *scriptOnceBuilder) LogWithTraces() {
	result := it.Result()
	stackTraces := codestack.StacksTo.String(codestack.Skip1, codestack.DefaultStackCount)
	output := result.DetailedOutput() +
		constants.DefaultLine +
		stackTraces

	fmt.Println(output)
}

func (it *scriptOnceBuilder) LogFatal() {
	errWp := it.CompiledErrorWrapper()

	errWp.LogFatal()
}

func (it *scriptOnceBuilder) LogFatalWithTraces() {
	errWp := it.CompiledErrorWrapper()

	errWp.LogFatalWithTraces()
}

func (it *scriptOnceBuilder) LogIf(isLog bool) {
	if !isLog {
		return
	}

	result := it.Result()
	fmt.Println(result.DetailedOutput())

	errWp := result.CompiledFullErrorWrapper()
	errWp.Log()
}

func (it *scriptOnceBuilder) IsDisposed() bool {
	return it.isDisposed
}

func (it *scriptOnceBuilder) IsFinalized() bool {
	return it.isDisposed
}

func (it *scriptOnceBuilder) ErrorWithTraces() error {
	return it.Error()
}

func (it *scriptOnceBuilder) ErrorStringWithTraces() string {
	return it.CompiledErrorWrapper().CompiledStackTracesString()
}

func (it *scriptOnceBuilder) Clear() {
	it.scriptLines.Clear()
}

// IsValid
//
//  return it != nil &&
//		it.scriptType.IsValid() &&
//		!it.HasError() &&
//		it.scriptLines.HasAnyItem()
func (it *scriptOnceBuilder) IsValid() bool {
	return it != nil &&
		it.scriptType.IsValid() &&
		!it.HasError() &&
		it.scriptLines.HasAnyItem()
}

func (it *scriptOnceBuilder) IsSuccess() bool {
	return !it.HasError()
}

func (it *scriptOnceBuilder) HasAnyIssues() bool {
	return it.Result().HasAnyError()
}

// ValidationError
//
//  Contains stack-traces
func (it *scriptOnceBuilder) ValidationError() error {
	return it.Error()
}

// Error
//
//  Contains stack-traces
func (it *scriptOnceBuilder) Error() error {
	return it.CompiledErrorWrapper().
		CompiledJsonErrorWithStackTraces()
}

// TrimmedOutput
//
//  Gives compiled output string without whitespace lines
func (it *scriptOnceBuilder) TrimmedOutput() string {
	return it.Result().CompiledTrimmedOutput()
}

// String
//
//  Gives compiled output
func (it *scriptOnceBuilder) String() string {
	return it.Output()
}

// Strings
//
//  Gives compiled output lines
func (it *scriptOnceBuilder) Strings() []string {
	return it.OutputLines()
}

func (it *scriptOnceBuilder) CompiledProcessArgs() (
	process string, args []string,
) {
	return New.
		Script.
		ProcessScriptsFormat(it.scriptType, (*it.scriptLines)...)
}

// SetSecure
//
// no error detail output
func (it *scriptOnceBuilder) SetSecure() ScriptOnceBuilder {
	it.isSecure = true

	return it
}

func (it *scriptOnceBuilder) SetPlainOutput() ScriptOnceBuilder {
	it.isSecure = false

	return it
}

func (it *scriptOnceBuilder) EnableOutput() ScriptOnceBuilder {
	it.hasOutput = true

	return it
}

func (it *scriptOnceBuilder) DisableOutput() ScriptOnceBuilder {
	it.hasOutput = false

	return it
}

func (it *scriptOnceBuilder) SetScriptType(
	scriptType scripttype.Variant,
) ScriptOnceBuilder {
	it.scriptType = scriptType

	return it
}

func (it *scriptOnceBuilder) SetScriptLines(
	scriptLines *corestr.SimpleSlice,
) ScriptOnceBuilder {
	it.scriptLines = scriptLines

	return it
}

func (it *scriptOnceBuilder) SetStdIn(
	stdIn *bytes.Buffer,
) ScriptOnceBuilder {
	it.stdIn = stdIn

	return it
}

func (it *scriptOnceBuilder) SetStdOut(
	stdOut *bytes.Buffer,
) ScriptOnceBuilder {
	it.stdOut = stdOut

	return it
}

func (it *scriptOnceBuilder) SetStderr(
	stderr *bytes.Buffer,
) ScriptOnceBuilder {
	it.stderr = stderr

	return it
}

func (it *scriptOnceBuilder) SetScript(
	scriptType scripttype.Variant,
) ScriptOnceBuilder {
	it.scriptType = scriptType

	return it
}

func (it *scriptOnceBuilder) AllBuffers() (
	stdIn, stdOut, stderr *bytes.Buffer,
) {
	return it.stdIn, it.stdOut, it.stderr
}

func (it *scriptOnceBuilder) SetInputOutput(
	stdIn, stdOut, stderr *bytes.Buffer,
) ScriptOnceBuilder {
	it.stdIn = stdIn
	it.stderr = stderr
	it.stdOut = stdOut

	return it
}

func (it *scriptOnceBuilder) AppendScriptLines(
	lines ...string,
) ScriptOnceBuilder {
	if len(lines) == 0 {
		return it
	}

	it.scriptLines.Adds(lines...)

	return it
}

func (it *scriptOnceBuilder) Args(
	args ...string,
) ScriptOnceBuilder {
	it.scriptLines.Adds(
		ArgsJoin(args...))

	return it
}

func (it *scriptOnceBuilder) ArgsIf(
	isAdd bool,
	args ...string,
) ScriptOnceBuilder {
	if !isAdd {
		return it
	}

	it.scriptLines.Adds(
		ArgsJoin(args...))

	return it
}

func (it *scriptOnceBuilder) ProcessArgs(
	process string,
	args ...string,
) ScriptOnceBuilder {
	it.scriptLines.Adds(
		ProcessArgsJoinAppend(process, args...))

	return it
}

func (it *scriptOnceBuilder) Process(
	process string,
) ScriptOnceBuilder {
	it.scriptLines.Adds(process)

	return it
}

func (it *scriptOnceBuilder) ProcessArgsIf(
	isAdd bool,
	process string, args ...string,
) ScriptOnceBuilder {
	if !isAdd {
		return it
	}

	it.scriptLines.Adds(
		ProcessArgsJoinAppend(process, args...))

	return it
}

func (it *scriptOnceBuilder) ProcessIf(
	isAdd bool,
	process string,
) ScriptOnceBuilder {
	if !isAdd {
		return it
	}

	it.scriptLines.Adds(process)

	return it
}

func (it *scriptOnceBuilder) Length() int {
	if it == nil {
		return 0
	}

	return it.scriptLines.Length()
}

func (it *scriptOnceBuilder) HasAnyItem() bool {
	return it != nil && it.scriptLines.HasAnyItem()
}

func (it *scriptOnceBuilder) IsEmpty() bool {
	return it == nil || it.scriptLines.IsEmpty()
}

func (it *scriptOnceBuilder) HasError() bool {
	return it != nil && it.Build().HasAnyIssues()
}

func (it *scriptOnceBuilder) IsFailed() bool {
	return it != nil && it.Build().HasAnyIssues()
}

func (it *scriptOnceBuilder) Result() *Result {
	return it.Build().CompiledResult()
}

func (it *scriptOnceBuilder) ResultMust() *Result {
	result := it.Result()

	errWp := result.CompiledFullErrorWrapper()
	errWp.MustBeEmpty()

	return result
}

func (it *scriptOnceBuilder) CompiledErrorWrapper() *errorwrapper.Wrapper {
	return it.Build().CompiledFullErrorWrapper()
}

// Output
//
//  Gives compiled output
func (it *scriptOnceBuilder) Output() string {
	return it.Result().OutputString()
}

func (it *scriptOnceBuilder) ErrorOutput() string {
	return it.Result().ErrorString()
}

func (it *scriptOnceBuilder) DetailedOutput() string {
	return it.Result().DetailedOutput()
}

func (it *scriptOnceBuilder) OutputLines() []string {
	return it.Result().CompiledOutputLines()
}

func (it *scriptOnceBuilder) TrimmedOutputLines() []string {
	return it.Result().CompiledTrimmedOutputLines()
}

func (it *scriptOnceBuilder) OutputBytes() []byte {
	return it.Result().OutputBytes()
}

func (it *scriptOnceBuilder) StandardOutputCmd() *exec.Cmd {
	if it.isDisposed {
		panic(errnew.FinalizedResourceCannotAccess)
	}

	cmd := it.BuildCmd()

	return CmdSetOsStandardWriters(cmd)
}

func (it *scriptOnceBuilder) Build() *CmdOnce {
	if it == nil {
		return nil
	}

	if it.isDisposed {
		panic(errnew.FinalizedResourceCannotAccess)
	}

	cmdOnce := New.Script.LinesWithStdIns(
		it.hasOutput,
		it.isSecure,
		it.stdIn,
		it.stdOut,
		it.stderr,
		it.scriptType,
		(*it.scriptLines)...)

	return it.setEnv(cmdOnce)
}

func (it *scriptOnceBuilder) setEnv(cmdOnce *CmdOnce) *CmdOnce {
	if it.isSetCurrentEnv {
		cmdOnce.InitializeEnvVars()
	}

	if it.envVars.HasAnyItem() {
		cmdOnce.AddEnvVarsHashmap(it.envVars.Hashmap())
	}

	return cmdOnce
}

func (it *scriptOnceBuilder) BuildClear() *CmdOnce {
	cmdOnce := it.Build()
	it.Clear()

	return cmdOnce
}

func (it *scriptOnceBuilder) BuildDispose() *CmdOnce {
	cmdOnce := it.Build()
	it.Dispose()

	return cmdOnce
}

func (it *scriptOnceBuilder) BuildCmd() *exec.Cmd {
	if it.isDisposed {
		panic(errnew.FinalizedResourceCannotAccess)
	}

	cmdOnce := New.Script.LinesWithStdIns(
		it.hasOutput,
		it.isSecure,
		it.stdIn,
		it.stdOut,
		it.stderr,
		it.scriptType,
		(*it.scriptLines)...)

	cmdOnce = it.setEnv(cmdOnce)
	cmd := cmdOnce.Cmd

	if !it.HasStdIn() {
		cmd.Stdin = nil
	}

	if !it.HasStdOut() {
		cmd.Stdout = nil
	}

	if !it.HasStdErr() {
		cmd.Stderr = nil
	}

	return cmd
}

func (it *scriptOnceBuilder) BuildCmdWithStdIn() (cmd *exec.Cmd, stdInWriter *bytes.Buffer) {
	cmdOnce := it.Build()

	return cmdOnce.Cmd, it.stdIn
}

func (it *scriptOnceBuilder) Dispose() {
	if it == nil {
		return
	}

	it.isDisposed = true
	it.isSecure = false
	it.hasOutput = false
	it.scriptType = scripttype.Invalid
	it.envVars = nil // todo https://github.com/alimtvnetwork/core-v9/-/issues/85
	it.scriptLines.Dispose()
	it.scriptLines = nil
	it.stdIn = nil
	it.stdOut = nil
	it.stderr = nil
}

func (it scriptOnceBuilder) AsScriptOnceBuilderContractsBinder() ScriptOnceBuilderContractsBinder {
	return &it
}
