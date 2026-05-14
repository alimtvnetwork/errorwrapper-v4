package errorwrapper

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecsv"
	"gitlab.com/evatix-go/errorwrapper/errconsts"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// SimpleReferencesCompileOptimized
//
// errconsts.SimpleReferenceCompileOptimizedFormat = `%typeName (..., "reference")`
func SimpleReferencesCompileOptimized(
	errType errtype.Variation,
	references ...interface{},
) string {
	variantStruct := errType.VariantStructurePtr()
	compiledString := corecsv.AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)

	if compiledString == constants.EmptyString {
		return variantStruct.Name
	}

	return fmt.Sprintf(
		errconsts.SimpleReferenceCompileOptimizedFormat,
		variantStruct.Name,
		compiledString)
}
