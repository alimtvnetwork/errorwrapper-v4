package errfunc

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type (
	SimpleErrorFunc                 func() error
	WrapperFunc                     func() *errorwrapper.Wrapper
	BoolReturnFunc                  func() *errbool.Result
	CollectionReturnFunc            func() *errstr.Collection
	CollectorFunc                   func(errorCollection *errwrappers.Collection)
	IsSuccessCollectorFunc          func(errorCollection *errwrappers.Collection) (isSuccess bool)
	IsSuccessProcessorCollectorFunc func(
		dynamicIn coredynamic.Dynamic,
		errorCollection *errwrappers.Collection,
	) (isSuccess bool)
	ConvertErrorFuncToWrapperFunc func(
		errorType errtype.Variation,
		simpleErrFunc SimpleErrorFunc,
	) WrapperFunc
)
