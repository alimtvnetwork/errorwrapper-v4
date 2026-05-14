package errcmd

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/alimtvnetwork/core-v9/conditional"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/core-v9/issetter"
	"github.com/alimtvnetwork/core-v9/simplewrap"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type CmdOnce struct {
	baseCmdWrapper                   BaseCmdWrapper
	toString                         corestr.SimpleStringOnce
	Cmd                              *exec.Cmd
	compiledResult                   *Result
	hasCmd                           bool
	isEnvironmentVariableInitialized bool           // Will get updated by AddEnvVarKeyVal, AddEnvVar,includeEnvVars  command
	isSuccessfullyExecuted           issetter.Value // Will get updated by RunOnce command
	isAlreadyRan                     bool           // Will get updated by RunOnce command
	hasOutput, hasSecureData         bool
}

// ArgumentsSingleLine Arguments Compiled using Space
func (it *CmdOnce) ArgumentsSingleLine() string {
	return it.baseCmdWrapper.argumentsSingleLine
}

// WholeCommandLine Process + Space + ArgumentsCompiledWithSpace
func (it *CmdOnce) WholeCommandLine() string {
	return it.baseCmdWrapper.wholeCommand
}

// DoubleQuoteWholeCommandLine Process + Space + ArgumentsCompiledWithSpace
func (it *CmdOnce) DoubleQuoteWholeCommandLine() string {
	return simplewrap.WithDoubleQuote(it.baseCmdWrapper.wholeCommand)
}

func (it *CmdOnce) includeEnvVars(cmd *exec.Cmd) {
	if it.isEnvironmentVariableInitialized {
		return
	}

	cmd.Env = os.Environ()
	it.isEnvironmentVariableInitialized = true
}

// InitializeEnvVars
//
//  Warning must be set before run.
func (it *CmdOnce) InitializeEnvVars() {
	if it.HasIssues() || it.isEnvironmentVariableInitialized {
		return
	}

	it.includeEnvVars(it.Cmd)
}

// AddEnvVarKeyVal cannot add slice if already executed.
// Warning : panic if called after execution.
func (it *CmdOnce) AddEnvVarKeyVal(keyVal corestr.KeyValuePair) *CmdOnce {
	return it.AddEnvVar(keyVal.Key, keyVal.Value)
}

// AddEnvVar cannot add slice if already executed.
// Warning : panic if called after execution.
func (it *CmdOnce) AddEnvVar(
	varName, varValue string,
) *CmdOnce {
	if it.HasIssues() {
		return it
	}

	it.panicOnAlreadyRan()

	return it.addEnvVarInternal(
		it.Cmd,
		varName,
		varValue)
}

func (it *CmdOnce) addEnvVarInternal(
	cmd *exec.Cmd,
	varName, varValue string,
) *CmdOnce {
	it.includeEnvVars(cmd)
	cmd.Env = append(
		cmd.Env,
		it.GetFormattedKeyValueData(varName, varValue))

	return it
}

func (it *CmdOnce) panicOnAlreadyRan() {
	if it.IsAlreadyRan() {
		// panic here.
		errnew.
			FinalizedResourceCannotAccess.
			HandleError()
	}
}

// GetFormattedKeyValueData "MY_VAR=some_value"
func (it *CmdOnce) GetFormattedKeyValueData(
	varName string,
	varValue string,
) string {
	return varName + constants.EqualSymbol + varValue
}

// AddEnvVarsHashmap
//
// cannot add slice if already executed.
// Warning : panic if called after execution.
func (it *CmdOnce) AddEnvVarsHashmap(
	hashmap *corestr.Hashmap,
) *CmdOnce {
	if hashmap.IsEmpty() {
		return it
	}

	slice := hashmap.ToStringsUsingCompiler(func(key, val string) string {
		return it.GetFormattedKeyValueData(key, val)
	})

	return it.AddEnvVarsSlice(*slice...)
}

// AddEnvVarsSlice
//
// cannot add slice if already executed.
// Warning : panic if called after execution.
// Data needs to be in this format "MY_VAR=some_value"
func (it *CmdOnce) AddEnvVarsSlice(
	slice ...string,
) *CmdOnce {
	if len(slice) == 0 || it.HasIssues() {
		return it
	}

	it.panicOnAlreadyRan()
	it.includeEnvVars(it.Cmd)
	it.Cmd.Env = append(it.Cmd.Env, slice...)

	return it
}

func (it *CmdOnce) HasCmd() bool {
	return it.hasCmd
}

func (it *CmdOnce) HasCompiledError() bool {
	return it.HasAnyIssues() || it.CompiledResult().HasError()
}

func (it *CmdOnce) setAlreadyRan() {
	it.isAlreadyRan = true
}

