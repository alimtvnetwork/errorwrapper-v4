package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/errorwrapper"
)

type HashsetsCollection struct {
	*corestr.HashsetsCollection
	ErrorWrapper *errorwrapper.Wrapper
}
