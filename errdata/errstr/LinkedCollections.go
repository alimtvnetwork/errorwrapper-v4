package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/errorwrapper"
)

// LinkedCollections TODO constructors
type LinkedCollections struct {
	*corestr.LinkedCollections
	ErrorWrapper *errorwrapper.Wrapper
}