func (it *CmdOnce) setExecutionState(isSuccess bool) {
	it.isSuccessfullyExecuted = issetter.GetBool(isSuccess)
}

// IsSuccessfullyExecuted Runs the execution and then returns the success result.
//
// Reveals if any issue after compile, thus runs Compile() or RunOnce() first then returns bool.
func (it *CmdOnce) IsSuccessfullyExecuted() bool {
	if it.isSuccessfullyExecuted.IsUninitialized() || !it.IsAlreadyRan() {
		// it.isSuccessfullyExecuted
		// will get updated by RunOnce command.
		it.RunOnce()
	}

	return it.isSuccessfullyExecuted.IsTrue()
}

// RunOnceWithSuccessFlag Runs the process then returns the isSuccess flag
func (it *CmdOnce) RunOnceWithSuccessFlag() (cmdResult *Result, isSuccess bool) {
	return it.RunOnce(), it.isSuccessfullyExecuted.IsTrue()
}

// RunOnce only runs once the process and returns the cached data many times.
func (it *CmdOnce) RunOnce() *Result {
	if it.compiledResult != nil {
		return it.compiledResult
	}

	if it.HasIssues() {
		compiledResult := it.compiledResultBasedOnIssues()
		it.setCompiledResult(compiledResult)

		return compiledResult
	}

	compiledResult := it.getCompiledResultByExecuting()
	it.setCompiledResult(compiledResult)

	return compiledResult
}

func (it *CmdOnce) NewArgs(additionalArgs ...string) *CmdOnce {
	existingArgs := corestr.New.SimpleSlice.Direct(
		true,
		it.baseCmdWrapper.arguments,
	)

	existingArgs.Adds(additionalArgs...)

	return New.Create(
		it.hasOutput,
		it.hasSecureData,
		it.ProcessName(),
		existingArgs.Items...)
}

// Clone calls constructor Lines to create itself using the existing data.
func (it *CmdOnce) Clone() *CmdOnce {
	return New.Create(
		it.hasOutput,
		it.hasSecureData,
		it.ProcessName(),
		it.Arguments()...)
}

func (it *CmdOnce) CmdCloneWithoutStd() *exec.Cmd {
	return it.CmdCloneUsingStds(nil, nil)
}

func (it *CmdOnce) CmdCloneCompiledOutputBytes() (cmd *exec.Cmd, allBytes []byte, err error) {
	cmd = it.CmdCloneUsingStds(nil, nil)

	errWholeLine := conditional.StringTrueFunc(
		!it.hasSecureData,
		it.WholeCommandLine)

	if cmd == nil {
		return nil, []byte{},
			errcore.FailedToCreateCmdType.Error(
				"cmd nil", errWholeLine)
	}

	allBytes, err = cmd.CombinedOutput()

	if err != nil {
		return cmd, allBytes, errtype.FailedProcess.ErrorNoRefs(
			"Command:" + errWholeLine)
	}

	return cmd, allBytes, err
}

func (it *CmdOnce) CmdCloneCompiledOutputString() (fullOutput corestr.SimpleStringOnce, err error) {
	_, allBytes, err := it.CmdCloneCompiledOutputBytes()

	if allBytes == nil {
		return corestr.Empty.SimpleStringOnce(), err
	}

	return corestr.New.SimpleStringOnce.Init(string(allBytes)), err
}

func (it *CmdOnce) CmdCloneCompiledOutputStringLines() (fullOutputLines []string, err error) {
	outputString, err := it.CmdCloneCompiledOutputString()

	return outputString.Split(constants.NewLineUnix), err
}

func (it *CmdOnce) CmdCloneCompiledOutputTrimStringLines() (fullOutputLines []string, err error) {
	lines, err := it.CmdCloneCompiledOutputStringLines()

	return stringslice.NonWhitespaceTrimSlice(lines), err
}

func (it *CmdOnce) CmdCloneUsingStds(
	outBuff, errBuff *bytes.Buffer,
) *exec.Cmd {
	cmd := exec.Command(
		it.ProcessName(),
		it.Arguments()...)

	hasCmd := cmd != nil
	hasSecureData := it.hasSecureData
	hasOutput := it.hasOutput

	if hasCmd && len(it.Cmd.Env) > 0 {
		cmd.Env = stringslice.Clone(it.Cmd.Env)
	}

	if hasCmd && !hasSecureData && errBuff != nil {
		cmd.Stderr = errBuff
	}

	if hasCmd && hasOutput && outBuff != nil {
		outBuff.Grow(constants.Capacity1G)
		cmd.Stdout = outBuff
	}

	return cmd
}

