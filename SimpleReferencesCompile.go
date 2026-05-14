package errorwrapper

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecsv"
	"gitlab.com/evatix-go/errorwrapper/errconsts"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
