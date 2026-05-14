package errorwrapper

import (
	"fmt"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/corecsv"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func SimpleReferencesCompile(
	errType errtype.Variation,
	references ...interface{},
) string {
	variantStruct := errType.VariantStructure()
	compiledString := corecsv.AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)

	if compiledString == constants.EmptyString {
		return fmt.Sprintf(
			errconsts.ValueHyphenValueFormat,
			variantStruct.String(),
			variantStruct.Name)
	}

	return fmt.Sprintf(
		errconsts.SimpleReferenceCompileFormat,
		variantStruct.String(),
		variantStruct.Name,
		compiledString)
}
