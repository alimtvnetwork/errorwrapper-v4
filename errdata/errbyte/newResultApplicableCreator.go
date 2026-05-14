package errbyte

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newResultApplicableCreator struct{}

func (it *newResultApplicableCreator) Empty() *ResultWithApplicable {
	return &ResultWithApplicable{}
}

func (it *newResultApplicableCreator) Error(
	errType errtype.Variation,
	err error,
) *ResultWithApplicable {
	return &ResultWithApplicable{
		Result: Result{
			ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
				codestack.Skip1,
				errType,
				err),
		},
	}
}

func (it *newResultApplicableCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *ResultWithApplicable {
	return &ResultWithApplicable{
		Result: Result{
			ErrorWrapper: errorWrapper,
		},
	}
}

func (it *newResultApplicableCreator) Create(
	value byte,
	isApplicable bool,
	errWrapper *errorwrapper.Wrapper,
) *ResultWithApplicable {
	return &ResultWithApplicable{
		Result: Result{
			Value:        value,
			ErrorWrapper: errWrapper,
		},
		IsApplicable: isApplicable,
	}
}

func (it *newResultApplicableCreator) ValuesOnly(
	isApplicable bool,
	value byte,
) *ResultWithApplicable {
	return &ResultWithApplicable{
		Result: Result{
			Value: value,
		},
		IsApplicable: isApplicable,
	}
}

func (it *newResultApplicableCreator) ApplicableZero() *ResultWithApplicable {
	return &ResultWithApplicable{
		IsApplicable: true,
	}
}

func (it *newResultApplicableCreator) ApplicableOne() *ResultWithApplicable {
	return &ResultWithApplicable{
		Result: Result{
			Value: constants.One,
		},
		IsApplicable: true,
	}
}

func (it *newResultApplicableCreator) ApplicableValue(
	value byte,
) *ResultWithApplicable {
	return &ResultWithApplicable{
		Result: Result{
			Value: value,
		},
		IsApplicable: true,
	}
}

func (it *newResultApplicableCreator) NonApplicableValue(
	value byte,
) *ResultWithApplicable {
	return &ResultWithApplicable{
		Result: Result{
			Value: value,
		},
		IsApplicable: false,
	}
}
