package errverifytestwrappers

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errverify"
)

type VerifyErrorCollectionTestWrapper struct {
	InputErrorCollections []*errorwrapper.Wrapper
	Verifier              errverify.CollectionVerifier
}
