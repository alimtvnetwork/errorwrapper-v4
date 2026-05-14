package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/errorwrapper"
)

type Hashset struct {
	*corestr.Hashset
	ErrorWrapper *errorwrapper.Wrapper
}
