package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

type Collection struct {
	*corestr.Collection
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Collection) HasError() bool {
	return it.ErrorWrapper != nil &&
		it.ErrorWrapper.HasError()
}

func (it *Collection) IsEmptyError() bool {
	return it.ErrorWrapper != nil &&
		it.ErrorWrapper.HasError()
}

func (it *Collection) IsEmpty() bool {
	return it.Collection == nil || it.Collection.IsEmpty()
}

// HasSafeItems No errors and has items
func (it *Collection) HasSafeItems() bool {
	return it.ErrorWrapper.IsEmpty() &&
		!it.IsEmpty()
}

func (it *Collection) IsSuccess() bool {
	return it.HasSafeItems()
}

func (it *Collection) IsValid() bool {
	return it.HasSafeItems()
}

func (it *Collection) IsFailed() bool {
	return it.HasError()
}
