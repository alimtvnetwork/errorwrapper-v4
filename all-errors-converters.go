package errorwrapper

import (
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/internal/consts"
)

// ErrorsToString
//
//  nil items will be ignored.
func ErrorsToString(errItems ...error) string {
	if len(errItems) == 0 {
		return ""
	}

	toString := ErrorsToStringUsingJoiner(
		consts.DefaultErrorLineSeparator,
		errItems...)

	return toString
}

// ErrorsToError
//
//  nil items will be ignored.
func ErrorsToError(errItems ...error) error {
	if len(errItems) == 0 {
		return nil
	}

	toString := ErrorsToStringUsingJoiner(
		consts.DefaultErrorLineSeparator,
		errItems...)

	if toString == "" {
		return nil
	}

	toWrap := NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.Generic,
		toString)
	toErr := toWrap.
		CompiledErrorWithStackTraces()

	go toWrap.Dispose()
	toWrap = nil

	return toErr
}

// ErrorsToWrap
//
//  nil items will be ignored.
func ErrorsToWrap(
	errType errtype.Variation,
	errItems ...error,
) *Wrapper {
	if len(errItems) == 0 {
		return nil
	}

	toString := ErrorsToString(errItems...)

	if toString == "" {
		return nil
	}

	return NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		toString)
}