// CloneCmd Clones and creates a new cmd
//
// If isUseStd true then it will clone it-self and returns the cmd and buffer will be created and injected,
//
// or else it will heavily depend on CompiledOutput or CombinedOutput
func (it *CmdOnce) CloneCmd(isUseStd bool) *exec.Cmd {
	if isUseStd {
		return it.Clone().Cmd
	}

	cmdNew := exec.Command(it.ProcessName(), it.Arguments()...)

	return cmdNew
}

func (it *CmdOnce) Arguments() []string {
	return it.baseCmdWrapper.arguments
}

func (it *CmdOnce) ProcessName() string {
	return it.baseCmdWrapper.process
}

// Run non lazy and runs as many times called.
//
// Warning: Costly operation, run it wisely.
//
// Under the hood, it creates the same cmdOnce and call it's RunOnce.
func (it *CmdOnce) Run() *Result {
	if !it.IsAlreadyRan() {
		return it.RunOnce()
	}

	return it.Clone().RunOnce()
}

func (it *CmdOnce) getCompiledResultByExecuting() *Result {
	err := it.Cmd.Run()
	execErrorWrapper, exitError := it.getCompileErrorWrapper(err)
	compiledResult := NewResultUsingBaseBuffer(
		execErrorWrapper,
		it.baseCmdWrapper.baseBufferStdOutError,
		exitError,
		true)

	return compiledResult
}

func (it *CmdOnce) compiledResultBasedOnIssues() *Result {
	compiledResult := NewResult(
		it.initializeErrorWrapper(),
		nil,
		nil,
		nil,
		false)

	return compiledResult
}

func (it *CmdOnce) setCompiledResult(
	compiledResult *Result,
) {
	it.setAlreadyRan()
	it.setExecutionState(compiledResult.IsEmptyError())
	it.compiledResult = compiledResult
}

// getCompileErrorWrapper
//
// takes in compile error and converts to errorwrapper.Wrapper
//
// Error + Std Error Lines : if execution has error and
// no secure output then adds error line to error
func (it *CmdOnce) getCompileErrorWrapper(err error) (
	compiledError *errorwrapper.Wrapper,
	exitError *exec.ExitError,
) {
	hasError := err != nil

	if !hasError {
		return nil, nil
	}

	if hasError && !it.hasSecureData {
		compiledError = errnew.Messages.Many(
			errtype.CommandExecution,
			"\n\"CmdOnce - Error\"\n",
			"CommandLine:\n"+it.WholeCommandLine(),
			it.stdErr().String(),
			err.Error())
	} else if hasError {
		compiledError = errnew.Type.Error(
			errtype.CommandExecution,
			err)
	}

	if exitError2, isOkay := err.(*exec.ExitError); isOkay {
		return compiledError, exitError2
	}

	return compiledError, nil
}

// stdErr it only gives the result after running the RunOnce command.
// Beware and use it carefully.
func (it *CmdOnce) stdErr() *bytes.Buffer {
	return it.
		baseCmdWrapper.
		baseBufferStdOutError.
		stdErr
}

// stdOut it only gives the result after running the RunOnce command.
// Beware and use it carefully.
func (it *CmdOnce) stdOut() *bytes.Buffer {
	return it.
		baseCmdWrapper.
		baseBufferStdOutError.
		stdOut
}

func (it *CmdOnce) baseBuffer() *baseBufferStdOutError {
	return it.
		baseCmdWrapper.
		baseBufferStdOutError
}

func (it *CmdOnce) initializeErrorWrapper() *errorwrapper.Wrapper {
	return it.
		baseCmdWrapper.
		initializeErrorWrapper
}

func (it *CmdOnce) IsAlreadyRan() bool {
	return it.isAlreadyRan
}

func (it *CmdOnce) IsNull() bool {
	return it.Cmd == nil
}

// CompiledErrorWrapper
//
// Runs CmdOnce
//
// using RunOnce then returns the final compiled error
func (it *CmdOnce) CompiledErrorWrapper() *errorwrapper.Wrapper {
	return it.CompiledResult().errorWrapper
}

// CompiledErrorWrapperWithErrorBufferLine
//
// Runs CmdOnce
//
// using RunOnce
// then returns the final compiled error with error buffer lines
//
// Calls Result.AllErrorString()
// Here, the value in result represent the buffer error string (from std)
// and error wrapper represents the compiled ErrorWrapper
func (it *CmdOnce) CompiledErrorWrapperWithErrorBufferLine() (
	errorBufferLine string, compiledErrorWrapper *errorwrapper.Wrapper,
) {
	return it.CompiledResult().AllErrorString()
}

// CompiledFullErrorWrapper
//
// Represents compiled error wrapper + std error line
//
// (if baseBufferStdOutError has any error bytes)
func (it *CmdOnce) CompiledFullErrorWrapper() *errorwrapper.Wrapper {
	return it.CompiledResult().CompiledFullErrorWrapper()
}

