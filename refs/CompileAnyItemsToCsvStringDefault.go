package refs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecsv"
)

// CompileAnyItemsToCsvStringDefault
//
// if references empty or len 0 then empty string returned.
//
// Final join whole lines with the joiner given (... joiner item)
//
// Formats :
//  - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
func CompileAnyItemsToCsvStringDefault(
	references ...interface{},
) string {
	return corecsv.AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
