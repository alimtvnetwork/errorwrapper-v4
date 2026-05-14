package errverifytestwrappers

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errverify"
)

type VerifyTestWrapper struct {
	ErrorWrapper *errorwrapper.Wrapper
	Verifier     errverify.Verifier
}
