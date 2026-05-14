package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

func convertWrapperWithAdditionalKeyMessageFunc(
	wrapperFunc WrapperFunc,
	errCollection *errwrappers.Collection,
	key string,
) func() *errorwrapper.Wrapper {
	return func() *errorwrapper.Wrapper {
		errWrapper := wrapperFunc()

		if errWrapper.HasError() {
			errCollection.AddWrapperWithAdditionalRefs(
				errWrapper,
				ref.Value{
					Variable: "LinuxType",
					Value:    key,
				})
		}

		return nil
	}
}
