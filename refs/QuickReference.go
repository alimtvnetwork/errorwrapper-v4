package refs

import (
	"github.com/alimtvnetwork/core-v9/corecsv"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
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
