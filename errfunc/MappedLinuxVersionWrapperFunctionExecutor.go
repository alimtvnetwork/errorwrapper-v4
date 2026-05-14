package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type MappedLinuxVersionWrapperFunctionExecutor struct {
	Mapping map[linuxtype.Variant]WrapperFunc
}

func (l *MappedLinuxVersionWrapperFunctionExecutor) GetWrapperFunctionByType(
	linuxType linuxtype.Variant,
) (
	wrapperFunction WrapperFunc,
	errorWrapper *errorwrapper.Wrapper,
) {
	mappingFunctions := l.Mapping
	wrapperFunc, has := mappingFunctions[linuxType]

	if has {
		return wrapperFunc, nil
	}

	return nil, errnew.Messages.Many(
		errtype.Null,
		"Given field doesn't have any defined function! Field has nil:",
		linuxType.Name())
}

func (l *MappedLinuxVersionWrapperFunctionExecutor) GetWrapperFunctionsByTypes(
	linuxTypes ...linuxtype.Variant,
) (
	wrapperFunctions map[linuxtype.Variant]WrapperFunc,
	errCollection *errwrappers.Collection,
) {
	errCollection = errwrappers.Empty()
	mappingResult := make(map[linuxtype.Variant]WrapperFunc, len(linuxTypes))

	for _, linuxType := range linuxTypes {
		wrapperFunc, errWrapper := l.GetWrapperFunctionByType(linuxType)

		if errWrapper.HasError() {
			errCollection.AddWrapperPtr(errWrapper)

			continue
		}

		mappingResult[linuxType] = wrapperFunc
	}

	return mappingResult, errCollection
}

func (l *MappedLinuxVersionWrapperFunctionExecutor) ExecuteAllAvailableFunctions() *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	for linuxType, wrapperFunc := range l.Mapping {
		newWrapperFunc := convertWrapperWithAdditionalLinuxTypeMessageFunc(
			linuxType,
			errCollection,
			wrapperFunc,
		)

		errCollection.AddAllFunctions(
			newWrapperFunc)
	}

	return errCollection
}

func (l *MappedLinuxVersionWrapperFunctionExecutor) ExecuteByLinuxTypes(
	linuxTypes ...linuxtype.Variant,
) *errwrappers.Collection {
	wrapperFunctionsMapping, errCollection := l.GetWrapperFunctionsByTypes(
		linuxTypes...)

	if errCollection.HasError() {
		return errCollection
	}

	for linuxType, wrapperFunc := range wrapperFunctionsMapping {
		newWrapperFunc := convertWrapperWithAdditionalLinuxTypeMessageFunc(
			linuxType,
			errCollection,
			wrapperFunc,
		)

		errCollection.AddAllFunctions(newWrapperFunc)
	}

	return errCollection
}
