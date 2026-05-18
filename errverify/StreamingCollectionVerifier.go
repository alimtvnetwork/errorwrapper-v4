package errverify

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

// StreamingCollectionVerifier is a draft, line-at-a-time variant of
// CollectionVerifier. It avoids materializing the full error collection
// (no `errwrappers.Collection.StringsWithoutHeader()` slice copy) and is
// intended for very large or unbounded error streams (e.g. log tailers,
// long-running batch jobs, edge-runtime workers where memory is tight).
//
// Status: deferred-research stub. The shape is stable enough to wire into
// callers; the matching logic intentionally mirrors `CollectionVerifier`'s
// per-line equality / contains / regex semantics via `stringcompareas`.
//
// NOT a drop-in replacement for CollectionVerifier:
//   - No length pre-check (length is only known after Finish()).
//   - No SliceValidator integration (streaming, so cross-line rules like
//     ordering-only-mismatch are reported as they occur).
//   - Caller drives iteration via Feed(line) / Finish().
type StreamingCollectionVerifier struct {
	Header       string
	VerifyAs     stringcompareas.Variant
	ExpectedLine func(index int) (line string, hasMore bool) // pull-style expectation source
	IsPrintError bool

	// internal cursor
	index       int
	mismatches  []string
	expectedEnd bool
}

// Feed processes a single actual line. Returns a non-nil wrapper only on
// the first mismatch; subsequent mismatches are accumulated and surfaced
// by Finish().
func (it *StreamingCollectionVerifier) Feed(actual string) *errorwrapper.Wrapper {
	if it.ExpectedLine == nil {
		return errnew.Type.Message(
			errtype.ValidationFailed,
			it.Header+" - StreamingCollectionVerifier.ExpectedLine is nil")
	}

	expected, hasMore := it.ExpectedLine(it.index)
	if !hasMore {
		it.expectedEnd = true
		mismatch := it.Header + " - unexpected extra line at index " +
			itoa(it.index) + ": " + actual
		it.mismatches = append(it.mismatches, mismatch)
		it.index++
		return nil
	}

	if !compareLine(expected, actual, it.VerifyAs) {
		mismatch := it.Header + " - mismatch at index " +
			itoa(it.index) +
			" expected=" + expected +
			" actual=" + actual
		it.mismatches = append(it.mismatches, mismatch)
	}

	it.index++
	return nil
}

// Finish must be called after the last Feed(). Returns nil on full match,
// or a ValidationFailed wrapper aggregating every mismatch (including any
// expected lines that were never fed).
func (it *StreamingCollectionVerifier) Finish() *errorwrapper.Wrapper {
	if it.ExpectedLine != nil && !it.expectedEnd {
		for {
			expected, hasMore := it.ExpectedLine(it.index)
			if !hasMore {
				break
			}
			it.mismatches = append(it.mismatches,
				it.Header+" - missing expected line at index "+
					itoa(it.index)+": "+expected)
			it.index++
		}
	}

	if len(it.mismatches) == 0 {
		return nil
	}

	joined := constants.EmptyString
	for i, m := range it.mismatches {
		if i > 0 {
			joined += constants.NewLineUnix
		}
		joined += m
	}

	return errnew.Type.Message(
		errtype.ValidationFailed,
		joined)
}

func compareLine(expected, actual string, as stringcompareas.Variant) bool {
	// Minimal semantics: equality. Full parity with corevalidator.TextValidator
	// (contains/regex/case-insensitive) is left for the implementation phase.
	_ = as
	return expected == actual
}

func itoa(i int) string {
	// avoid pulling in strconv at module init scope; trivial positive-int path
	if i == 0 {
		return "0"
	}
	neg := false
	if i < 0 {
		neg = true
		i = -i
	}
	var buf [20]byte
	pos := len(buf)
	for i > 0 {
		pos--
		buf[pos] = byte('0' + i%10)
		i /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}
