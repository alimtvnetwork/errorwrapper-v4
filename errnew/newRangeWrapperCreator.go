package errnew

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type newRangeWrapperCreator struct{}

func (it newRangeWrapperCreator) Within(
	currentValue interface{},
	min, max interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.WithinRange,
		"",
		ref.Value{
			Variable: "Value Received",
			Value:    currentValue,
		},
		ref.Value{
			Variable: "Min",
			Value:    min,
		},
		ref.Value{
			Variable: "Max",
			Value:    max,
		})
}

func (it newRangeWrapperCreator) OutOf(
	currentValue interface{},
	min, max interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		"",
		ref.Value{
			Variable: "Value Received",
			Value:    currentValue,
		},
		ref.Value{
			Variable: "Min",
			Value:    min,
		},
		ref.Value{
			Variable: "Max",
			Value:    max,
		})
}

func (it newRangeWrapperCreator) OutOfRanges(
	currentValue interface{},
	ranges ...interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		"Didn't meet the expectation!",
		ref.Value{
			Variable: "Value Received",
			Value:    currentValue,
		},
		ref.Value{
			Variable: "Ranges",
			Value:    ranges,
		})
}

func (it newRangeWrapperCreator) MessageOutOf(
	message string,
	reference interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		message,
		ref.Value{
			Variable: "Reference",
			Value:    reference,
		})
}

func (it newRangeWrapperCreator) Error(
	err error,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errtype.OutOfRange,
		err.Error(),
	)
}

func (it newRangeWrapperCreator) ErrorRef(
	err error,
	reference interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		err.Error(),
		ref.Value{
			Variable: "Reference",
			Value:    reference,
		})
}

func (it newRangeWrapperCreator) ErrorRefStartEnd(
	err error,
	currentValue interface{},
	start, end interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		err.Error(),
		ref.Value{
			Variable: "Value Received",
			Value:    currentValue,
		},
		ref.Value{
			Variable: "start",
			Value:    start,
		},
		ref.Value{
			Variable: "end",
			Value:    end,
		})
}

func (it newRangeWrapperCreator) ErrorRefStartEndValues(
	err error,
	currentValue interface{},
	start, end interface{},
	possibleValues interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		err.Error(),
		ref.Value{
			Variable: "Value Received",
			Value:    currentValue,
		},
		ref.Value{
			Variable: "start",
			Value:    start,
		},
		ref.Value{
			Variable: "end",
			Value:    end,
		},
		ref.Value{
			Variable: "Possible Values",
			Value:    possibleValues,
		})
}

func (it newRangeWrapperCreator) RefStartEndValues(
	currentValue interface{},
	start, end interface{},
	possibleValues interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		"Didn't meet the expectation!",
		ref.Value{
			Variable: "Value Received",
			Value:    currentValue,
		},
		ref.Value{
			Variable: "start",
			Value:    start,
		},
		ref.Value{
			Variable: "end",
			Value:    end,
		},
		ref.Value{
			Variable: "Possible Values",
			Value:    possibleValues,
		})
}

func (it newRangeWrapperCreator) StartEnd(
	start, end interface{},
) *errorwrapper.Wrapper {
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errtype.OutOfRange,
		"Didn't meet the expectation!",
		ref.Value{
			Variable: "start",
			Value:    start,
		},
		ref.Value{
			Variable: "end",
			Value:    end,
		})
}
