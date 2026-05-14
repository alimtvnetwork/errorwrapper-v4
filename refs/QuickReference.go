package refs

import (
	"gitlab.com/evatix-go/core/corecsv"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type QuickReference struct {
	ErrorType              errtype.Variation
	ReferencesCompiledLine string
}

func (it QuickReference) CompileLine() string {
	return it.ErrorType.CodeTypeNameWithReference(it.ReferencesCompiledLine)
}

func NewQuickReference(
	errorType errtype.Variation,
	referenceItems ...interface{},
) QuickReference {
	return QuickReference{
		ErrorType:              errorType,
		ReferencesCompiledLine: corecsv.AnyItemsToStringDefault(referenceItems...),
	}
}

func NewQuickReferenceStrings(
	errorType errtype.Variation,
	referenceItems ...string,
) QuickReference {
	return QuickReference{
		ErrorType:              errorType,
		ReferencesCompiledLine: corecsv.StringsToStringDefault(referenceItems...),
	}
}
