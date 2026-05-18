package errverify

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

// ConsumeCollection feeds every wrapper line of `coll` into `verifier`
// without materializing an intermediate `[]string` slice. Caller chooses
// whether reference lines are included.
//
// Returns the aggregated mismatch wrapper from verifier.Finish() (nil on
// full match), or the first setup error from verifier.Feed().
//
// This is the streaming-consumer adapter referenced in
// docs/extensibility.md §6.2.
func ConsumeCollection(
	verifier *StreamingCollectionVerifier,
	coll *errwrappers.Collection,
	isIncludeReferences bool,
) *errorwrapper.Wrapper {
	if verifier == nil {
		return nil
	}
	if coll == nil || coll.IsEmpty() {
		return verifier.Finish()
	}

	for _, w := range coll.Items() {
		if w == nil {
			continue
		}
		line := wrapperLine(w, isIncludeReferences)
		if setupErr := verifier.Feed(line); setupErr != nil {
			return setupErr
		}
	}

	return verifier.Finish()
}

// ConsumeChannel feeds every string from `lines` into `verifier`. Use
// for true streaming sources (log tailers, scanners, network streams)
// where the collection never exists in memory.
//
// The channel must be closed by the producer when finished; this
// function blocks until it drains.
func ConsumeChannel(
	verifier *StreamingCollectionVerifier,
	lines <-chan string,
) *errorwrapper.Wrapper {
	if verifier == nil {
		return nil
	}
	if lines == nil {
		return verifier.Finish()
	}

	for line := range lines {
		if setupErr := verifier.Feed(line); setupErr != nil {
			return setupErr
		}
	}

	return verifier.Finish()
}

func wrapperLine(w *errorwrapper.Wrapper, isIncludeReferences bool) string {
	if isIncludeReferences {
		return w.FullString()
	}
	return w.FullStringWithoutReferences()
}
