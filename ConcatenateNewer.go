package errorwrapper

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type ConcatenateNewer interface {
	NewStackSkip(stackSkip int) *Wrapper
	Error(err error) *Wrapper
	ErrorUsingStackSkip(skipStackIndex int, err error) *Wrapper

	Msg(errMsg string) *Wrapper
	MsgUsingStackSkip(skipStackIndex int, errMsg string) *Wrapper
	MessagesUsingStackSkip(skipStackIndex int, errMessages ...string) *Wrapper

	MsgRefs(
		skipStackIndex int,
		errMsg string,
		references ...ref.Value,
	) *Wrapper
	MsgRefsOnly(
		references ...ref.Value,
	) *Wrapper

	Errors(
		errItems ...error,
	) *Wrapper
	ErrorsUsingStackSkip(skipStackIndex int, errItems ...error) *Wrapper
	Wrapper(another *Wrapper) *Wrapper
	WrapperUsingStackSkip(skipStackIndex int, another *Wrapper) *Wrapper

	ErrorInterface(errInf errcoreinf.BaseErrorOrCollectionWrapper) *Wrapper
	ErrorInterfaceUsingStackSkip(
		stackSkip int,
		errInf errcoreinf.BaseErrorOrCollectionWrapper,
	) *Wrapper

	BasicError(
		basicErr errcoreinf.BasicErrWrapper,
	) *Wrapper
	BasicErrorUsingStackSkip(
		stackSkip int,
		basicErr errcoreinf.BasicErrWrapper,
	) *Wrapper

	NewStackTraces(
		stackTraces ...codestack.Trace,
	) *Wrapper
	CloneStackSkip(
		skipStackIndex int,
	) *Wrapper
}
