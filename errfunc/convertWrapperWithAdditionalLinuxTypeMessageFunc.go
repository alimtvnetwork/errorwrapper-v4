package errfunc

import (
	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
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
