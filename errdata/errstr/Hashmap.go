package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

type Hashmap struct {
	*corestr.Hashmap
	ErrorWrapper *errorwrapper.Wrapper
}
