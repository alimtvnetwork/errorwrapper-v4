package errnew

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

func Empty() *errorwrapper.Wrapper {
	return errorwrapper.EmptyPtr()
}

func OnEmpty(
	wrapper *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if wrapper == nil {
		return Empty()
	}

	return wrapper
}

func WasExpecting(
	variant errtype.Variation,
	title string,
	expecting, actual interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		defaultSkipInternal,
		variant,
		constants.Space,
		errcore.ExpectingSimpleNoType(title, expecting, actual))
}

func WasExpectingUsingStackSkip(
	stackSkipIndex int,
	variant errtype.Variation,
	title string,
	expecting, actual interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMessagesUsingJoiner(
		stackSkipIndex+defaultSkipInternal,
		variant,
		constants.Space,
		errcore.ExpectingSimpleNoType(title, expecting, actual))
}

func NotSupportedOptionUsingStackSkip(
	stackSkipIndex int,
	varName string, val interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	refsCollection := refs.
		New(1).
		Add(varName, val)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errtype.NotSupportedOption,
		strings.Join(messages, constants.Space),
		refsCollection)
}

func NotSupportedOption(
	varName string, val interface{},
	messages ...string,
) *errorwrapper.Wrapper {
	refsCollection := refs.
		New(1).
		Add(varName, val)

	return errorwrapper.NewMsgUsingAllParams(
		defaultSkipInternal,
		errtype.NotSupportedOption,
		true,
		strings.Join(messages, constants.Space),
		refsCollection)
}

func OutOfRange(
	start, end, ranges interface{},
	message string,
) *errorwrapper.Wrapper {
	msg := fmt.Sprintf(
		errconsts.RangeWithOutRangeFormat,
		message,
		start,
		end)

	finalMessage := errtype.OutOfRangeValue.Combine(
		msg,
		errconsts.Ranges,
		ranges,
	)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errtype.OutOfRangeValue,
		finalMessage,
		nil)
}

func NotImplPanic(
	url string,
	todoMessage string,
) {
	NotImpl(url).HandleErrorWithMsg(todoMessage)
}

func NotImpl(
	url string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errtype.NotImplemented,
		"TODO: url("+url+")",
		nil)
}

func NotImplPtrUsingStackSkip(
	stackSkipIndex int,
	url string,
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errtype.NotImplemented,
		"TODO: url("+url+")",
		nil)
}