// CompiledErrorWrapperWithErrorBufferBytes
//
// Runs CmdOnce
//
// using RunOnce
// then returns the final compiled error with error buffer lines
//
// Calls Result.AllErrorBytes()
// Here, the value in result represent the buffer error bytes (from std)
// and error wrapper represents the compiled ErrorWrapper
func (it *CmdOnce) CompiledErrorWrapperWithErrorBufferBytes() (
	errorBufferBytes []byte, compiledErrorWrapper *errorwrapper.Wrapper,
) {
	return it.CompiledResult().AllErrorBytes()
}

// CompiledResult
//
// RunOnce and submits the compiled result structure.
func (it *CmdOnce) CompiledResult() *Result {
	if it.compiledResult != nil {
		return it.compiledResult
	}

	return it.RunOnce()
}

// CompiledErrorBytes returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledErrorBytes() []byte {
	return it.CompiledResult().ErrorBytes()
}

// CompiledError returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledError() string {
	return it.CompiledResult().ErrorString()
}

// CompiledOutputBytes returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledOutputBytes() []byte {
	return it.CompiledResult().OutputBytes()
}

// CompiledOutput returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledOutput() string {
	return it.CompiledResult().OutputString()
}

// CompiledOutputLines returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledOutputLines() []string {
	return it.CompiledResult().CompiledOutputLines()
}

// CompiledErrorLines returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledErrorLines() []string {
	return it.CompiledResult().CompiledErrorLines()
}

// CompiledTrimmedOutputLines returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledTrimmedOutputLines() []string {
	return it.CompiledResult().CompiledTrimmedOutputLines()
}

// CompiledTrimmedErrorLines returns lazy once result.
//
// Developer may check:
//  - IsSuccessfullyExecuted() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - HasAnyIssues() reveals if any issue after compile, thus runs the compile or RunOnce() first then returns bool. Or,
//  - isAlreadyRan() reveals already executed or not. Doesn't run the process. Or,
//  - HasCmd() reveals if cmd is nil or not. Or,
//  - HasIssues() reveals any issues before running cmd. Or,
//  - CompiledErrorWrapper() to check the final error condition.
func (it *CmdOnce) CompiledTrimmedErrorLines() []string {
	return it.CompiledResult().CompiledTrimmedErrorLines()
}

// CommandLine Gets commandline based on security.
// if receiver.hasSecureData then only returns process name
func (it *CmdOnce) CommandLine() string {
	if it.hasSecureData {
		return simplewrap.WithDoubleQuote(it.ProcessName())
	}

	return it.DoubleQuoteWholeCommandLine()
}

func (it *CmdOnce) GetCommandLineDataDependingOnSecurity() string {
	return it.CommandLine()
}

// HasAnyIssues
//
// executes run command and gets the compiled
// result and check if any compilation error
// Any issues include compile or cmd running error.
func (it *CmdOnce) HasAnyIssues() bool {
	return it.HasIssues() ||
		it.CompiledResult().HasError()
}

// HasIssues reveals either cmd is nil or has any other
// cmd initialize error but it doesn't conclude running issues.
func (it *CmdOnce) HasIssues() bool {
	return it.IsNull() ||
		it.initializeErrorWrapper().HasError()
}

func (it *CmdOnce) Dispose() {
	if it == nil {
		return
	}

	it.DisposeWithoutResult()
	it.DisposeWithoutErrorWrapper()
	it.compiledResult.Dispose()
	it.compiledResult = nil
}

func (it *CmdOnce) DisposeWithoutResult() {
	if it == nil {
		return
	}

	it.baseCmdWrapper.DisposeWithoutErrorWrapper()
	it.Cmd = nil
	it.baseCmdWrapper.Dispose()
}

func (it *CmdOnce) DisposeWithoutErrorWrapper() {
	if it == nil {
		return
	}

	it.baseCmdWrapper.DisposeWithoutErrorWrapper()
	it.compiledResult.DisposeWithoutErrorWrapper()
	it.compiledResult = nil
	it.Cmd = nil
}

// String()
//
// If cmd is present and HasIssues false then it will call RunOnce and get the output and returns it.
func (it *CmdOnce) String() string {
	if it.toString.IsInitialized() {
		return it.toString.Value()
	}

	if it.HasIssues() {
		toString := it.initializeErrorWrapper().
			FullStringWithTraces

		return it.
			toString.
			GetOnceFunc(
				toString)
	}

	compiledOutput := it.
		CompiledResult().
		DetailedOutput()

	toString := "Command : " +
		it.CommandLine() +
		", Outputs : \n" +
		compiledOutput

	return it.
		toString.
		GetSetOnce(
			toString)
}
