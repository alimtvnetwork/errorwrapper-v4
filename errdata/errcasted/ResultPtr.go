package errcasted

import (
	"fmt"

	"github.com/alimtvnetwork/core-v9/constants"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type ResultPtr struct {
	Wrapper          *errorwrapper.Wrapper
	IsCastedProperly bool
}

func FailedTypeCastPtr(any interface{}, toType interface{}, msg string) *ResultPtr {
	fromType := fmt.Sprint(constants.SprintTypeFormat, any)
	typeName := fmt.Sprint(constants.SprintTypeFormat, toType)
	msg1 := "From (" + fromType + ") to (" + typeName + "). "

	return &ResultPtr{
		Wrapper:          errnew.Messages.Many(errtype.CastingFailed, msg1, msg),
		IsCastedProperly: false,
	}
}

func EmptyPtr() *ResultPtr {
	return &ResultPtr{
		Wrapper:          nil,
		IsCastedProperly: false,
	}
}

func NewPtr(wrapper *errorwrapper.Wrapper) *ResultPtr {
	return &ResultPtr{
		Wrapper:          wrapper,
		IsCastedProperly: true,
	}
}

func (resultPtr *ResultPtr) ToResult() Result {
	return Result{
		Wrapper:          resultPtr.Wrapper,
		IsCastedProperly: resultPtr.IsCastedProperly,
	}
}
