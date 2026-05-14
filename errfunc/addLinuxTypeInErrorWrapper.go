package errfunc

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/ref"
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
