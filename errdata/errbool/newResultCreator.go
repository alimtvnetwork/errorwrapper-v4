package errbool

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type newResultCreator struct{}

func (it *newResultCreator) Empty() *Result {
	return &Result{}
}

func (it *newResultCreator) Item(
	item bool,
) *Result {
	return &Result{
		Value: item,
	}
}

func (it *newResultCreator) Bool(
	item bool,
) *Result {
	return &Result{
		Value: item,
	}
}

func (it *newResultCreator) True() *Result {
	return &Result{
		Value: true,
	}
}

func (it *newResultCreator) IsSetter(
	value issetter.Value,
) *Result {
	return &Result{
		Value: value.IsTrue(),
	}
}

func (it *newResultCreator) OnOffLater(
	onOff enuminf.OnOffLater,
) *Result {
	return &Result{
		Value: onOff.IsOn(),
	}
}

func (it *newResultCreator) YesNoRejector(
	yesNoRejector coreinterface.YesNoAcceptRejecter,
) *Result {
	return &Result{
		Value: yesNoRejector.IsYes(),
	}
}

func (it *newResultCreator) TrueWithErr(errWp *errorwrapper.Wrapper) *Result {
	return &Result{
		Value:        true,
		ErrorWrapper: errWp,
	}
}

func (it *newResultCreator) False() *Result {
	return &Result{}
}

func (it *newResultCreator) FalseWithErr(errWp *errorwrapper.Wrapper) *Result {
	return &Result{
		ErrorWrapper: errWp,
	}
}

func (it *newResultCreator) Error(
	errType errtype.Variation,
	err error,
) *Result {
	return &Result{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newResultCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) Create(
	result bool,
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		Value:        result,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) ValueOnly(
	result bool,
) *Result {
	return &Result{
		Value: result,
	}
}
