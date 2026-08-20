package errorwrapper

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/refs"
)

type WrapperDataModel struct {
	IsDisplayableError bool
	CurrentError       string
	ErrorType          errtype.Variation
	StackTraces        codestack.TraceCollection
	References         *refs.Collection
	HasError           bool
}

func NewDataModel(wrapper *Wrapper) WrapperDataModel {
	toModel := WrapperDataModel{}

	if wrapper == nil {
		return toModel
	}

	return transpileWrapperToModel(
		wrapper,
		&toModel)
}
