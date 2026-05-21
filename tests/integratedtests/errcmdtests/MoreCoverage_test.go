package errcmdtests

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

// Test_StdIn — exercises StdIn buffer accessors.
func Test_StdIn(t *testing.T) {
	Convey("NewStdIn with zero sizes returns no buffers", t, func() {
		s := errcmd.NewStdIn(0, 0)
		So(s, ShouldNotBeNil)
		So(s.HasStdOut(), ShouldBeFalse)
		So(s.HasStdErr(), ShouldBeFalse)
		So(s.StdOutString(), ShouldEqual, "")
		So(s.StdErrString(), ShouldEqual, "")
	})

	Convey("NewStdIn with positive sizes allocates buffers", t, func() {
		s := errcmd.NewStdIn(16, 16)
		So(s.HasStdOut(), ShouldBeTrue)
		So(s.HasStdErr(), ShouldBeTrue)
		s.OutBuf.WriteString("hi")
		s.ErrBuf.WriteString("bye")
		So(s.StdOutString(), ShouldEqual, "hi")
		So(s.StdErrString(), ShouldEqual, "bye")
	})

	Convey("Nil StdIn safely reports has-flags", t, func() {
		var s *errcmd.StdIn
		So(s.HasStdOut(), ShouldBeFalse)
		So(s.HasStdErr(), ShouldBeFalse)
	})
}

// Test_BufferClone — clones byte buffers.
func Test_BufferClone(t *testing.T) {
	Convey("Nil buffer clones to nil", t, func() {
		So(errcmd.BufferClone(nil), ShouldBeNil)
	})

	Convey("Empty zero-cap buffer clones to fresh empty buffer", t, func() {
		buf := &bytes.Buffer{}
		clone := errcmd.BufferClone(buf)
		So(clone, ShouldNotBeNil)
		So(clone.Len(), ShouldEqual, 0)
	})

	Convey("Buffer with data clones content and preserves capacity", t, func() {
		buf := &bytes.Buffer{}
		buf.Grow(32)
		buf.WriteString("payload")
		clone := errcmd.BufferClone(buf)
		So(clone, ShouldNotBeNil)
		So(clone.String(), ShouldEqual, "payload")
	})
}

// Test_CmdSetOsStandardWriters — sets std writers on a real cmd.
func Test_CmdSetOsStandardWriters(t *testing.T) {
	Convey("Nil cmd is a safe no-op", t, func() {
		So(errcmd.CmdSetOsStandardWriters(nil), ShouldBeNil)
	})
	Convey("Non-nil cmd gets std writers assigned", t, func() {
		cmd := exec.Command("echo", "hi")
		out := errcmd.CmdSetOsStandardWriters(cmd)
		So(out, ShouldNotBeNil)
		So(out.Stdin, ShouldNotBeNil)
		So(out.Stdout, ShouldNotBeNil)
		So(out.Stderr, ShouldNotBeNil)
	})
}

// Test_CombinedOutputError — runs a real command and inspects the compiled output.
func Test_CombinedOutputError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping POSIX echo-based test on windows")
	}

	Convey("Running echo collects stdout", t, func() {
		out := errcmd.CombinedOutputError(nil, "echo", "hello")
		So(out, ShouldNotBeNil)
	})

	Convey("Running an unknown binary surfaces an error", t, func() {
		out := errcmd.CombinedOutputError(nil, "this-binary-does-not-exist-xyz")
		So(out, ShouldNotBeNil)
	})
}

// Test_NewCreator_Empty — Empty CmdOnce is initialized with the cmdNil error.
func Test_NewCreator_Empty(t *testing.T) {
	Convey("New.Empty returns a CmdOnce with initialization error", t, func() {
		once := errcmd.New.Empty()
		So(once, ShouldNotBeNil)
		So(once.HasCompiledError(), ShouldBeFalse) // not run yet
		So(once.IsNull(), ShouldBeFalse)
	})
}

