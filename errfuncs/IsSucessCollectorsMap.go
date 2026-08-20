package errfuncs

import (
	"sync"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errcmd"
	"github.com/alimtvnetwork/errorwrapper-v4/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

type IsSuccessCollectorsMap struct {
	items *map[string]errfunc.IsSuccessCollectorFunc
	sync.Mutex
}

func NewIsSuccessCollectorsMap(capacity int) *IsSuccessCollectorsMap {
	items := make(
		map[string]errfunc.IsSuccessCollectorFunc,
		capacity)

	return &IsSuccessCollectorsMap{items: &items}
}

func (it *IsSuccessCollectorsMap) Length() int {
	if it.items == nil {
		return 0
	}

	return len(*it.items)
}

func (it *IsSuccessCollectorsMap) Count() int {
	return it.Length()
}

func (it *IsSuccessCollectorsMap) IsEmpty() bool {
	return it.Length() == 0
}

func (it *IsSuccessCollectorsMap) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *IsSuccessCollectorsMap) LastIndex() int {
	return it.Length() - 1
}

func (it *IsSuccessCollectorsMap) HasIndex(
	index int,
) bool {
	return it.LastIndex() >= index
}

func (it *IsSuccessCollectorsMap) Items() *map[string]errfunc.IsSuccessCollectorFunc {
	return it.items
}

func (it *IsSuccessCollectorsMap) Get(
	key string,
) errfunc.IsSuccessCollectorFunc {
	return (*it.items)[key]
}

func (it *IsSuccessCollectorsMap) IsKeyExist(
	key string,
) bool {
	_, has := (*it.items)[key]

	return has
}

func (it *IsSuccessCollectorsMap) IsAllKeysExist(
	keys ...string,
) bool {
	for _, key := range keys {
		_, has := (*it.items)[key]

		if !has {
			return false
		}
	}

	return true
}

func (it *IsSuccessCollectorsMap) GetItemsByKeys(
	keys ...string,
) []errfunc.IsSuccessCollectorFunc {
	slice := make(
		[]errfunc.IsSuccessCollectorFunc,
		constants.Zero,
		len(keys))

	if len(keys) == 0 {
		return slice
	}

	for _, key := range keys {
		item := it.Get(key)

		if item != nil {
			slice = append(slice, item)
		}
	}

	return slice
}

func (it *IsSuccessCollectorsMap) Add(
	key string,
	wrapperFunc errfunc.IsSuccessCollectorFunc,
) *IsSuccessCollectorsMap {
	if wrapperFunc == nil {
		return it
	}

	(*it.items)[key] = wrapperFunc

	return it
}

func (it *IsSuccessCollectorsMap) AddLinuxIsSuccessPlusErrorCollectActionsIf(
	isAdd bool,
	processorActions ...errfunc.LinuxIsSuccessPlusErrorCollectAction,
) *IsSuccessCollectorsMap {
	if !isAdd || len(processorActions) == 0 {
		return it
	}

	for _, action := range processorActions {
		it.Add(
			action.LinuxType.Name(),
			action.IsSuccessCollectorFunc)
	}

	return it
}

func (it *IsSuccessCollectorsMap) AddLinuxIsSuccessCollectorFunc(
	linuxType linuxtype.Variant,
	isSuccessCollectorFunc errfunc.IsSuccessCollectorFunc,
) *IsSuccessCollectorsMap {
	return it.Add(
		linuxType.Name(),
		isSuccessCollectorFunc)
}

func (it *IsSuccessCollectorsMap) AddLinuxWrapperFunc(
	linuxType linuxtype.Variant,
	wrapperFunc errfunc.WrapperFunc,
) *IsSuccessCollectorsMap {
	return it.Add(
		linuxType.Name(),
		errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(
			wrapperFunc))
}

func (it *IsSuccessCollectorsMap) AddLinuxWrapperAction(
	linuxWrapperAction errfunc.LinuxWrapperAction,
) *IsSuccessCollectorsMap {
	return it.Add(
		linuxWrapperAction.LinuxType.Name(),
		errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(
			linuxWrapperAction.WrapperFunc))
}

