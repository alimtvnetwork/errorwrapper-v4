package errfunc

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errbool"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errstr"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
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
