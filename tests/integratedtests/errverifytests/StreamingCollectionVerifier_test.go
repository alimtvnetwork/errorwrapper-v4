package errverifytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
	. "github.com/smartystreets/goconvey/convey"
)

func sliceSource(lines []string) func(int) (string, bool) {
	return func(i int) (string, bool) {
		if i < 0 || i >= len(lines) {
			return "", false
		}
		return lines[i], true
	}
}

func Test_StreamingCollectionVerifier_Equal(t *testing.T) {
	Convey("Equal mode — all lines match", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Header:       "case",
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{"a", "b", "c"}),
		}
		for _, l := range []string{"a", "b", "c"} {
			So(v.Feed(l), ShouldBeNil)
		}
		So(v.Finish(), ShouldBeNil)
	})

	Convey("Equal mode — mismatch surfaces in Finish", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{"a", "b"}),
		}
		v.Feed("a")
		v.Feed("X")
		err := v.Finish()
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "mismatch at index 1")
	})

	Convey("Extra actual lines flagged", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{"a"}),
		}
		v.Feed("a")
		v.Feed("b")
		err := v.Finish()
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unexpected extra line")
	})

	Convey("Missing expected lines flagged", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{"a", "b", "c"}),
		}
		v.Feed("a")
		err := v.Finish()
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "missing expected line")
	})
}

func Test_StreamingCollectionVerifier_Modes(t *testing.T) {
	Convey("EqualFold matches case-insensitive", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqualFold,
			ExpectedLine: sliceSource([]string{"Hello"}),
		}
		So(v.Feed("HELLO"), ShouldBeNil)
		So(v.Finish(), ShouldBeNil)
	})

	Convey("Contains matches substring", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchContains,
			ExpectedLine: sliceSource([]string{"err"}),
		}
		So(v.Feed("internal error occurred"), ShouldBeNil)
		So(v.Finish(), ShouldBeNil)
	})

	Convey("ContainsFold matches case-insensitive substring", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchContainsFold,
			ExpectedLine: sliceSource([]string{"ERR"}),
		}
		So(v.Feed("internal error"), ShouldBeNil)
		So(v.Finish(), ShouldBeNil)
	})

	Convey("Regex matches pattern", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchRegex,
			ExpectedLine: sliceSource([]string{`^code=\d+$`}),
		}
		So(v.Feed("code=42"), ShouldBeNil)
		So(v.Finish(), ShouldBeNil)
	})

	Convey("Invalid regex returns setup error from Feed", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchRegex,
			ExpectedLine: sliceSource([]string{`(`}),
		}
		err := v.Feed("anything")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "invalid regex")
	})
}

func Test_StreamingCollectionVerifier_LengthCheck(t *testing.T) {
	Convey("ExpectedLength enforced in Finish", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:           errverify.StreamMatchEqual,
			ExpectedLine:   sliceSource([]string{"a", "b"}),
			ExpectedLength: 3,
		}
		v.Feed("a")
		v.Feed("b")
		err := v.Finish()
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "length mismatch")
	})
}

func Test_StreamingCollectionVerifier_NilSource(t *testing.T) {
	Convey("Nil ExpectedLine returns setup error", t, func() {
		v := &errverify.StreamingCollectionVerifier{}
		err := v.Feed("anything")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "ExpectedLine is nil")
	})
}
