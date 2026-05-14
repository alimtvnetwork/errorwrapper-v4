package errtype

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/coreonce"
)

var (
	typeName                = coredynamic.TypeName(NoError)
	rangesCsvNameStringOnce = coreonce.NewStringOnce(func() string {
		return generateRangesCsvString()
	})
	stringToVariantMapOnce = coreonce.NewAnyOnce(func() interface{} {
		return generateAllErrorTypeMap()
	})

	allNameWithValuesOnce = coreonce.NewStringsOnce(func() []string {
		return generateAllNameWithValues()
	})
)
