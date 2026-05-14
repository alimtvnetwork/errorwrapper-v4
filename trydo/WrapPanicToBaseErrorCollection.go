package trydo

import (
	"errors"

	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

func WrapPanicToBaseErrorCollection(voidFunc func()) errcoreinf.BaseErrorOrCollectionWrapper {
	finalErr := errwrappers.Empty()

	Block{
		Try: func() {
			voidFunc()
		},
		Catch: func(e Exception) {
			if e == nil {
				return
			}

			switch cast := e.(type) {
			case *errwrappers.Collection:
				finalErr = cast
			case *errorwrapper.Wrapper:
				finalErr.AddWrapperPtr(cast)
			case errorwrapper.BasicErrWrapper:
				finalErr.AddWrapperPtr(cast.AsErrorWrapper())
			case error:
				finalErr.AddError(cast)
			case string:
				finalErr.AddError(errors.New(cast))
			case *corejson.Result:
				finalErr, err := finalErr.ParseInjectUsingJson(cast)

				if err != nil {
					finalErr.AddTypeError(errtype.Unmarshalling, err)
				}
			default:
				finalErr.AddError(errors.New(converters.AnyToString(true, cast)))
			}
		},
		Finally: nil,
	}.Do()

	return finalErr
}
