package errcmdportabletests

import (
	"errors"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errcmdportable"
	"github.com/alimtvnetwork/errorwrapper-v3/errcmdportable/osadapter"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_NoProcessRunner(t *testing.T) {
	Convey("NoProcessRunner advertises CapabilityNoProcess", t, func() {
		r := errcmdportable.NoProcessRunner{}
		So(r.Capability(), ShouldEqual, errcmdportable.CapabilityNoProcess)
	})

	Convey("NoProcessRunner.Run returns ErrNotSupported", t, func() {
		r := errcmdportable.NoProcessRunner{}
		res := r.Run("echo", "hi")
		So(res.HasError(), ShouldBeTrue)
		So(res.IsNotSupported(), ShouldBeTrue)
		So(errors.Is(res.Err, errcmdportable.ErrNotSupported), ShouldBeTrue)
		So(res.ExitCode, ShouldEqual, -1)
	})
}

func Test_Detect(t *testing.T) {
	Convey("Detect returns NoProcessRunner by default", t, func() {
		r := errcmdportable.Detect()
		So(r.Capability(), ShouldEqual, errcmdportable.CapabilityNoProcess)
	})
}

func Test_OsAdapter(t *testing.T) {
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Skip("os/exec unsupported on edge target")
	}

	Convey("osadapter.New advertises CapabilityOsExec", t, func() {
		r := osadapter.New()
		So(r.Capability(), ShouldEqual, errcmdportable.CapabilityOsExec)
	})

	Convey("osadapter runs a real process", t, func() {
		r := osadapter.New()
		var res errcmdportable.Result
		switch runtime.GOOS {
		case "windows":
			res = r.Run("cmd", "/C", "echo hi")
		default:
			res = r.Run("echo", "hi")
		}
		So(res.HasError(), ShouldBeFalse)
		So(res.ExitCode, ShouldEqual, 0)
		So(res.Stdout, ShouldContainSubstring, "hi")
	})

	Convey("osadapter surfaces non-zero exit code", t, func() {
		r := osadapter.New()
		var res errcmdportable.Result
		switch runtime.GOOS {
		case "windows":
			res = r.Run("cmd", "/C", "exit 7")
		default:
			res = r.Run("sh", "-c", "exit 7")
		}
		So(res.ExitCode, ShouldEqual, 7)
		So(res.HasError(), ShouldBeTrue)
		So(res.IsNotSupported(), ShouldBeFalse)
	})
}
