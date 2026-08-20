package errfuncs

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

type Collectors struct {
	items *[]errfunc.CollectorFunc
}

func NewCollectors(
	capacity int,
) *Collectors {
	items := make(
		[]errfunc.CollectorFunc,
		constants.Zero,
		capacity)

	return &Collectors{items: &items}
}

func (it *Collectors) Length() int {
	if it.items == nil {
		return 0
	}

	return len(*it.items)
}

func (it *Collectors) Count() int {
	return it.Length()
}

func (it *Collectors) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Collectors) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Collectors) LastIndex() int {
	return it.Length() - 1
}

func (it *Collectors) HasIndex(
	index int,
) bool {
	return it.LastIndex() >= index
}

func (it *Collectors) Items() *[]errfunc.CollectorFunc {
	return it.items
}

func (it *Collectors) GetAt(
	index int,
) errfunc.CollectorFunc {
	return (*it.items)[index]
}

func (it *Collectors) GetSafeAt(
	index int,
) errfunc.CollectorFunc {
	if it.HasIndex(index) {
		return (*it.items)[index]
	}

	return nil
}

func (it *Collectors) Add(
	collectorFunc errfunc.CollectorFunc,
) *Collectors {
	if collectorFunc == nil {
		return it
	}

	*it.items = append(
		*it.items,
		collectorFunc)

	return it
}

func (it *Collectors) AddsIf(
	isAdd bool,
	errorCollectorFunctions ...errfunc.CollectorFunc,
) *Collectors {
	if !isAdd {
		return it
	}

	return it.Adds(errorCollectorFunctions...)
}

func (it *Collectors) Adds(
	errorCollectorFunctions ...errfunc.CollectorFunc,
) *Collectors {
	if len(errorCollectorFunctions) == 0 {
		return it
	}

	for _, errorCollectorFunction := range errorCollectorFunctions {
		if errorCollectorFunction == nil {
			continue
		}

		*it.items = append(
			*it.items,
			errorCollectorFunction)
	}

	return it
}

func (it *Collectors) ExecuteAll() *errorwrapper.Wrapper {
	return it.
		ExecuteAllCollectionWithNewEmpty().
		GetAsErrorWrapperPtr()
}

func (it *Collectors) ExecuteAllCollectionWithNewEmpty() *errwrappers.Collection {
	return it.ExecuteAllCollection(errwrappers.Empty())
}

func (it *Collectors) ExecuteAllCollection(
	errorCollection *errwrappers.Collection,
) *errwrappers.Collection {

	if it.IsEmpty() {
		return errorCollection
	}

	for _, errorCollectorFunc := range *it.items {
		errorCollectorFunc(errorCollection)
	}

	return errorCollection
}

func (it Collectors) String() string {
	return it.
		ExecuteAllCollectionWithNewEmpty().
		String()
}

func (it *Collectors) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
