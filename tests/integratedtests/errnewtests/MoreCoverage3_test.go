package errnewtests

import (
	"errors"
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func Test_MoreCoverage3_NotFound_Direct(t *testing.T) {
	Convey("NotFound zero-arg type-default constructors", t, func() {
		So(errnew.NotFound.Type().HasError(), ShouldBeTrue)
		So(errnew.NotFound.Record().HasError(), ShouldBeTrue)
		So(errnew.NotFound.Process().HasError(), ShouldBeTrue)
		So(errnew.NotFound.User().HasError(), ShouldBeTrue)
		So(errnew.NotFound.RedisCmd().HasError(), ShouldBeTrue)
		So(errnew.NotFound.Id().HasError(), ShouldBeTrue)
		So(errnew.NotFound.Name().HasError(), ShouldBeTrue)
		So(errnew.NotFound.Data().HasError(), ShouldBeTrue)
		So(errnew.NotFound.Group().HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage3_NotFound_ErrorPaths(t *testing.T) {
	Convey("NotFound error/process/cmd variants", t, func() {
		So(errnew.NotFound.Error(nil), ShouldBeNil)
		So(errnew.NotFound.Error(errors.New("e")).HasError(), ShouldBeTrue)

		So(errnew.NotFound.Simple("missing").HasError(), ShouldBeTrue)

		So(errnew.NotFound.ErrorReference(nil, "r"), ShouldBeNil)
		So(errnew.NotFound.ErrorReference(errors.New("e"), "r").HasError(), ShouldBeTrue)

		So(errnew.NotFound.ProcessError(nil), ShouldBeNil)
		So(errnew.NotFound.ProcessError(errors.New("e")).HasError(), ShouldBeTrue)

		So(errnew.NotFound.ProcessWholeCommand(""), ShouldBeNil)
		So(errnew.NotFound.ProcessWholeCommand("ls -la").HasError(), ShouldBeTrue)

		// Cmd / CmdWholeCommand return nil when cmd != nil, wrapper when nil.
		cmd := exec.Command("echo", "hi")
		So(errnew.NotFound.Cmd(cmd), ShouldBeNil)
		So(errnew.NotFound.Cmd(nil).HasError(), ShouldBeTrue)
		So(errnew.NotFound.CmdWholeCommand(cmd, "echo hi"), ShouldBeNil)
		So(errnew.NotFound.CmdWholeCommand(nil, "echo hi").HasError(), ShouldBeTrue)

		So(errnew.NotFound.ProcessMessageError("msg", nil), ShouldBeNil)
		So(errnew.NotFound.ProcessMessageError("msg", errors.New("e")).HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage3_Fmt_All(t *testing.T) {
	Convey("Fmt creator family", t, func() {
		So(errnew.Fmt.If(false, errtype.IO, "x %d", 1), ShouldBeNil)
		So(errnew.Fmt.If(true, errtype.IO, "x %d", 1).HasError(), ShouldBeTrue)

		So(errnew.Fmt.Fmt(errtype.IO, "x %s", "y").HasError(), ShouldBeTrue)
		So(errnew.Fmt.Error(errtype.IO, errors.New("e"), "x %s", "y").HasError(), ShouldBeTrue)
		So(errnew.Fmt.Message(errtype.IO, "m", "x %s", "y").HasError(), ShouldBeTrue)
		So(errnew.Fmt.MessageError(errtype.IO, "m", errors.New("e"), "x %s", "y").HasError(), ShouldBeTrue)

		m := map[string]interface{}{"name": "alice", "id": 7}

		So(errnew.Fmt.UsingMapOptions(false, errtype.IO, "hi %name%", m).HasError(), ShouldBeTrue)
		So(errnew.Fmt.UsingMapOptions(true, errtype.IO, "hi {name}", m).HasError(), ShouldBeTrue)

		So(errnew.Fmt.ErrorFormatUsingMap(errtype.IO, errors.New("e"), "hi %name%", m).HasError(), ShouldBeTrue)

		So(errnew.Fmt.MsgFormatUsingMapIf(false, errtype.IO, "msg", "hi %name%", m), ShouldBeNil)
		So(errnew.Fmt.MsgFormatUsingMapIf(true, errtype.IO, "msg", "hi %name%", m).HasError(), ShouldBeTrue)

		So(errnew.Fmt.MsgFormatUsingMap(errtype.IO, "msg", "hi %name%", m).HasError(), ShouldBeTrue)
		So(errnew.Fmt.FormatUsingMap(errtype.IO, "hi %name%", m).HasError(), ShouldBeTrue)
		So(errnew.Fmt.CurlyFormatUsingMap(errtype.IO, "hi {name}", m).HasError(), ShouldBeTrue)
		So(errnew.Fmt.CurlyMsgFormatUsingMap(errtype.IO, "msg", "hi {name}", m).HasError(), ShouldBeTrue)
	})
}
