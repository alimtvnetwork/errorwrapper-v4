package errcmd

import (
	"bytes"
	"os/exec"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corerange"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type Result struct {
	*baseBufferStdOutError
	ExitError                *exec.ExitError
	ExitCode                 int
	hasCommandExecuted       bool
	detailedOutput           corestr.SimpleStringOnce
	errorWrapper             *errorwrapper.Wrapper
	compiledFullErrorWrapper *errorwrapper.Wrapper // represents errorwrapper.Wrapper using error + std error line
}

func NewResult(
	errorWrapper *errorwrapper.Wrapper,
	stdOut, stdErr *bytes.Buffer,
	exitError *exec.ExitError,
	hasCommandExecuted bool,
) *Result {
	exitCode := GetExitCode(errorWrapper, exitError)

	return &Result{
		baseBufferStdOutError: &baseBufferStdOutError{
			stdErr: stdErr,
			stdOut: stdOut,
		},
		errorWrapper:       errorWrapper,
		ExitError:          exitError,
		ExitCode:           exitCode,
		hasCommandExecuted: hasCommandExecuted,
	}
}

func NewResultUsingBaseBuffer(
	errorWrapper *errorwrapper.Wrapper,
	baseBuffer *baseBufferStdOutError,
	exitError *exec.ExitError,
	hasCommandExecuted bool,
) *Result {
	exitCode := GetExitCode(errorWrapper, exitError)

	return &Result{
		baseBufferStdOutError: baseBuffer,
		ExitError:             exitError,
		ExitCode:              exitCode,
		errorWrapper:          errorWrapper,
		hasCommandExecuted:    hasCommandExecuted,
	}
}

func (it *Result) HasCommandExecuted() bool {
	return it.hasCommandExecuted
}

func (it *Result) IsInvalidExitCode() bool {
	return it.ExitCode == InvalidExitCode
}

func (it *Result) IsRunSuccessfully() bool {
	return it.ExitCode == SuccessfullyRunningExitCode
}

func (it *Result) IsSuccess() bool {
	return it.ExitCode == SuccessfullyRunningExitCode
}

func (it *Result) HasAnyError() bool {
	return it != nil && it.HasError() || it.baseBufferStdOutError.HasErrorBufferData()
}

func (it *Result) IsFailed() bool {
	return it.ExitCode != SuccessfullyRunningExitCode
}

func (it *Result) IsInvalid() bool {
	return it.ExitCode != SuccessfullyRunningExitCode
}

func (it *Result) ExitCodeByte() byte {
	if it == nil || it.ExitError == nil || it.IsInvalidExitCode() {
		return constants.Zero
	}

	value, _ := corerange.
		Within.
		RangeByteDefault(it.ExitCode)

	return value
}

func (it *Result) IsExitCode(exitCode int) bool {
	return it.ExitCode == exitCode
}

func (it *Result) IsExitCodeByte(exitCode byte) bool {
	return it.ExitCode > 0 &&
		it.ExitCode == int(exitCode)
}

func (it *Result) HasExitError() bool {
	return it.ExitError != nil
}

func (it *Result) HasValidExitCode() bool {
	return !it.IsInvalidExitCode()
}

func (it *Result) IsEmptyError() bool {
	return it == nil || it.errorWrapper.IsEmpty()
}

func (it *Result) HasNoError() bool {
	return it == nil ||
		it.errorWrapper.IsEmpty() &&
			!it.baseBufferStdOutError.HasErrorBufferData()
}

func (it *Result) HasError() bool {
	return it != nil && it.errorWrapper.HasError()
}

// CompiledFullErrorWrapper
//
// Represents compiled error wrapper + std error line
//
// (if baseBufferStdOutError has any error bytes)
func (it *Result) CompiledFullErrorWrapper() *errorwrapper.Wrapper {
	if it.compiledFullErrorWrapper != nil {
		return it.compiledFullErrorWrapper
	}

	if it.HasNoError() {
		return nil
	}

	if it.HasError() {
		// has error indicates
		// that error is already merged
		it.compiledFullErrorWrapper = it.
			errorWrapper
	} else if it.HasErrorBufferData() {
		// has no error but has error lines
		it.compiledFullErrorWrapper = errorwrapper.NewMsgDisplayErrorNoReference(
			codestack.Skip1,
			errtype.FailedProcess,
			it.ErrorString(),
		)
	}

	return it.compiledFullErrorWrapper
}

func (it *Result) ErrorWrapper() *errorwrapper.Wrapper {
	return it.errorWrapper
}

// AllErrorString
//
// errstr.Result Value represents std error buffer lines
// errstr.Result ErrorWrapper represents the Result.ErrorWrapper
func (it *Result) AllErrorString() (errorBufferLine string, compiledErrorWrapper *errorwrapper.Wrapper) {
	if it.HasNoError() {
		return "", nil
	}

	return it.ErrorString(), it.errorWrapper
}

// AllErrorBytes
//
// errstr.Result Value represents std error buffer lines
// errstr.Result ErrorWrapper represents the Result.ErrorWrapper
func (it *Result) AllErrorBytes() (errorBufferBytes []byte, compiledErrorWrapper *errorwrapper.Wrapper) {
	if it.HasNoError() {
		return nil, nil
	}

	return it.ErrorBytes(), it.errorWrapper
}

func (it Result) String() string {
	if it.HasError() {
		return it.errorWrapper.FullString()
	}

	if it.stdOut == nil && it.stdErr != nil {
		return it.ErrorString()
	}

	return it.OutputString()
}

func (it Result) StringIf(isErrorOnly bool) string {
	if it.HasError() {
		return it.errorWrapper.FullString()
	}

	if isErrorOnly || it.stdOut == nil && it.stdErr != nil {
		return it.ErrorString()
	}

	return it.OutputString()
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.DisposeWithoutErrorWrapper()
	it.ExitError = nil
	it.errorWrapper.Dispose()
	it.compiledFullErrorWrapper.Dispose()
}

func (it *Result) DisposeWithoutErrorWrapper() {
	if it == nil {
		return
	}

	it.baseBufferStdOutError.Dispose()
	it.detailedOutput.Dispose()
	it.baseBufferStdOutError = nil
}

func (it *Result) DetailedOutput() string {
	if it.detailedOutput.IsInitialized() {
		return it.detailedOutput.String()
	}

	var toString string
	if it.HasAnyError() {
		toString = it.
			CompiledFullErrorWrapper().
			FullString() +
			constants.NewLineUnix
	}

	toString += it.OutputString()

	return it.
		detailedOutput.
		GetSetOnce(toString)
}

func (it *Result) HandleError() {
	if it.HasAnyError() {
		it.CompiledFullErrorWrapper().LogFatalWithTraces()
	}
}
