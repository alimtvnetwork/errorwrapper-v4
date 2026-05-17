package errcmdtests

import (
	"os/exec"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
	"github.com/alimtvnetwork/errorwrapper-v3/errorwrapper"
	. "github.com/smartystreets/goconvey/convey"
)

type emptyChecker bool

func (e emptyChecker) IsEmpty() bool { return bool(e) }

func Test_ArgsJoinFamily(t *testing.T) {
	Convey("ArgsJoin / ArgsJoinSlice / ArgsJoinSlicePtr", t, func() {
		Convey("Empty inputs return empty string", func() {
			So(errcmd.ArgsJoin(), ShouldEqual, "")
			So(errcmd.ArgsJoinSlice(nil), ShouldEqual, "")
			So(errcmd.ArgsJoinSlice([]string{}), ShouldEqual, "")
			So(errcmd.ArgsJoinSlicePtr(nil), ShouldEqual, "")
			empty := []string{}
			So(errcmd.ArgsJoinSlicePtr(&empty), ShouldEqual, "")
		})

		Convey("Joins with spaces", func() {
			So(errcmd.ArgsJoin("ls", "-la"), ShouldEqual, "ls -la")
			So(errcmd.ArgsJoinSlice([]string{"a", "b", "c"}), ShouldEqual, "a b c")
			slice := []string{"x", "y"}
			So(errcmd.ArgsJoinSlicePtr(&slice), ShouldEqual, "x y")
		})
	})

	Convey("ArgsJoinPrepend prepends the first arg", t, func() {
		So(errcmd.ArgsJoinPrepend("sudo"), ShouldEqual, "sudo")
		So(errcmd.ArgsJoinPrepend("sudo", "ls", "-la"), ShouldEqual, "sudo ls -la")
	})

	Convey("ArgsJoinWithSingle appends the leading arg after the rest", t, func() {
		So(errcmd.ArgsJoinWithSingle("solo"), ShouldEqual, "")
		So(errcmd.ArgsJoinWithSingle("end", "a", "b"), ShouldEqual, "a b end")
	})

	Convey("ProcessArgsJoinAppend places process after args", t, func() {
		So(errcmd.ProcessArgsJoinAppend("proc"), ShouldEqual, "proc")
		So(errcmd.ProcessArgsJoinAppend("proc", "--flag", "v"), ShouldEqual, "--flag v proc")
	})
}

func Test_GetFormattedKeyValueData(t *testing.T) {
	Convey("Produces KEY=VALUE", t, func() {
		So(errcmd.GetFormattedKeyValueData("MY_VAR", "value"), ShouldEqual, "MY_VAR=value")
		So(errcmd.GetFormattedKeyValueData("", ""), ShouldEqual, "=")
	})
}

func Test_GetExitCode(t *testing.T) {
	Convey("Returns success when no error and no exit error", t, func() {
		So(errcmd.GetExitCode(emptyChecker(true), nil), ShouldEqual, errcmd.SuccessfullyRunningExitCode)
	})

	Convey("Returns invalid exit code when wrapper not empty and no exit error", t, func() {
		So(errcmd.GetExitCode(emptyChecker(false), nil), ShouldEqual, errcmd.InvalidExitCode)
	})
}

func Test_Conditional(t *testing.T) {
	Convey("Picks true/false branches", t, func() {
		trueOne := &errcmd.CmdOnce{}
		falseOne := &errcmd.CmdOnce{}
		So(errcmd.Conditional(true, trueOne, falseOne), ShouldEqual, trueOne)
		So(errcmd.Conditional(false, trueOne, falseOne), ShouldEqual, falseOne)
	})
}

func Test_CmdToScriptLine(t *testing.T) {
	Convey("Nil / empty path returns empty string", t, func() {
		So(errcmd.CmdToScriptLine(nil), ShouldEqual, "")
		So(errcmd.CmdToScriptLine(&exec.Cmd{}), ShouldEqual, "")
	})

	Convey("Joins path with args", t, func() {
		cmd := &exec.Cmd{Path: "/bin/ls", Args: []string{"-la"}}
		So(errcmd.CmdToScriptLine(cmd), ShouldEqual, "-la /bin/ls")
	})
}

func Test_CurrentOsDetails(t *testing.T) {
	Convey("Returns details or wrapped error on current OS", t, func() {
		details, w := errcmd.CurrentOsDetails()
		if w != nil {
			So(w, ShouldNotBeNil)
			So(w.HasError(), ShouldBeTrue)
			return
		}
		So(details, ShouldNotBeNil)

		var _ *errorwrapper.Wrapper = w
	})
}
