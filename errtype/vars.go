package errtype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coredata/coreonce"
)

var (
	typeName                = coredynamic.SafeTypeName(NoError)
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
