package errfuncs

import (
	"sync"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

type WrappersMap struct {
	items *map[string]errfunc.WrapperFunc
	sync.Mutex
}

func NewWrappersMap(capacity int) *WrappersMap {
	items := make(
		map[string]errfunc.WrapperFunc,
		capacity)

	return &WrappersMap{items: &items}
}

func (it *WrappersMap) Length() int {
	if it.items == nil {
		return 0
	}

	return len(*it.items)
}

func (it *WrappersMap) Count() int {
	return it.Length()
}

func (it *WrappersMap) IsEmpty() bool {
	return it.Length() == 0
}

func (it *WrappersMap) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *WrappersMap) LastIndex() int {
	return it.Length() - 1
}

func (it *WrappersMap) HasIndex(
	index int,
) bool {
	return it.LastIndex() >= index
}

func (it *WrappersMap) Items() *map[string]errfunc.WrapperFunc {
	return it.items
}

func (it *WrappersMap) Get(
	key string,
) errfunc.WrapperFunc {
	return (*it.items)[key]
}

func (it *WrappersMap) IsKeyExist(
	key string,
) bool {
	_, has := (*it.items)[key]

	return has
}

func (it *WrappersMap) IsAllKeysExist(
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

func (it *WrappersMap) GetItemsByKeys(
	keys ...string,
) []errfunc.WrapperFunc {
	slice := make(
		[]errfunc.WrapperFunc,
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

func (it *WrappersMap) Add(
	key string,
	wrapperFunc errfunc.WrapperFunc,
) *WrappersMap {
	if wrapperFunc == nil {
		return it
	}

	(*it.items)[key] = wrapperFunc

	return it
}

func (it *WrappersMap) AddLinuxIsSuccessPlusErrorCollectActionsIf(
	isAdd bool,
	processorActions ...errfunc.LinuxIsSuccessPlusErrorCollectAction,
) *WrappersMap {
	if !isAdd || len(processorActions) == 0 {
		return it
	}

	for _, action := range processorActions {
		wrapperFunc := errfunc.ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc(
			action)

		it.Add(
			action.LinuxType.Name(),
			wrapperFunc)
	}

	return it
}

func (it *WrappersMap) AddLinuxWrapperActions(
	linuxType linuxtype.Variant,
	wrapperFunc errfunc.WrapperFunc,
) *WrappersMap {
	return it.Add(
		linuxType.Name(),
		wrapperFunc)
}

func (it *WrappersMap) AddLinuxWrapperActionsIf(
	isAdd bool,
	processorActions ...errfunc.LinuxWrapperAction,
) *WrappersMap {
	if !isAdd || len(processorActions) == 0 {
		return it
	}

	for _, action := range processorActions {
		it.Add(
			action.LinuxType.Name(),
			action.WrapperFunc)
	}

	return it
}

func (it *WrappersMap) AddLinuxErrorCollectorActionsIf(
	isAdd bool,
	processorActions ...errfunc.LinuxErrorCollectorAction,
) *WrappersMap {
	if !isAdd || len(processorActions) == 0 {
		return it
	}

	for _, action := range processorActions {
		wrapperFunc := errfunc.ConvertLinuxErrorCollectorActionToErrWrapperFunc(
			action)

		it.Add(action.LinuxType.Name(), wrapperFunc)
	}

	return it
}

func (it *WrappersMap) ExecuteByKeysAndCollection(
	errorWrapperCollection *errwrappers.Collection,
	keys ...string,
) *errwrappers.Collection {
	if errorWrapperCollection == nil {
		errorWrapperCollection = errwrappers.Empty()
	}

	if it.IsEmpty() && len(keys) == 0 {
		return errorWrapperCollection
	}

	for _, key := range keys {
		errorWrapperFunc, has := (*it.items)[key]

		if !has {
			errorWrapperCollection.AddRefOne(
				errtype.DictionaryNotContainsKey,
				"key",
				key)
		}

		errorWrapperCollection.AddAnyFunctions(
			errorWrapperFunc)
	}

	return errorWrapperCollection
}

func (it *WrappersMap) IsAllLinuxTypesExist(
	linuxTypes ...linuxtype.Variant,
) bool {
	if linuxTypes == nil {
		return true
	}

	keys := errfunc.LinuxTypesToNameSlice(
		linuxTypes...)

	return it.IsAllKeysExist(keys...)
}

func (it *WrappersMap) ExecuteByLinuxTypes(
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

func (it *WrappersMap) AddLinuxTypeCmdOnce(
	linuxType linuxtype.Variant,
	cmdOnce *errcmd.CmdOnce,
) *WrappersMap {
	if cmdOnce == nil {
		return it
	}

	return it.Add(linuxType.Name(), func() *errorwrapper.Wrapper {
		return cmdOnce.CompiledErrorWrapper()
	})
}

func (it *WrappersMap) AddLinuxTypeCmdOnceUsingProcessor(
	linuxType linuxtype.Variant,
	cmdOnce *errcmd.CmdOnce,
	cmdOnceProcessor func(cmdOnce *errcmd.CmdOnce) *errorwrapper.Wrapper,
) *WrappersMap {
	if cmdOnce == nil {
		return it
	}

	return it.Add(linuxType.Name(), func() *errorwrapper.Wrapper {
		return cmdOnceProcessor(cmdOnce)
	})
}

func (it *WrappersMap) ExecuteByKeys(
	keys ...string,
) *errwrappers.Collection {
	errorWrapperCollection := errwrappers.Empty()

	return it.ExecuteByKeysAndCollection(
		errorWrapperCollection,
		keys...)
}

func (it *WrappersMap) ExecuteAll() *errorwrapper.Wrapper {
	return it.
		ExecuteAllByDefault().
		GetAsErrorWrapperPtr()
}

func (it *WrappersMap) ExecuteAllUsingCollection(
	errorWrapperCollection *errwrappers.Collection,
) *errwrappers.Collection {
	if it.IsEmpty() {
		return errorWrapperCollection
	}

	for _, wrapperFunc := range *it.items {
		errorWrapper := wrapperFunc()

		errorWrapperCollection.AddWrapperPtr(errorWrapper)
	}

	return errorWrapperCollection
}

func (it *WrappersMap) ExecuteAllByDefault() *errwrappers.Collection {
	errorWrapperCollection := errwrappers.Empty()

	return it.ExecuteAllUsingCollection(errorWrapperCollection)
}

func (it WrappersMap) String() string {
	return it.
		ExecuteAllByDefault().
		String()
}

func (it *WrappersMap) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
