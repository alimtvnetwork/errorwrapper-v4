package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/ref"
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
