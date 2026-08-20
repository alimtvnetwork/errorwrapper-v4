package errverifytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ConsumeChannel(t *testing.T) {
	Convey("ConsumeChannel feeds lines and finishes on close", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{"a", "b", "c"}),
		}

		ch := make(chan string, 3)
		ch <- "a"
		ch <- "b"
		ch <- "c"
		close(ch)

		err := errverify.ConsumeChannel(v, ch)
		So(err, ShouldBeNil)
	})

	Convey("ConsumeChannel surfaces mismatches", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{"a", "b"}),
		}

		ch := make(chan string, 2)
		ch <- "a"
		ch <- "X"
		close(ch)

		err := errverify.ConsumeChannel(v, ch)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "mismatch at index 1")
	})

	Convey("Nil verifier is a no-op", t, func() {
		ch := make(chan string)
		close(ch)
		So(errverify.ConsumeChannel(nil, ch), ShouldBeNil)
	})

	Convey("Nil channel still finishes verifier", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{"missing"}),
		}
		err := errverify.ConsumeChannel(v, nil)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "missing expected line")
	})
}

func Test_ConsumeCollection_NilSafe(t *testing.T) {
	Convey("Nil collection still triggers Finish", t, func() {
		v := &errverify.StreamingCollectionVerifier{
			Mode:         errverify.StreamMatchEqual,
			ExpectedLine: sliceSource([]string{}),
		}
		So(errverify.ConsumeCollection(v, nil, false), ShouldBeNil)
	})

	Convey("Nil verifier is a no-op", t, func() {
		So(errverify.ConsumeCollection(nil, nil, false), ShouldBeNil)
	})
}
