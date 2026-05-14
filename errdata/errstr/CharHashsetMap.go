package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/errorwrapper"
)

type CharHashsetMap struct {
	*corestr.CharHashsetMap
	ErrorWrapper *errorwrapper.Wrapper
}
