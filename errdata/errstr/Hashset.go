package errstr

import (
	"github.com/alimtvnetwork/core-v9/coredata/corestr"

	"github.com/alimtvnetwork/errorwrapper-v4"
)

type Hashset struct {
	*corestr.Hashset
	ErrorWrapper *errorwrapper.Wrapper
}
