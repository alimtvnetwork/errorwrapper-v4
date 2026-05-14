package errfunc

import (
	"reflect"
	"sync"

	"gitlab.com/evatix-go/enum/linuxtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/internal/reflectinternal"
)

type LinuxVersionWrapperFunctionExecutor struct {
	sync.Mutex
	mappingFunctions     map[string]WrapperFunc
	Unknown              WrapperFunc
	UbuntuServer         WrapperFunc
	UbuntuServer18       WrapperFunc
	UbuntuServer19       WrapperFunc
	UbuntuServer20       WrapperFunc
	UbuntuServer21       WrapperFunc
	UbuntuServer22       WrapperFunc
	UbuntuServer23       WrapperFunc
	Centos               WrapperFunc
	Centos7              WrapperFunc
	Centos8              WrapperFunc
	Centos9              WrapperFunc
	Centos10             WrapperFunc
	Centos11             WrapperFunc
	Centos12             WrapperFunc
	CentosStream         WrapperFunc
	DebianServer         WrapperFunc
	DebianServer7        WrapperFunc
	DebianServer8        WrapperFunc
	DebianServer9        WrapperFunc
	DebianServer10       WrapperFunc
	DebianServer11       WrapperFunc
	DebianServer12       WrapperFunc
	DebianServer13       WrapperFunc
	DebianServer14       WrapperFunc
	DebianDesktop        WrapperFunc
	Docker               WrapperFunc
	DockerUbuntuServer   WrapperFunc
	DockerUbuntuServer18 WrapperFunc
	DockerUbuntuServer19 WrapperFunc
	DockerUbuntuServer20 WrapperFunc
	DockerUbuntuServer21 WrapperFunc
	DockerUbuntuServer22 WrapperFunc
	DockerCentos7        WrapperFunc
	DockerCentos8        WrapperFunc
	DockerCentos9        WrapperFunc
	DockerCentos10       WrapperFunc
	Android              WrapperFunc
	UbuntuDesktop        WrapperFunc
}

func (l *LinuxVersionWrapperFunctionExecutor) GetFunctionsMappingLock() map[string]WrapperFunc {
	l.Lock()
	defer l.Unlock()

	return l.GetFunctionsMapping()
}

func (l *LinuxVersionWrapperFunctionExecutor) GetFunctionsMapping() map[string]WrapperFunc {
	if l.mappingFunctions != nil {
		return l.mappingFunctions
	}

	hashMapping := make(
		map[string]WrapperFunc,
		len(linuxtype.Ranges))
	l.mappingFunctions = hashMapping

	reflectVal := reflect.ValueOf(l)
	indirectReflectVal := reflect.Indirect(reflectVal)

	for _, linuxTypeName := range linuxtype.Ranges {
		f := indirectReflectVal.FieldByName(linuxTypeName)
		val := reflectinternal.GetFieldValue(f)
		conv, isSuccess := val.(WrapperFunc)

		if isSuccess && conv != nil {
			hashMapping[linuxTypeName] = conv
		}
	}

	return l.mappingFunctions
}

func (l *LinuxVersionWrapperFunctionExecutor) GetWrapperFunctionByType(
	linuxType linuxtype.Variant,
) (
	wrapperFunction WrapperFunc,
	errorWrapper *errorwrapper.Wrapper,
) {
	mappingFunctions := l.GetFunctionsMapping()
	wrapperFunc, has := mappingFunctions[linuxType.Name()]

	if has {
		return wrapperFunc, nil
	}

	return nil, errnew.Messages.Many(
		errtype.Null,
		"Given field doesn't have any defined function! Field has nil:",
		linuxType.Name())
}

func (l *LinuxVersionWrapperFunctionExecutor) GetWrapperFunctionsByTypes(
	linuxTypes ...linuxtype.Variant,
) (
	wrapperFunctions map[string]WrapperFunc,
	errCollection *errwrappers.Collection,
) {
	errCollection = errwrappers.Empty()
	mappingResult := make(map[string]WrapperFunc, len(linuxTypes))

	for _, linuxType := range linuxTypes {
		wrapperFunc, errWrapper := l.GetWrapperFunctionByType(linuxType)

		if errWrapper.HasError() {
			errCollection.AddWrapperPtr(errWrapper)

			continue
		}

		mappingResult[linuxType.Name()] = wrapperFunc
	}

	return mappingResult, errCollection
}

func (l *LinuxVersionWrapperFunctionExecutor) ExecuteAllAvailableFunctionsLock() *errwrappers.Collection {
	l.Lock()
	defer l.Unlock()

	return l.ExecuteAllAvailableFunctions()
}

func (l *LinuxVersionWrapperFunctionExecutor) ExecuteAllAvailableFunctions() *errwrappers.Collection {
	mapping := l.GetFunctionsMapping()
	errCollection := errwrappers.Empty()

	for key, wrapperFunc := range mapping {
		newWrapperFunc := convertWrapperWithAdditionalKeyMessageFunc(
			wrapperFunc,
			errCollection,
			key)

		errCollection.AddAllFunctions(
			newWrapperFunc)
	}

	return errCollection
}

func (l *LinuxVersionWrapperFunctionExecutor) ExecuteByLinuxTypesLock(
	linuxTypes ...linuxtype.Variant,
) *errwrappers.Collection {
	l.Lock()
	defer l.Unlock()

	return l.ExecuteByLinuxTypes(linuxTypes...)
}

func (l *LinuxVersionWrapperFunctionExecutor) ExecuteByLinuxTypes(
	linuxTypes ...linuxtype.Variant,
) *errwrappers.Collection {
	wrapperFunctionsMapping, errCollection := l.GetWrapperFunctionsByTypes(
		linuxTypes...)

	if errCollection.HasError() {
		return errCollection
	}

	for name, wrapperFunc := range wrapperFunctionsMapping {
		newWrapperFunc := convertWrapperWithAdditionalKeyMessageFunc(
			wrapperFunc,
			errCollection,
			name)

		errCollection.AddAllFunctions(newWrapperFunc)
	}

	return errCollection
}
