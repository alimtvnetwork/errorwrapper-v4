package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/errorwrapper"
)

type CharCollectionMap struct {
	*corestr.CharCollectionMap
	ErrorWrapper *errorwrapper.Wrapper
}
