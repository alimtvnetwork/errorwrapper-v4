package errcasted

import (
	"fmt"

	"github.com/alimtvnetwork/core-v9/constants"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type Result struct {
	Wrapper          *errorwrapper.Wrapper
	IsCastedProperly bool
}

func FailedTypeCast(any interface{}, toType interface{}, msg string) Result {
	fromType := fmt.Sprintf(constants.SprintTypeFormat, any)
	typeName := fmt.Sprintf(constants.SprintTypeFormat, toType)
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
