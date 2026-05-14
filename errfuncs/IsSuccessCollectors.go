package errfuncs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errfunc"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type IsSuccessCollectors struct {
	items *[]errfunc.IsSuccessCollectorFunc
}

func NewIsSuccessCollectors(capacity int) *IsSuccessCollectors {
	items := make(
		[]errfunc.IsSuccessCollectorFunc,
		constants.Zero,
		capacity)

	return &IsSuccessCollectors{items: &items}
}

func (receiver *IsSuccessCollectors) Length() int {
	if receiver.items == nil {
		return 0
	}

	return len(*receiver.items)
}

func (receiver *IsSuccessCollectors) Count() int {
	return receiver.Length()
}

func (receiver *IsSuccessCollectors) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *IsSuccessCollectors) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *IsSuccessCollectors) LastIndex() int {
	return receiver.Length() - 1
}

func (receiver *IsSuccessCollectors) HasIndex(
	index int,
) bool {
	return receiver.LastIndex() >= index
}

func (receiver *IsSuccessCollectors) Items() *[]errfunc.IsSuccessCollectorFunc {
	return receiver.items
}

func (receiver *IsSuccessCollectors) GetAt(
	index int,
) errfunc.IsSuccessCollectorFunc {
	return (*receiver.items)[index]
}

func (receiver *IsSuccessCollectors) GetSafeAt(
	index int,
) errfunc.IsSuccessCollectorFunc {
	if receiver.HasIndex(index) {
		return (*receiver.items)[index]
	}

	return nil
}

func (receiver *IsSuccessCollectors) Add(
	wrapperFunc errfunc.IsSuccessCollectorFunc,
) *IsSuccessCollectors {
	if wrapperFunc == nil {
		return receiver
	}

	*receiver.items = append(
		*receiver.items,
		wrapperFunc)

	return receiver
}

func (receiver *IsSuccessCollectors) AddsIf(
	isAdd bool,
	wrapperFunctions ...errfunc.IsSuccessCollectorFunc,
) *IsSuccessCollectors {
	if !isAdd {
		return receiver
	}

	return receiver.Adds(wrapperFunctions...)
}

func (receiver *IsSuccessCollectors) SimpleErrorFunctionsAdds(
	errorType errtype.Variation,
	simpleErrorFunctions ...errfunc.SimpleErrorFunc,
) *IsSuccessCollectors {
	if len(simpleErrorFunctions) == 0 {
		return receiver
	}

	for _, simpleErrorFunc := range simpleErrorFunctions {
		isSuccessCollectorFunc := errfunc.ConvertErrorFuncToIsSuccessCollectorFunc(
			errorType,
			simpleErrorFunc)

		receiver.Add(isSuccessCollectorFunc)
	}

	return receiver
}

func (receiver *IsSuccessCollectors) Adds(
	wrapperFunctions ...errfunc.IsSuccessCollectorFunc,
) *IsSuccessCollectors {
	if len(wrapperFunctions) == 0 {
		return receiver
	}

	for _, wrapperFunc := range wrapperFunctions {
		if wrapperFunc == nil {
			continue
		}

		*receiver.items = append(
			*receiver.items,
			wrapperFunc)
	}

	return receiver
}

func (receiver *IsSuccessCollectors) ExecuteAll() *errorwrapper.Wrapper {
	return receiver.
		ExecuteAllCollection().
		GetAsErrorWrapperPtr()
}

func (receiver *IsSuccessCollectors) ExecuteAllCollection() *errwrappers.Collection {
	errorWrapperCollection := errwrappers.Empty()

	if receiver.IsEmpty() {
		return errorWrapperCollection
	}

	for _, wrapperFunc := range *receiver.items {
		wrapperFunc(errorWrapperCollection)
	}

	return errorWrapperCollection
}

func (receiver *IsSuccessCollectors) String() string {
	return receiver.
		ExecuteAllCollection().
		String()
}

func (receiver *IsSuccessCollectors) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return receiver
}
