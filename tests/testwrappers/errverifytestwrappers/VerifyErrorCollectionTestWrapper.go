package errverifytestwrappers

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errverify"
)

type VerifyErrorCollectionTestWrapper struct {
	InputErrorCollections []*errorwrapper.Wrapper
	Verifier              errverify.CollectionVerifier
}
