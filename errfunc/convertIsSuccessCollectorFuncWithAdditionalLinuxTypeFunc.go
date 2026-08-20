package errfunc

import (
	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

func convertIsSuccessCollectorFuncWithAdditionalLinuxTypeFunc(
	linuxType linuxtype.Variant,
	errCollection *errwrappers.Collection,
	isSuccessCollectorFunc IsSuccessCollectorFunc,
) IsSuccessCollectorFunc {
	return func(errorCollection *errwrappers.Collection) (isSuccess bool) {
		isSuccess = isSuccessCollectorFunc(errorCollection)

		if !isSuccess {
			first := errorCollection.First()
			errCollection.AddWrapperWithAdditionalRefs(
				first,
				ref.Value{
					Variable: "LinuxType",
					Value:    linuxType.Name(),
				})
		}

		return isSuccess
	}
}
