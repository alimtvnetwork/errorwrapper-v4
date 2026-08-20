package errverifytestwrappers

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
)

type VerifyTestWrapper struct {
	ErrorWrapper *errorwrapper.Wrapper
	Verifier     errverify.Verifier
}
