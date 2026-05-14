package errfuncs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errfunc"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type Wrappers struct {
	items *[]errfunc.WrapperFunc
}

func NewWrappers(capacity int) *Wrappers {
	items := make(
		[]errfunc.WrapperFunc,
		constants.Zero,
		capacity)

	return &Wrappers{items: &items}
}

func (it *Wrappers) Length() int {
	if it.items == nil {
		return 0
	}

	return len(*it.items)
}

func (it *Wrappers) Count() int {
	return it.Length()
}

func (it *Wrappers) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Wrappers) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Wrappers) LastIndex() int {
	return it.Length() - 1
}

func (it *Wrappers) HasIndex(
	index int,
) bool {
	return it.LastIndex() >= index
}

func (it *Wrappers) Items() *[]errfunc.WrapperFunc {
	return it.items
}

func (it *Wrappers) GetAt(
	index int,
) errfunc.WrapperFunc {
	return (*it.items)[index]
}

func (it *Wrappers) GetSafeAt(
	index int,
) errfunc.WrapperFunc {
	if it.HasIndex(index) {
		return (*it.items)[index]
	}

	return nil
}

func (it *Wrappers) Add(
	wrapperFunc errfunc.WrapperFunc,
) *Wrappers {
	if wrapperFunc == nil {
		return it
	}

	*it.items = append(
		*it.items,
		wrapperFunc)

	return it
}

func (it *Wrappers) AddsIf(
	isAdd bool,
	wrapperFunctions ...errfunc.WrapperFunc,
) *Wrappers {
	if !isAdd {
		return it
	}

	return it.Adds(wrapperFunctions...)
}

func (it *Wrappers) SimpleErrorFunctionsAdds(
	errorType errtype.Variation,
	simpleErrorFunctions ...errfunc.SimpleErrorFunc,
) *Wrappers {
	if len(simpleErrorFunctions) == 0 {
		return it
	}

	for _, simpleErrorFunc := range simpleErrorFunctions {
		wrapperFunc := errfunc.ConvertErrorFuncToWrapper(
			errorType,
			simpleErrorFunc)

		it.Add(wrapperFunc)
	}

	return it
}

func (it *Wrappers) Adds(
	wrapperFunctions ...errfunc.WrapperFunc,
) *Wrappers {
	if len(wrapperFunctions) == 0 {
		return it
	}

	for _, wrapperFunc := range wrapperFunctions {
		if wrapperFunc == nil {
			continue
		}

		*it.items = append(
			*it.items,
			wrapperFunc)
	}

	return it
}

func (it *Wrappers) ExecuteAll() *errorwrapper.Wrapper {
	return it.
		ExecuteAllCollection().
		GetAsErrorWrapperPtr()
}

func (it *Wrappers) ExecuteAllCollection() *errwrappers.Collection {
	errorWrapperCollection := errwrappers.Empty()

	if it.IsEmpty() {
		return errorWrapperCollection
	}

	for _, wrapperFunc := range *it.items {
		errorWrapper := wrapperFunc()

		errorWrapperCollection.AddWrapperPtr(errorWrapper)
	}

	return errorWrapperCollection
}

func (it Wrappers) String() string {
	return it.
		ExecuteAllCollection().
		String()
}

func (it *Wrappers) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
