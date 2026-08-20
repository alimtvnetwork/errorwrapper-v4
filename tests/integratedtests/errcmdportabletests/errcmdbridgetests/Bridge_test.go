package errcmdbridgetests

import (
	"bytes"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errcmd"
	"github.com/alimtvnetwork/errorwrapper-v4/errcmdportable/errcmdbridge"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_FromErrcmdResult_Nil(t *testing.T) {
	Convey("FromErrcmdResult returns zero-value Result for nil input", t, func() {
		got := errcmdbridge.FromErrcmdResult(nil)
		So(got.ExitCode, ShouldEqual, 0)
		So(got.Stdout, ShouldEqual, "")
		So(got.Stderr, ShouldEqual, "")
		So(got.Err, ShouldBeNil)
		So(got.HasError(), ShouldBeFalse)
		So(got.IsNotSupported(), ShouldBeFalse)
	})
}

func Test_FromErrcmdResult_SuccessWithStdout(t *testing.T) {
	Convey("FromErrcmdResult mirrors stdout for a successful errcmd.Result", t, func() {
		stdOut := bytes.NewBufferString("hello world\n")
		stdErr := bytes.NewBufferString("")
		src := errcmd.NewResult(nil, stdOut, stdErr, nil, true)

		got := errcmdbridge.FromErrcmdResult(src)

		So(got.HasError(), ShouldBeFalse)
		So(got.Err, ShouldBeNil)
		So(got.Stdout, ShouldEqual, "hello world")
		So(got.Stderr, ShouldEqual, "")
	})
}

func Test_FromErrcmdResult_CarriesErrorWrapper(t *testing.T) {
	Convey("FromErrcmdResult carries the underlying errorWrapper as Err", t, func() {
		wrap := errnew.Messages.Single(errtype.FailedProcess, "boom")
		stdOut := bytes.NewBufferString("")
		stdErr := bytes.NewBufferString("boom\n")
		src := errcmd.NewResult(wrap, stdOut, stdErr, nil, true)

		got := errcmdbridge.FromErrcmdResult(src)

		So(got.HasError(), ShouldBeTrue)
		So(got.Err, ShouldNotBeNil)
		So(got.Stderr, ShouldEqual, "boom")
		So(got.IsNotSupported(), ShouldBeFalse)
	})
}
