package errorwrapper

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/refs"
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
