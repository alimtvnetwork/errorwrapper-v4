package errstr

import (
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v3"
)

type Collection struct {
	*corestr.Collection
	ErrorWrapper *errorwrapper.Wrapper
}

func (it *Collection) HasError() bool {
	return it != nil && it.ErrorWrapper != nil &&
		it.ErrorWrapper.HasError()
}

func (it *Collection) IsEmptyError() bool {
	return it == nil || it.ErrorWrapper == nil ||
		it.ErrorWrapper.IsEmpty()
}

func (it *Collection) IsEmpty() bool {
	return it == nil || it.Collection == nil || it.Collection.IsEmpty()
}

// HasSafeItems No errors and has items
func (it *Collection) HasSafeItems() bool {
	return it != nil && it.ErrorWrapper.IsEmpty() &&
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
