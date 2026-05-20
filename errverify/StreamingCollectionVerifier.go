package errverify

import (
	"errors"
	"regexp"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
)

// StreamMatchMode controls how StreamingCollectionVerifier matches an
// actual line against an expected line. Independent of the upstream
// `stringcompareas.Variant` enum so this stub stays self-contained;
// callers can map their preferred Variant via FromVariant().
type StreamMatchMode int

const (
	// StreamMatchEqual — exact byte equality (default).
	StreamMatchEqual StreamMatchMode = iota
	// StreamMatchEqualFold — case-insensitive equality.
	StreamMatchEqualFold
	// StreamMatchContains — actual must contain expected as substring.
	StreamMatchContains
	// StreamMatchContainsFold — case-insensitive contains.
	StreamMatchContainsFold
	// StreamMatchRegex — expected is a Go regexp.MustCompile pattern,
	// matched against the actual line.
	StreamMatchRegex
)

// FromVariant maps the upstream stringcompareas.Variant to the local
// StreamMatchMode. Unknown variants fall back to StreamMatchEqual.
func FromVariant(v stringcompareas.Variant) StreamMatchMode {
	// Defensive: stringcompareas constants aren't enumerated here to
	// avoid coupling. Match by name to stay forward-compatible.
	switch v.Name() {
	case "Equal", "":
		return StreamMatchEqual
	case "EqualFold", "EqualCaseInsensitive":
		return StreamMatchEqualFold
	case "Contains":
		return StreamMatchContains
	case "ContainsFold", "ContainsCaseInsensitive":
		return StreamMatchContainsFold
	case "Regex", "Regexp":
		return StreamMatchRegex
	}
	return StreamMatchEqual
}

// StreamingCollectionVerifier is a draft, line-at-a-time variant of
// CollectionVerifier. It avoids materializing the full error collection
// (no `errwrappers.Collection.StringsWithoutHeader()` slice copy) and is
// intended for very large or unbounded error streams (e.g. log tailers,
// long-running batch jobs, edge-runtime workers where memory is tight).
//
// NOT a drop-in replacement for CollectionVerifier:
//   - Length is only known after Finish(); use ExpectedLength (>0) for a
//     soft check that runs in Finish().
//   - SliceValidator ordering rules are not applied; this is positional
//     line-by-line matching.
//   - Caller drives iteration via Feed(line) / Finish().
type StreamingCollectionVerifier struct {
	Header         string
	Mode           StreamMatchMode
	ExpectedLine   func(index int) (line string, hasMore bool)
	ExpectedLength int // 0 = no check
	IsPrintError   bool

	// internal cursor
	index       int
	mismatches  []string
	expectedEnd bool
	regexCache  map[string]*regexp.Regexp
}

// Feed processes a single actual line. Mismatches are accumulated and
// surfaced by Finish(); Feed itself returns non-nil only for setup
// errors (nil ExpectedLine, bad regex).
func (it *StreamingCollectionVerifier) Feed(actual string) error {
	if it.ExpectedLine == nil {
		return errors.New(it.Header + " - StreamingCollectionVerifier.ExpectedLine is nil")
	}

	expected, hasMore := it.ExpectedLine(it.index)
	if !hasMore {
		it.expectedEnd = true
		it.mismatches = append(it.mismatches,
			it.Header+" - unexpected extra line at index "+
				itoa(it.index)+": "+actual)
		it.index++
		return nil
	}

	ok, setupErr := it.compare(expected, actual)
	if setupErr != nil {
		return setupErr
	}
	if !ok {
		it.mismatches = append(it.mismatches,
			it.Header+" - mismatch at index "+itoa(it.index)+
				" expected="+expected+" actual="+actual)
	}

	it.index++
	return nil
}

// Finish must be called after the last Feed(). Returns nil on full match,
// or a ValidationFailed wrapper aggregating every mismatch.
func (it *StreamingCollectionVerifier) Finish() error {
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

	if it.ExpectedLength > 0 && it.ExpectedLength != it.fedCount() {
		it.mismatches = append(it.mismatches,
			it.Header+" - length mismatch expected="+
				itoa(it.ExpectedLength)+" actual="+itoa(it.fedCount()))
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

	return errors.New(joined)
}

// fedCount approximates how many actual lines were fed (index advances
// once per Feed, including over-feeds past expected end).
func (it *StreamingCollectionVerifier) fedCount() int { return it.index }

func (it *StreamingCollectionVerifier) compare(expected, actual string) (bool, error) {
	switch it.Mode {
	case StreamMatchEqual:
		return expected == actual, nil
	case StreamMatchEqualFold:
		return strings.EqualFold(expected, actual), nil
	case StreamMatchContains:
		return strings.Contains(actual, expected), nil
	case StreamMatchContainsFold:
		return strings.Contains(strings.ToLower(actual), strings.ToLower(expected)), nil
	case StreamMatchRegex:
		if it.regexCache == nil {
			it.regexCache = map[string]*regexp.Regexp{}
		}
		re, ok := it.regexCache[expected]
		if !ok {
			compiled, err := regexp.Compile(expected)
			if err != nil {
				return false, errors.New(it.Header + " - invalid regex at index " + itoa(it.index) + ": " + err.Error())
			}
			re = compiled
			it.regexCache[expected] = re
		}
		return re.MatchString(actual), nil
	}
	return expected == actual, nil
}

func itoa(i int) string {
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
