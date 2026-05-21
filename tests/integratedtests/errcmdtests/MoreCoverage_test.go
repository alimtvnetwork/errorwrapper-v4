package errcmdtests

import (
	"bytes"
	"testing"

	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
	. "github.com/smartystreets/goconvey/convey"
)

// MoreCoverage_test.go — pure construction / getter / disposal paths
// for the errcmd package. Avoids spawning real subprocesses so the suite
// stays portable across Linux CI and Windows developer machines.

func Test_MoreCoverage_BufferClone(t *testing.T) {
	Convey("BufferClone behaviour", t, func() {
		Convey("nil input returns nil", func() {
			So(errcmd.BufferClone(nil), ShouldBeNil)
		})

		Convey("zero-capacity buffer returns fresh empty buffer", func() {
			out := errcmd.BufferClone(&bytes.Buffer{})
			So(out, ShouldNotBeNil)
			So(out.Len(), ShouldEqual, 0)
		})

		Convey("buffer with capacity & data is cloned with same bytes", func() {
			src := &bytes.Buffer{}
			src.Grow(64)
			src.WriteString("hello")
			out := errcmd.BufferClone(src)
			So(out, ShouldNotBeNil)
			So(out.String(), ShouldEqual, "hello")
		})
	})
}

func Test_MoreCoverage_StdIn(t *testing.T) {
	Convey("StdIn constructor + accessors", t, func() {
		empty := errcmd.NewStdIn(0, 0)
		So(empty.HasStdOut(), ShouldBeFalse)
		So(empty.HasStdErr(), ShouldBeFalse)
		So(empty.StdOutString(), ShouldEqual, "")
		So(empty.StdErrString(), ShouldEqual, "")

		full := errcmd.NewStdIn(16, 16)
		So(full.HasStdOut(), ShouldBeTrue)
		So(full.HasStdErr(), ShouldBeTrue)
		full.OutBuf.WriteString("o")
		full.ErrBuf.WriteString("e")
		So(full.StdOutString(), ShouldEqual, "o")
		So(full.StdErrString(), ShouldEqual, "e")
	})
}

func Test_MoreCoverage_NewCreator(t *testing.T) {
	Convey("errcmd.New.Empty produces an issue-laden CmdOnce", t, func() {
		empty := errcmd.New.Empty()
		So(empty, ShouldNotBeNil)
		So(empty.HasAnyIssues(), ShouldBeTrue)
		So(empty.HasIssues(), ShouldBeTrue)
		// Dispose paths
		empty.DisposeWithoutResult()
		empty.Dispose()
	})

	Convey("errcmd.New.Default constructs CmdOnce with expected metadata", t, func() {
		one := errcmd.New.Default("echo", "hello", "world")
		So(one, ShouldNotBeNil)
		So(one.ProcessName(), ShouldEqual, "echo")
		So(one.Arguments(), ShouldResemble, []string{"hello", "world"})
		So(one.ArgumentsSingleLine(), ShouldEqual, "hello world")
		So(one.WholeCommandLine(), ShouldEqual, "echo hello world")
		So(one.DoubleQuoteWholeCommandLine(), ShouldContainSubstring, "echo hello world")
		So(one.HasCmd(), ShouldBeTrue)
		_ = one.IsAlreadyRan()
		_ = one.IsNull()
		So(one.CommandLine(), ShouldEqual, "echo hello world")
		_ = one.GetCommandLineDataDependingOnSecurity()

		Convey("Clone reproduces the same surface", func() {
			cloned := one.Clone()
			So(cloned.ProcessName(), ShouldEqual, one.ProcessName())
			So(cloned.WholeCommandLine(), ShouldEqual, one.WholeCommandLine())
		})

		Convey("NewArgs appends additional arguments", func() {
			extra := one.NewArgs("--flag", "v")
			So(extra.Arguments(), ShouldResemble, []string{"hello", "world", "--flag", "v"})
		})

		Convey("CloneCmd returns underlying exec.Cmd in both modes", func() {
			So(one.CloneCmd(false), ShouldNotBeNil)
			So(one.CloneCmd(true), ShouldNotBeNil)
			So(one.CmdCloneWithoutStd(), ShouldNotBeNil)
		})

		Convey("CmdCloneUsingStds wires provided buffers", func() {
			out := &bytes.Buffer{}
			errBuf := &bytes.Buffer{}
			cmd := one.CmdCloneUsingStds(out, errBuf)
			So(cmd, ShouldNotBeNil)
		})

		Convey("GetFormattedKeyValueData emits KEY=VAL", func() {
			So(one.GetFormattedKeyValueData("A", "B"), ShouldEqual, "A=B")
		})
	})

	Convey("Secure-data variant masks the command line", t, func() {
		secure := errcmd.New.Create(true, true, "echo", "secret")
		So(secure.GetCommandLineDataDependingOnSecurity(), ShouldEqual, "echo")
	})

	Convey("UsingCmd / CreateUsingStdIns construct valid wrappers", t, func() {
		So(errcmd.New.UsingCmd(true, false, nil), ShouldNotBeNil)
		stdIn := errcmd.NewStdIn(8, 8)
		c := errcmd.New.CreateUsingStdIns(true, false, &bytes.Buffer{}, stdIn.OutBuf, stdIn.ErrBuf, "echo", "x")
		So(c, ShouldNotBeNil)
		So(c.HasCmd(), ShouldBeTrue)
	})
}

func Test_MoreCoverage_NewCmd(t *testing.T) {
	Convey("New.Cmd / CmdOsStd return non-nil exec.Cmd + joined command", t, func() {
		c, line := errcmd.New.Cmd.Cmd("echo", "hi")
		So(c, ShouldNotBeNil)
		So(line, ShouldContainSubstring, "echo")

		c2, line2 := errcmd.New.Cmd.CmdOsStd("echo", "hi")
		So(c2, ShouldNotBeNil)
		So(line2, ShouldContainSubstring, "echo")
	})
}