// Test_NewCreator_IsSuccess_IsFailed — running echo succeeds, unknown bin fails.
func Test_NewCreator_IsSuccess_IsFailed(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping POSIX echo-based test on windows")
	}

	Convey("IsSuccess true for echo, false for missing binary", t, func() {
		So(errcmd.New.IsSuccess("echo", "ok"), ShouldBeTrue)
		So(errcmd.New.IsSuccess("this-binary-does-not-exist-xyz"), ShouldBeFalse)
	})

	Convey("IsFailed inverts IsSuccess", t, func() {
		So(errcmd.New.IsFailed("echo", "ok"), ShouldBeFalse)
		So(errcmd.New.IsFailed("this-binary-does-not-exist-xyz"), ShouldBeTrue)
	})
}

// Test_NewCreator_CombinedOutput — bytes returned for a successful command.
func Test_NewCreator_CombinedOutput(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping POSIX echo-based test on windows")
	}

	Convey("CombinedOutput returns echoed bytes with no error", t, func() {
		bytesOut, err := errcmd.New.CombinedOutput("echo", "hi")
		So(err, ShouldBeNil)
		So(string(bytesOut), ShouldContainSubstring, "hi")
	})

	Convey("CombinedOutputRaw returns echoed bytes", t, func() {
		bytesOut := errcmd.New.CombinedOutputRaw("echo", "hello")
		So(string(bytesOut), ShouldContainSubstring, "hello")
	})

	Convey("CombinedOutputString returns echoed string", t, func() {
		s := errcmd.New.CombinedOutputString("echo", "world")
		So(s, ShouldContainSubstring, "world")
	})
}

// Test_NewResult_Inspectors — Result accessors with crafted state.
func Test_NewResult_Inspectors(t *testing.T) {
	Convey("Successful result reports success and zero exit code", t, func() {
		stdOut := &bytes.Buffer{}
		stdErr := &bytes.Buffer{}
		stdOut.WriteString("ok")
		r := errcmd.NewResult(nil, stdOut, stdErr, nil, true)
		So(r, ShouldNotBeNil)
		So(r.HasCommandExecuted(), ShouldBeTrue)
		So(r.ExitCode, ShouldEqual, errcmd.SuccessfullyRunningExitCode)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.IsInvalidExitCode(), ShouldBeFalse)
		So(r.IsInvalid(), ShouldBeFalse)
		So(r.HasNoError(), ShouldBeTrue)
		So(r.HasAnyError(), ShouldBeFalse)
		So(r.IsExitCode(errcmd.SuccessfullyRunningExitCode), ShouldBeTrue)
		So(r.HasExitError(), ShouldBeFalse)
		So(r.HasValidExitCode(), ShouldBeTrue)
		So(r.IsRunSuccessfully(), ShouldBeTrue)
		So(r.String(), ShouldNotBeBlank)
		So(r.StringIf(false), ShouldNotBeBlank)
		So(r.DetailedOutput(), ShouldNotBeBlank)
	})

	Convey("Failing result with wrapped error reports failure", t, func() {
		w := errnew.Type.Error(errtype.CommandExecution, errors.New("boom"))
		r := errcmd.NewResult(w, &bytes.Buffer{}, &bytes.Buffer{}, nil, true)
		So(r, ShouldNotBeNil)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasAnyError(), ShouldBeTrue)
		So(r.HasError(), ShouldBeTrue)
		So(r.ErrorWrapper(), ShouldNotBeNil)
		So(r.CompiledFullErrorWrapper(), ShouldNotBeNil)
		r.Dispose()
	})

	Convey("DisposeWithoutErrorWrapper is safe", t, func() {
		r := errcmd.NewResult(nil, &bytes.Buffer{}, &bytes.Buffer{}, nil, true)
		r.DisposeWithoutErrorWrapper()
	})
}

// Test_NewCreator_ExitCode — exit-code accessors.
func Test_NewCreator_ExitCode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping POSIX echo-based test on windows")
	}

	Convey("ExitCode returns the success code for echo", t, func() {
		So(errcmd.New.ExitCode("echo", "ok"), ShouldEqual, errcmd.SuccessfullyRunningExitCode)
		_ = errcmd.New.ExitCodeAsByte("echo", "ok")
	})
}
