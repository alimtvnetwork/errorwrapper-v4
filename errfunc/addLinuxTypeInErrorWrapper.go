package errfunc

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
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
