package errstr

import (
	"github.com/alimtvnetwork/core-v9/coredata/corestr"

	"github.com/alimtvnetwork/errorwrapper-v4"
)

type CharCollectionMap struct {
	*corestr.CharCollectionMap
	ErrorWrapper *errorwrapper.Wrapper
}
