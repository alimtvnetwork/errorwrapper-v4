package errorwrapper

import (
	"fmt"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/corecsv"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
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
