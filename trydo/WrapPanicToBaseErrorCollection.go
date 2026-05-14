package trydo

import (
	"errors"

	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
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