func Test_MoreCoverage_CmdSetOsStandardWriters(t *testing.T) {
	Convey("nil input returns nil", t, func() {
		So(errcmd.CmdSetOsStandardWriters(nil), ShouldBeNil)
	})
}

func Test_MoreCoverage_Result(t *testing.T) {
	Convey("NewResult with no error reports success", t, func() {
		r := errcmd.NewResult(nil, nil, nil, nil, true)
		So(r, ShouldNotBeNil)
		So(r.HasCommandExecuted(), ShouldBeTrue)
		So(r.IsRunSuccessfully(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.IsInvalid(), ShouldBeFalse)
		So(r.IsInvalidExitCode(), ShouldBeFalse)
		So(r.HasValidExitCode(), ShouldBeTrue)
		So(r.HasExitError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.HasNoError(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasAnyError(), ShouldBeFalse)
		So(r.ExitCode, ShouldEqual, errcmd.SuccessfullyRunningExitCode)
		So(r.IsExitCode(errcmd.SuccessfullyRunningExitCode), ShouldBeTrue)
		So(r.IsExitCodeByte(0), ShouldBeFalse) // ExitCode 0 fails the > 0 guard
		So(r.ExitCodeByte(), ShouldEqual, byte(0))
		So(r.ErrorWrapper(), ShouldBeNil)
		So(r.CompiledFullErrorWrapper(), ShouldBeNil)

		eb, ew := r.AllErrorBytes()
		So(eb, ShouldBeNil)
		So(ew, ShouldBeNil)
		es, ew2 := r.AllErrorString()
		So(es, ShouldEqual, "")
		So(ew2, ShouldBeNil)

		// DetailedOutput / String paths
		_ = r.DetailedOutput()
		_ = r.String()
		_ = r.StringIf(true)
		_ = r.StringIf(false)
		r.HandleError() // no-op when no error
		r.Dispose()
	})

	Convey("NewResultUsingBaseBuffer accepts a nil buffer", t, func() {
		r := errcmd.NewResultUsingBaseBuffer(nil, nil, nil, false)
		So(r, ShouldNotBeNil)
		So(r.HasCommandExecuted(), ShouldBeFalse)
		r.DisposeWithoutErrorWrapper()
	})
}

func Test_MoreCoverage_CurrentOsDetails(t *testing.T) {
	Convey("CurrentOsDetails returns details on supported OS", t, func() {
		details, err := errcmd.CurrentOsDetails()
		// Either we got details (Linux/macOS/Windows known) or a wrapped error.
		if err != nil {
			So(err.HasError(), ShouldBeTrue)
		} else {
			So(details, ShouldNotBeNil)
		}
	})
}

func Test_MoreCoverage_CmdOnceCollection(t *testing.T) {
	Convey("Empty constructors", t, func() {
		So(errcmd.NewCmdOnceCollection2().IsEmpty(), ShouldBeTrue)
		So(errcmd.NewCmdOnceCollection(4).Length(), ShouldEqual, 0)
		So(errcmd.NewCmdOnceCollectionUsingLinesOfScripts(scripttype.Bash).IsEmpty(), ShouldBeTrue)
		So(errcmd.NewCmdOnceCollectionUsingLinesDirect(scripttype.Bash, true).IsEmpty(), ShouldBeTrue)
		So(errcmd.BashScriptsLinesCmdOneCollection(true).IsEmpty(), ShouldBeTrue)
		So(errcmd.NewCmdOnceCollectionUsingLinesPtr(scripttype.Bash, true, nil).IsEmpty(), ShouldBeTrue)
	})

	Convey("Populated collection exposes count / getters", t, func() {
		coll := errcmd.NewCmdOnceCollectionUsingLinesOfScripts(
			scripttype.Bash,
			"echo a",
			"echo b",
		)
		So(coll.IsEmpty(), ShouldBeFalse)
		So(coll.HasAnyItem(), ShouldBeTrue)
		So(coll.Length(), ShouldEqual, 2)
		So(coll.Count(), ShouldEqual, 2)
		So(coll.LastIndex(), ShouldEqual, 1)
		So(coll.HasIndex(0), ShouldBeTrue)
		So(coll.HasIndex(99), ShouldBeFalse)
		So(len(coll.Items()), ShouldEqual, 2)
		So(len(coll.CommandsStrings()), ShouldEqual, 2)

		Convey("Add / AddMany append items", func() {
			extra := errcmd.New.Default("echo", "c")
			coll.Add(extra).Add(nil)
			coll.AddMany(errcmd.New.Default("echo", "d"))
			So(coll.Length(), ShouldEqual, 4)
		})

		Convey("Clone reproduces the same shape", func() {
			cloned := coll.Clone()
			So(cloned.Length(), ShouldEqual, coll.Length())
		})

		Convey("AddDefaultScript / AddDefaultScriptArgs", func() {
			c2 := errcmd.NewCmdOnceCollection2()
			c2.AddDefaultScript(scripttype.Bash, "echo x")
			c2.AddDefaultScriptArgs(scripttype.Bash, "echo", "y")
			c2.AddBashArgsDefault("echo", "z")
			So(c2.Length(), ShouldBeGreaterThanOrEqualTo, 1)
		})

		Convey("AsBasicSlice contracts binders are non-nil", func() {
			So(coll.AsBasicSliceContractsBinder(), ShouldNotBeNil)
			So(coll.AsBasicSlicerContractsBinder(), ShouldNotBeNil)
		})
	})
}
