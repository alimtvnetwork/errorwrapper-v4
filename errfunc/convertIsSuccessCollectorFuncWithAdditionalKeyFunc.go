package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

func convertIsSuccessCollectorFuncWithAdditionalKeyFunc(
	isSuccessCollectorFunc IsSuccessCollectorFunc,
	errCollection *errwrappers.Collection,
	key string,
) IsSuccessCollectorFunc {
	return func(errorCollection *errwrappers.Collection) (isSuccess bool) {
		isSuccess = isSuccessCollectorFunc(errorCollection)

		if !isSuccess {
			first := errorCollection.First()
			errCollection.AddWrapperWithAdditionalRefs(
				first,
				ref.Value{
					Variable: "LinuxType",
					Value:    key,
				})
		}

		return isSuccess
	}
}
