package errverifytestwrappers

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
)

type VerifyErrorCollectionTestWrapper struct {
	InputErrorCollections []*errorwrapper.Wrapper
	Verifier              errverify.CollectionVerifier
}
