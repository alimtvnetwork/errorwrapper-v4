package errverifytestwrappers

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errverify"
)

type VerifyTestWrapper struct {
	ErrorWrapper *errorwrapper.Wrapper
	Verifier     errverify.Verifier
}
