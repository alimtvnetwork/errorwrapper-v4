package errany

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newResultApplicable2Creator struct{}

func (it *newResultApplicable2Creator) Empty() *ResultWithApplicable2 {
	return &ResultWithApplicable2{}
}

func (it *newResultApplicable2Creator) Error(
	errType errtype.Variation,
	err error,
) *ResultWithApplicable2 {
	return &ResultWithApplicable2{
		Result2: Result2{
			Result: Result{
				ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
					codestack.Skip1,
					errType,
					err),
			},
		},
	}
}

func (it *newResultApplicable2Creator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *ResultWithApplicable2 {
	return &ResultWithApplicable2{
		Result2: Result2{
			Result: Result{
				ErrorWrapper: errorWrapper,
			},
		},
	}
}

func (it *newResultApplicable2Creator) Create(
	value interface{},
	isApplicable bool,
	errWrapper *errorwrapper.Wrapper,
) *ResultWithApplicable2 {
	return &ResultWithApplicable2{
		Result2: Result2{
			Result: Result{
				Value:        value,
				ErrorWrapper: errWrapper,
			},
		},
		IsApplicable: isApplicable,
	}
}

func (it *newResultApplicable2Creator) ValuesOnly(
	isApplicable bool,
	value interface{},
) *ResultWithApplicable2 {
	return &ResultWithApplicable2{
		Result2: Result2{
			Result: Result{
				Value: value,
			},
		},
		IsApplicable: isApplicable,
	}
}

func (it *newResultApplicable2Creator) ApplicableValue(
	value, value2 interface{},
) *ResultWithApplicable2 {
	return &ResultWithApplicable2{
		Result2: Result2{
			Result: Result{
				Value: value,
			},
			Value2: value2,
		},
		IsApplicable: true,
	}
}

func (it *newResultApplicable2Creator) NonApplicableValue(
	value, value2 interface{},
) *ResultWithApplicable2 {
	return &ResultWithApplicable2{
		Result2: Result2{
			Result: Result{
				Value: value,
			},
			Value2: value2,
		},
		IsApplicable: false,
	}
}
