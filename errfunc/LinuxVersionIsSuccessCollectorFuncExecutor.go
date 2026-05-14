package errfunc

import (
	"reflect"
	"sync"

	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/internal/reflectinternal"
)

type LinuxVersionIsSuccessCollectorFuncExecutor struct {
	sync.Mutex
	mappingFunctions     map[string]IsSuccessCollectorFunc
	Unknown              IsSuccessCollectorFunc
	UbuntuServer         IsSuccessCollectorFunc
	UbuntuServer18       IsSuccessCollectorFunc
	UbuntuServer19       IsSuccessCollectorFunc
	UbuntuServer20       IsSuccessCollectorFunc
	UbuntuServer21       IsSuccessCollectorFunc
	UbuntuServer22       IsSuccessCollectorFunc
	UbuntuServer23       IsSuccessCollectorFunc
	Centos               IsSuccessCollectorFunc
	Centos7              IsSuccessCollectorFunc
	Centos8              IsSuccessCollectorFunc
	Centos9              IsSuccessCollectorFunc
	Centos10             IsSuccessCollectorFunc
	Centos11             IsSuccessCollectorFunc
	Centos12             IsSuccessCollectorFunc
	CentosStream         IsSuccessCollectorFunc
	DebianServer         IsSuccessCollectorFunc
	DebianServer7        IsSuccessCollectorFunc
	DebianServer8        IsSuccessCollectorFunc
	DebianServer9        IsSuccessCollectorFunc
	DebianServer10       IsSuccessCollectorFunc
	DebianServer11       IsSuccessCollectorFunc
	DebianServer12       IsSuccessCollectorFunc
	DebianServer13       IsSuccessCollectorFunc
	DebianServer14       IsSuccessCollectorFunc
	DebianDesktop        IsSuccessCollectorFunc
	Docker               IsSuccessCollectorFunc
	DockerUbuntuServer   IsSuccessCollectorFunc
	DockerUbuntuServer18 IsSuccessCollectorFunc
	DockerUbuntuServer19 IsSuccessCollectorFunc
	DockerUbuntuServer20 IsSuccessCollectorFunc
	DockerUbuntuServer21 IsSuccessCollectorFunc
	DockerUbuntuServer22 IsSuccessCollectorFunc
	DockerCentos7        IsSuccessCollectorFunc
	DockerCentos8        IsSuccessCollectorFunc
	DockerCentos9        IsSuccessCollectorFunc
	DockerCentos10       IsSuccessCollectorFunc
	Android              IsSuccessCollectorFunc
	UbuntuDesktop        IsSuccessCollectorFunc
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) GetFunctionsMappingLock() map[string]IsSuccessCollectorFunc {
	it.Lock()
	defer it.Unlock()

	return it.GetFunctionsMapping()
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) GetFunctionsMapping() map[string]IsSuccessCollectorFunc {
	if it.mappingFunctions != nil {
		return it.mappingFunctions
	}

	hashMapping := make(
		map[string]IsSuccessCollectorFunc,
		len(linuxtype.Ranges))
	it.mappingFunctions = hashMapping

	reflectVal := reflect.ValueOf(it)
	indirectReflectVal := reflect.Indirect(reflectVal)

	for _, linuxTypeName := range linuxtype.Ranges {
		f := indirectReflectVal.FieldByName(linuxTypeName)
		val := reflectinternal.GetFieldValue(f)
		conv, isSuccess := val.(IsSuccessCollectorFunc)

		if isSuccess && conv != nil {
			hashMapping[linuxTypeName] = conv
		}
	}

	return it.mappingFunctions
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) GetWrapperFunctionByType(
	linuxType linuxtype.Variant,
) (
	wrapperFunction IsSuccessCollectorFunc,
	errorWrapper *errorwrapper.Wrapper,
) {
	mappingFunctions := it.GetFunctionsMapping()
	wrapperFunc, has := mappingFunctions[linuxType.Name()]

	if has {
		return wrapperFunc, nil
	}

	return nil, errnew.Messages.Many(
		errtype.Null,
		"Given field doesn't have any defined function! Field has nil:",
		linuxType.Name())
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) GetFunctionsByTypes(
	linuxTypes ...linuxtype.Variant,
) (
	wrapperFunctions map[string]IsSuccessCollectorFunc,
	errCollection *errwrappers.Collection,
) {
	errCollection = errwrappers.Empty()
	mappingResult := make(map[string]IsSuccessCollectorFunc, len(linuxTypes))

	for _, linuxType := range linuxTypes {
		wrapperFunc, errWrapper := it.GetWrapperFunctionByType(linuxType)

		if errWrapper.HasError() {
			errCollection.AddWrapperPtr(errWrapper)

			continue
		}

		mappingResult[linuxType.Name()] = wrapperFunc
	}

	return mappingResult, errCollection
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) ExecuteAllAvailableFunctionsLock() *errwrappers.Collection {
	it.Lock()
	defer it.Unlock()

	return it.ExecuteAllAvailableFunctions()
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) ExecuteAllAvailableFunctions() *errwrappers.Collection {
	mapping := it.GetFunctionsMapping()
	errCollection := errwrappers.Empty()

	for key, wrapperFunc := range mapping {
		newWrapperFunc := convertIsSuccessCollectorFuncWithAdditionalKeyFunc(
			wrapperFunc,
			errCollection,
			key)

		errCollection.AddAllIsSuccessCollectorFunctions(
			newWrapperFunc)
	}

	return errCollection
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) ExecuteByLinuxTypesLock(
	linuxTypes ...linuxtype.Variant,
) *errwrappers.Collection {
	it.Lock()
	defer it.Unlock()

	return it.ExecuteByLinuxTypes(linuxTypes...)
}

func (it *LinuxVersionIsSuccessCollectorFuncExecutor) ExecuteByLinuxTypes(
	linuxTypes ...linuxtype.Variant,
) *errwrappers.Collection {
	wrapperFunctionsMapping, errCollection := it.GetFunctionsByTypes(
		linuxTypes...)

	if errCollection.HasError() {
		return errCollection
	}

	for name, wrapperFunc := range wrapperFunctionsMapping {
		newWrapperFunc := convertIsSuccessCollectorFuncWithAdditionalKeyFunc(
			wrapperFunc,
			errCollection,
			name)

		errCollection.AddAllIsSuccessCollectorFunctions(newWrapperFunc)
	}

	return errCollection
}
