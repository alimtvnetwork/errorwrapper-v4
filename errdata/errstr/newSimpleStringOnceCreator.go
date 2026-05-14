package errstr

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newSimpleStringOnceCreator struct{}

func (it *newSimpleStringOnceCreator) Empty() *SimpleStringOnce {
	return &SimpleStringOnce{}
}

func (it *newSimpleStringOnceCreator) String(
	isInit bool,
	input string,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		Value: corestr.
			New.
			SimpleStringOnce.
			Create(input, isInit),
	}
}

func (it *newSimpleStringOnceCreator) Initialized(
	input string,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		Value: corestr.
			New.
			SimpleStringOnce.
			Create(
				input,
				true),
	}
}

func (it *newSimpleStringOnceCreator) InitializedWithErrorWrapper(
	input string,
	errorWrapper *errorwrapper.Wrapper,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		Value:        corestr.New.SimpleStringOnce.Create(input, true),
		ErrorWrapper: errorWrapper,
	}
}

func (it *newSimpleStringOnceCreator) Error(
	errType errtype.Variation,
	err error,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newSimpleStringOnceCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newSimpleStringOnceCreator) Create(
	simpleStringOnce corestr.SimpleStringOnce,
	errorWrapper *errorwrapper.Wrapper,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		Value:        simpleStringOnce,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newSimpleStringOnceCreator) Input(
	input string,
	errorWrapper *errorwrapper.Wrapper,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		Value:        corestr.New.SimpleStringOnce.Create(input, input != "" && errorWrapper.IsEmpty()),
		ErrorWrapper: errorWrapper,
	}
}

func (it *newSimpleStringOnceCreator) ValueOnly(
	result string,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		Value: corestr.
			New.
			SimpleStringOnce.
			Init(result),
	}
}
