package errstr

import (
	"github.com/alimtvnetwork/core-v9/coredata/corestr"

	"github.com/alimtvnetwork/errorwrapper-v3"
)

// LinkedCollections TODO constructors
type LinkedCollections struct {
	*corestr.LinkedCollections
	ErrorWrapper *errorwrapper.Wrapper
}
