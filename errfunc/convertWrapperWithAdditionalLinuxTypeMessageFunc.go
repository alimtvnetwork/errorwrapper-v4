package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

func convertWrapperWithAdditionalLinuxTypeMessageFunc(
	variation linuxtype.Variant,
	errCollection *errwrappers.Collection,
	wrapperFunc WrapperFunc,
) func() *errorwrapper.Wrapper {
	return func() *errorwrapper.Wrapper {
		errWrapper := wrapperFunc()

		if errWrapper.HasError() {
			errCollection.AddWrapperWithAdditionalRefs(
				errWrapper,
				ref.Value{
					Variable: "LinuxType",
					Value:    variation.Name(),
				})
		}

		return nil
	}
}