func (it *IsSuccessCollectorsMap) ExecuteByKeysAndCollection(
	errorWrapperCollection *errwrappers.Collection,
	keys ...string,
) *errwrappers.Collection {
	if it.IsEmpty() {
		return errorWrapperCollection
	}

	for _, key := range keys {
		isSuccessCollectorFunc, has := (*it.items)[key]

		if !has {
			errorWrapperCollection.AddRefOne(
				errtype.DictionaryNotContainsKey,
				"key",
				key)
		}

		errorWrapperCollection.AddIsSuccessCollectorWithKey(
			key,
			isSuccessCollectorFunc)
	}

	return errorWrapperCollection
}

func (it *IsSuccessCollectorsMap) IsAllLinuxTypesExist(
	linuxTypes ...linuxtype.Variant,
) bool {
	if linuxTypes == nil {
		return true
	}

	keys := errfunc.LinuxTypesToNameSlice(
		linuxTypes...)

	return it.IsAllKeysExist(keys...)
}

func (it *IsSuccessCollectorsMap) ExecuteByLinuxTypes(
	linuxTypes ...linuxtype.Variant,
) *errwrappers.Collection {
	if linuxTypes == nil {
		return errwrappers.Empty()
	}

	keys := errfunc.LinuxTypesToNameSlice(
		linuxTypes...)

	return it.ExecuteByKeys(
		keys...)
}

func (it *IsSuccessCollectorsMap) AddLinuxTypeCmdOnce(
	linuxType linuxtype.Variant,
	cmdOnce *errcmd.CmdOnce,
) *IsSuccessCollectorsMap {
	if cmdOnce == nil {
		return it
	}

	return it.Add(
		linuxType.Name(),
		func(errorCollector *errwrappers.Collection) (isSuccess bool) {
			errorWrapper := cmdOnce.CompiledErrorWrapper()
			errorCollector.AddWrapperPtr(errorWrapper)

			return errorWrapper.IsEmptyError()
		})
}

func (it *IsSuccessCollectorsMap) AddLinuxTypeCmdOnceUsingProcessor(
	linuxType linuxtype.Variant,
	cmdOnce *errcmd.CmdOnce,
	cmdOnceProcessor func(cmdOnce *errcmd.CmdOnce) *errorwrapper.Wrapper,
) *IsSuccessCollectorsMap {
	if cmdOnce == nil {
		return it
	}

	return it.Add(
		linuxType.Name(),
		func(errorCollector *errwrappers.Collection) (isSuccess bool) {
			errorWrapper := cmdOnceProcessor(cmdOnce)
			errorCollector.AddWrapperPtr(errorWrapper)

			return errorWrapper.IsEmptyError()
		})
}

func (it *IsSuccessCollectorsMap) ExecuteByKeys(
	keys ...string,
) *errwrappers.Collection {
	errorWrapperCollection := errwrappers.Empty()

	return it.ExecuteByKeysAndCollection(
		errorWrapperCollection,
		keys...)
}

func (it *IsSuccessCollectorsMap) ExecuteAll() *errorwrapper.Wrapper {
	return it.
		ExecuteAllByDefault().
		GetAsErrorWrapperPtr()
}

func (it *IsSuccessCollectorsMap) ExecuteAllUsingCollection(
	errorWrapperCollection *errwrappers.Collection,
) *errwrappers.Collection {
	if it.IsEmpty() {
		return errorWrapperCollection
	}

	for key, isSuccessCollectorFunc := range *it.items {
		errorWrapperCollection.AddIsSuccessCollectorWithKey(
			key,
			isSuccessCollectorFunc,
		)
	}

	return errorWrapperCollection
}

func (it *IsSuccessCollectorsMap) ExecuteAllByDefault() *errwrappers.Collection {
	errorWrapperCollection := errwrappers.Empty()

	return it.ExecuteAllUsingCollection(errorWrapperCollection)
}

func (it IsSuccessCollectorsMap) String() string {
	return it.
		ExecuteAllByDefault().
		String()
}

func (it *IsSuccessCollectorsMap) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
