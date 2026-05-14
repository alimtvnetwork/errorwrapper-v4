package errcasted

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type Result struct {
	Wrapper          *errorwrapper.Wrapper
	IsCastedProperly bool
}

func FailedTypeCast(any interface{}, toType interface{}, msg string) Result {
	fromType := fmt.Sprint(constants.SprintTypeFormat, any)
	typeName := fmt.Sprint(constants.SprintTypeFormat, toType)
	msg1 := "From (" + fromType + ") to (" + typeName + "). "

	return Result{
		Wrapper:          errnew.Messages.Many(errtype.CastingFailed, msg1, msg),
		IsCastedProperly: false,
	}
}

func New(wrapper *errorwrapper.Wrapper) Result {
	return Result{
		Wrapper:          wrapper,
		IsCastedProperly: true,
	}
}

func Empty() Result {
	return Result{
		Wrapper:          errorwrapper.EmptyPtr(),
		IsCastedProperly: false,
	}
}

func (it Result) ToResultPtr() *ResultPtr {
	return &ResultPtr{
		Wrapper:          it.Wrapper,
		IsCastedProperly: it.IsCastedProperly,
	}
}
