package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

type Result struct {
	Request      Request
	CmdOnce      *errcmd.CmdOnce
	ErrorWrapper *errorwrapper.Wrapper
	ExitCode     linuxservicestate.ExitCode
}

func (it *Result) ServiceName() string {
	return it.Request.ServiceName
}

func (it *Result) IsSuccess() bool {
	return it.ExitCode.IsActiveRunning()
}

func (it *Result) IsFailed() bool {
	return !it.ExitCode.IsActiveRunning()
}

func (it *Result) IsUnknownService() bool {
	return it.ExitCode.IsUnknownService()
}

func (it *Result) IsEmpty() bool {
	return it == nil ||
		it.CmdOnce == nil ||
		it.ExitCode.IsInvalidCode()
}

func (it *Result) ErrorWrapperUsingOpt(isDetailedErr bool) *errorwrapper.Wrapper {
	if isDetailedErr {
		return it.CompiledErrorWrapper()
	}

	return it.SimplifiedError()
}

func (it *Result) CompiledErrorWrapper() *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.WithMessage("CmdResult is nil", it)
	}

	if it.CmdOnce == nil {
		return errnew.Null.WithMessage("CmdOnce nil", it.CmdOnce)
	}

	if it.ErrorWrapper != nil {
		return it.ErrorWrapper
	}

	return it.CmdOnce.CompiledErrorWrapper()
}

func (it *Result) SimplifiedError() *errorwrapper.Wrapper {
	return exitCodeToMappedError(it)
}

func (it *Result) CollectSimpleError(
	errCollection *errwrappers.Collection,
) bool {
	err := it.SimplifiedError()
	errCollection.AddWrapperPtr(err)

	return err.IsSuccess()
}

func (it *Result) CollectError(
	errCollection *errwrappers.Collection,
	isSimpleError bool,
) bool {
	if isSimpleError {
		return it.CollectSimpleError(errCollection)
	}

	err := it.CompiledErrorWrapper()
	errCollection.AddWrapperPtr(err)

	return err.IsSuccess()
}

func (it *Result) VerifyExitCode(expectedExitCode linuxservicestate.ExitCode) *errorwrapper.Wrapper {
	if it.ExitCode == expectedExitCode {
		return nil
	}

	return errnew.WasExpecting(
		errtype.ServiceStatusCodeMismatch,
		it.ServiceName()+" is not in the expected status!",
		expectedExitCode.NameValue(),
		it.ExitCode.NameValue())
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.DisposeWithoutErrorWrapper()
	it.ErrorWrapper.Dispose()
}

func (it *Result) DisposeWithoutErrorWrapper() {
	if it == nil {
		return
	}

	it.Request.Dispose()
	it.CmdOnce.Dispose()
	it.ExitCode = linuxservicestate.Invalid
}
