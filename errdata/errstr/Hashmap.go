package errstr

import (
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

type Hashmap struct {
	*corestr.Hashmap
	ErrorWrapper *errorwrapper.Wrapper
}
