package errfunc

import (
	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

type MappedLinuxVersionIsSuccessCollectorFuncExecutor struct {
	Mapping map[linuxtype.Variant]IsSuccessCollectorFunc
}

func (it *MappedLinuxVersionIsSuccessCollectorFuncExecutor) GetWrapperFunctionByType(
	linuxType linuxtype.Variant,
) (
	wrapperFunction IsSuccessCollectorFunc,
	errorWrapper *errorwrapper.Wrapper,
) {
	wrapperFunc, has := it.Mapping[linuxType]

	if has {
		return wrapperFunc, nil
	}

	return nil, errnew.Messages.Many(
		errtype.Null,
		"Given field doesn't have any defined function! Field has nil:",
		linuxType.Name())
}

func (it *MappedLinuxVersionIsSuccessCollectorFuncExecutor) GetFunctionsByTypes(
	linuxTypes ...linuxtype.Variant,
) (
	wrapperFunctions map[linuxtype.Variant]IsSuccessCollectorFunc,
	errCollection *errwrappers.Collection,
) {
	errCollection = errwrappers.Empty()
	mappingResult := make(map[linuxtype.Variant]IsSuccessCollectorFunc, len(linuxTypes))

	for _, linuxType := range linuxTypes {
		wrapperFunc, errWrapper := it.GetWrapperFunctionByType(linuxType)

		if errWrapper.HasError() {
			errCollection.AddWrapperPtr(errWrapper)

			continue
		}

		mappingResult[linuxType] = wrapperFunc
	}

	return mappingResult, errCollection
}

func (it *MappedLinuxVersionIsSuccessCollectorFuncExecutor) ExecuteAllAvailableFunctions() *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	for linuxType, wrapperFunc := range it.Mapping {
		newWrapperFunc := convertIsSuccessCollectorFuncWithAdditionalLinuxTypeFunc(
			linuxType,
			errCollection,
			wrapperFunc,
		)

		errCollection.AddAllIsSuccessCollectorFunctions(
			newWrapperFunc)
	}

	return errCollection
}

func (it *MappedLinuxVersionIsSuccessCollectorFuncExecutor) ExecuteByLinuxTypes(

	linuxTypes ...linuxtype.Variant,
) *errwrappers.Collection {
	wrapperFunctionsMapping, errCollection := it.GetFunctionsByTypes(
		linuxTypes...)

	if errCollection.HasError() {
		return errCollection
	}

	for linuxType, wrapperFunc := range wrapperFunctionsMapping {
		newWrapperFunc := convertIsSuccessCollectorFuncWithAdditionalLinuxTypeFunc(
			linuxType,
			errCollection,
			wrapperFunc,
		)

		errCollection.AddAllIsSuccessCollectorFunctions(newWrapperFunc)
	}

	return errCollection
}
