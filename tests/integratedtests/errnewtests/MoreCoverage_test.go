package errnewtests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

func Test_Type_Creators(t *testing.T) {
	Convey("Type.Message + Type.Messages", t, func() {
		w := errnew.Type.Message(errtype.InvalidInput, "bad")
		So(w, ShouldNotBeNil)
		So(w.Type(), ShouldEqual, errtype.InvalidInput)
		So(w.FullString(), ShouldContainSubstring, "bad")

		w2 := errnew.Type.Messages(errtype.InvalidInput, "a", "b")
		So(w2, ShouldNotBeNil)
		So(w2.FullString(), ShouldContainSubstring, "a")
	})

	Convey("Type.Error nil and non-nil", t, func() {
		So(errnew.Type.Error(errtype.IO, nil), ShouldBeNil)
		w := errnew.Type.Error(errtype.IO, errors.New("boom"))
		So(w, ShouldNotBeNil)
		So(w.FullString(), ShouldContainSubstring, "boom")
	})

	Convey("Type.New, Marshal, Unmarshal nil-safe", t, func() {
		So(errnew.Type.New(errtype.IO, nil), ShouldBeNil)
		So(errnew.Type.Marshal(nil), ShouldBeNil)
		So(errnew.Type.Unmarshal(nil), ShouldBeNil)

		So(errnew.Type.Marshal(errors.New("m")), ShouldNotBeNil)
		So(errnew.Type.Unmarshal(errors.New("u")), ShouldNotBeNil)
	})

	Convey("Type.ErrorWithMessage nil + non-nil", t, func() {
		coll := refs.New(1)
		So(errnew.Type.ErrorWithMessage(errtype.IO, nil, "msg", *coll), ShouldBeNil)

		w := errnew.Type.ErrorWithMessage(errtype.IO, errors.New("e"), "msg", *coll)
		So(w, ShouldNotBeNil)
		So(w.FullString(), ShouldContainSubstring, "e")
	})

	Convey("Type.ErrorWithMessages nil + non-nil", t, func() {
		So(errnew.Type.ErrorWithMessages(errtype.IO, nil, "x"), ShouldBeNil)
		w := errnew.Type.ErrorWithMessages(errtype.IO, errors.New("e"), "x", "y")
		So(w, ShouldNotBeNil)
		So(w.FullString(), ShouldContainSubstring, "e")
	})

	Convey("Type.Default + DefaultUsingStackSkip + Create", t, func() {
		So(errnew.Type.Default(errtype.NoError), ShouldBeNil)
		So(errnew.Type.DefaultUsingStackSkip(0, errtype.NoError), ShouldBeNil)

		w := errnew.Type.Default(errtype.IO)
		So(w, ShouldNotBeNil)
		So(w.Type(), ShouldEqual, errtype.IO)

		w2 := errnew.Type.DefaultUsingStackSkip(0, errtype.IO)
		So(w2, ShouldNotBeNil)
	})

	Convey("Type.MessagesUsingStackSkip + ErrorWithMessagesUsingStackSkip", t, func() {
		w := errnew.Type.MessagesUsingStackSkip(0, errtype.IO, "x")
		So(w, ShouldNotBeNil)
		So(errnew.Type.ErrorWithMessagesUsingStackSkip(0, errtype.IO, nil, "x"), ShouldBeNil)
		w2 := errnew.Type.ErrorWithMessagesUsingStackSkip(0, errtype.IO, errors.New("e"), "x")
		So(w2, ShouldNotBeNil)
	})

	Convey("Type.ErrorWithMessageUsingStackSkip nil-safe", t, func() {
		coll := refs.New(1)
		So(errnew.Type.ErrorWithMessageUsingStackSkip(0, errtype.IO, nil, "m", *coll), ShouldBeNil)
		w := errnew.Type.ErrorWithMessageUsingStackSkip(0, errtype.IO, errors.New("e"), "m", *coll)
		So(w, ShouldNotBeNil)
	})

	Convey("Type.DirectRefs + Refs + References", t, func() {
		w := errnew.Type.DirectRefs(errtype.IO, "ctx", "a", "b")
		So(w, ShouldNotBeNil)
		So(w.FullString(), ShouldContainSubstring, "ctx")
	})
}

func Test_Error_Creators(t *testing.T) {
	Convey("Error.Many nil + non-nil", t, func() {
		So(errnew.Error.Many(errtype.IO), ShouldBeNil)
		w := errnew.Error.Many(errtype.IO, errtype.InvalidInput, errtype.NotFound)
		So(w, ShouldNotBeNil)
	})

	Convey("Error.NoTypeUsingStackSkip nil + non-nil", t, func() {
		So(errnew.Error.NoTypeUsingStackSkip(0, nil), ShouldBeNil)
		w := errnew.Error.NoTypeUsingStackSkip(0, errors.New("x"))
		So(w, ShouldNotBeNil)
	})

	Convey("Error.Default nil + non-nil", t, func() {
		So(errnew.Error.Default(errtype.IO, nil), ShouldBeNil)
		w := errnew.Error.Default(errtype.IO, errors.New("x"))
		So(w, ShouldNotBeNil)
	})
}

func Test_Path_Creators(t *testing.T) {
	Convey("Path basics", t, func() {
		So(errnew.Path.Invalid("/p"), ShouldNotBeNil)
		So(errnew.Path.NotDir("/p"), ShouldNotBeNil)
		So(errnew.Path.NotFile("/p"), ShouldNotBeNil)
		So(errnew.Path.Empty(), ShouldNotBeNil)
	})

	Convey("Path.Dir / File error-aware", t, func() {
		So(errnew.Path.Dir(nil, "/p"), ShouldBeNil)
		w := errnew.Path.Dir(errors.New("x"), "/p")
		So(w, ShouldNotBeNil)
		So(errnew.Path.File(nil, "/p"), ShouldBeNil)
		So(errnew.Path.File(errors.New("x"), "/p"), ShouldNotBeNil)
	})

	Convey("Path.InvalidMany + InvalidUsingStackSkip", t, func() {
		So(errnew.Path.InvalidMany("/p1", "/p2"), ShouldNotBeNil)
		So(errnew.Path.InvalidUsingStackSkip(0, "/p"), ShouldNotBeNil)
		So(errnew.Path.InvalidManyUsingStackSkip(0, "/p1", "/p2"), ShouldNotBeNil)
	})
}

func Test_Range_Creators(t *testing.T) {
	Convey("Range builders", t, func() {
		So(errnew.Range.Within(5, 1, 10), ShouldNotBeNil)
		So(errnew.Range.OutOf(20, 1, 10), ShouldNotBeNil)
		So(errnew.Range.OutOfRanges(20, 1, 5, 10), ShouldNotBeNil)
		So(errnew.Range.StartEnd(1, 10), ShouldNotBeNil)
		So(errnew.Range.MessageOutOf("msg", 5, 1, 10), ShouldNotBeNil)
		So(errnew.Range.RefStartEndValues("name", 1, 10, 20), ShouldNotBeNil)
	})
}

func Test_Formatter_Creators(t *testing.T) {
	Convey("Fmt builders produce wrappers", t, func() {
		So(errnew.Fmt.Default(errtype.IO, "x=%d", 1), ShouldNotBeNil)
		So(errnew.Fmt.Format(errtype.IO, "x=%d", 1), ShouldNotBeNil)
		So(errnew.Fmt.Fmt(errtype.IO, "x=%d", 1), ShouldNotBeNil)
		So(errnew.Fmt.Message(errtype.IO, "msg"), ShouldNotBeNil)
		So(errnew.Fmt.MessageRef(errtype.IO, "msg", "name", "val"), ShouldNotBeNil)
		So(errnew.Fmt.MessageRefs(errtype.IO, "msg", "a", "b"), ShouldNotBeNil)
	})

	Convey("Fmt.Error nil + non-nil", t, func() {
		So(errnew.Fmt.Error(errtype.IO, nil), ShouldBeNil)
		So(errnew.Fmt.Error(errtype.IO, errors.New("x")), ShouldNotBeNil)
		So(errnew.Fmt.MessageError(errtype.IO, "msg", nil), ShouldBeNil)
		So(errnew.Fmt.MessageError(errtype.IO, "msg", errors.New("x")), ShouldNotBeNil)
		So(errnew.Fmt.ErrorRefs(errtype.IO, errors.New("x"), "ref"), ShouldNotBeNil)
		So(errnew.Fmt.ErrorRefs(errtype.IO, nil, "ref"), ShouldBeNil)
	})

	Convey("Fmt.If gates by condition", t, func() {
		So(errnew.Fmt.If(false, errtype.IO, "x"), ShouldBeNil)
		So(errnew.Fmt.If(true, errtype.IO, "x"), ShouldNotBeNil)
	})
}

func Test_Null_Creators(t *testing.T) {
	Convey("Null.Simple nil-tagged", t, func() {
		var nilErr error = nil
		w := errnew.Null.Simple(nilErr)
		So(w, ShouldNotBeNil)
		So(w.Type(), ShouldEqual, errtype.Null)
	})

	Convey("Null.WithMessage / Message", t, func() {
		var v interface{} = nil
		So(errnew.Null.WithMessage("missing", v), ShouldNotBeNil)
		So(errnew.Null.Message("missing"), ShouldNotBeNil)
		So(errnew.Null.Message(""), ShouldBeNil)
	})

	Convey("Null.Error nil + non-nil", t, func() {
		So(errnew.Null.Error(nil), ShouldBeNil)
		So(errnew.Null.Error(errors.New("x")), ShouldNotBeNil)
		So(errnew.Null.ErrorWithMessage("msg", nil), ShouldBeNil)
		So(errnew.Null.ErrorWithMessage("msg", errors.New("x")), ShouldNotBeNil)
	})
}

func Test_MessageWithRef_Creators(t *testing.T) {
	Convey("MessageWithRef.Default + Error", t, func() {
		r := ref.New("k", "v")
		So(errnew.MessageWithRef.Default(errtype.IO, "msg", r), ShouldNotBeNil)
		So(errnew.MessageWithRef.Error(errtype.IO, errors.New("e"), r), ShouldNotBeNil)
		So(errnew.MessageWithRef.Error(errtype.IO, nil, r), ShouldBeNil)
		So(errnew.MessageWithRef.References(errtype.IO, "msg", r), ShouldNotBeNil)
		So(errnew.MessageWithRef.DefaultVarName(errtype.IO, "msg", "k", "v"), ShouldNotBeNil)
		So(errnew.MessageWithRef.ErrorVarName(errtype.IO, errors.New("e"), "k", "v"), ShouldNotBeNil)
	})
}

func Test_Message_Creators(t *testing.T) {
	Convey("Message.New + Many + Default", t, func() {
		So(errnew.Message.New(errtype.IO, "msg"), ShouldNotBeNil)
		So(errnew.Message.Many(errtype.IO, "a", "b"), ShouldNotBeNil)
		So(errnew.Message.Default(errtype.IO, "msg"), ShouldNotBeNil)
		So(errnew.Message.Type(errtype.IO, "msg"), ShouldNotBeNil)
	})

	Convey("Message.ErrorWithMany nil + non-nil", t, func() {
		So(errnew.Message.ErrorWithMany(errtype.IO, nil, "x"), ShouldBeNil)
		So(errnew.Message.ErrorWithMany(errtype.IO, errors.New("e"), "x"), ShouldNotBeNil)
	})
}

func Test_TopLevel_Helpers(t *testing.T) {
	Convey("Empty + OnEmpty", t, func() {
		So(errnew.Empty(), ShouldNotBeNil)
		So(errnew.OnEmpty(nil), ShouldNotBeNil)
		existing := errnew.Type.Default(errtype.IO)
		So(errnew.OnEmpty(existing), ShouldEqual, existing)
	})

	Convey("WasExpecting + stack-skip variant", t, func() {
		w := errnew.WasExpecting(errtype.IO, "mismatch", "a", "b")
		So(w, ShouldNotBeNil)
		So(w.FullString(), ShouldContainSubstring, "mismatch")
		So(errnew.WasExpectingUsingStackSkip(0, errtype.IO, "m", "a", "b"), ShouldNotBeNil)
	})

	Convey("NotSupportedOption + stack-skip", t, func() {
		So(errnew.NotSupportedOption("k", "v", "msg"), ShouldNotBeNil)
		So(errnew.NotSupportedOptionUsingStackSkip(0, "k", "v", "msg"), ShouldNotBeNil)
	})

	Convey("OutOfRange", t, func() {
		So(errnew.OutOfRange(1, 10, "1..10", "value out"), ShouldNotBeNil)
	})

	Convey("NotImpl + stack-skip", t, func() {
		So(errnew.NotImpl("http://x"), ShouldNotBeNil)
		So(errnew.NotImplPtrUsingStackSkip(0, "http://x"), ShouldNotBeNil)
	})
}

func Test_NotFound_Extras(t *testing.T) {
	Convey("Constant constructors return non-nil", t, func() {
		So(errnew.NotFound.Id(), ShouldNotBeNil)
		So(errnew.NotFound.Key(), ShouldNotBeNil)
		So(errnew.NotFound.Name(), ShouldNotBeNil)
		So(errnew.NotFound.Data(), ShouldNotBeNil)
		So(errnew.NotFound.Group(), ShouldNotBeNil)
		So(errnew.NotFound.Block(), ShouldNotBeNil)
		So(errnew.NotFound.Process(), ShouldNotBeNil)
		So(errnew.NotFound.Record(), ShouldNotBeNil)
		So(errnew.NotFound.User(), ShouldNotBeNil)
		So(errnew.NotFound.RedisCmd(), ShouldNotBeNil)
		So(errnew.NotFound.Bytes(), ShouldNotBeNil)
		So(errnew.NotFound.Payload(), ShouldNotBeNil)
		So(errnew.NotFound.EmptyPath(), ShouldNotBeNil)
		So(errnew.NotFound.EmptyString(), ShouldNotBeNil)
		So(errnew.NotFound.EmptyStatus(), ShouldNotBeNil)
		So(errnew.NotFound.EmptyStates(), ShouldNotBeNil)
	})

	Convey("Message-bearing constructors", t, func() {
		So(errnew.NotFound.Message("m"), ShouldNotBeNil)
		So(errnew.NotFound.Simple("m"), ShouldNotBeNil)
		So(errnew.NotFound.Missing("m"), ShouldNotBeNil)
		So(errnew.NotFound.Invalid("m", "a", "b"), ShouldNotBeNil)
		So(errnew.NotFound.MessageRef("m", "v"), ShouldNotBeNil)
		So(errnew.NotFound.MessageRefName("m", "k", "v"), ShouldNotBeNil)
		So(errnew.NotFound.MessageReference("m", "v"), ShouldNotBeNil)
		So(errnew.NotFound.IdMessage("m"), ShouldNotBeNil)
		So(errnew.NotFound.IdReference("v"), ShouldNotBeNil)
		So(errnew.NotFound.KeyMessage("m"), ShouldNotBeNil)
		So(errnew.NotFound.KeyReference("v"), ShouldNotBeNil)
		So(errnew.NotFound.NameMessage("m"), ShouldNotBeNil)
		So(errnew.NotFound.NameReference("v"), ShouldNotBeNil)
		So(errnew.NotFound.GroupReference("v"), ShouldNotBeNil)
		So(errnew.NotFound.PathMessage("/p"), ShouldNotBeNil)
		So(errnew.NotFound.DirMessage("/d"), ShouldNotBeNil)
		So(errnew.NotFound.FileMessage("/f"), ShouldNotBeNil)
		So(errnew.NotFound.BlockMessage("m"), ShouldNotBeNil)
		So(errnew.NotFound.PayloadMessage("m"), ShouldNotBeNil)
		So(errnew.NotFound.PayloadReference("v"), ShouldNotBeNil)
	})

	Convey("Error-bearing constructors nil-safe", t, func() {
		So(errnew.NotFound.PathErr(nil, "/p"), ShouldBeNil)
		So(errnew.NotFound.FileErr(nil, "/f"), ShouldBeNil)
		So(errnew.NotFound.DirErr(nil, "/d"), ShouldBeNil)
		So(errnew.NotFound.PayloadErr(nil), ShouldBeNil)
		So(errnew.NotFound.RecordError(nil), ShouldBeNil)
		So(errnew.NotFound.KeyError(nil), ShouldBeNil)

		So(errnew.NotFound.PathErr(errors.New("x"), "/p"), ShouldNotBeNil)
		So(errnew.NotFound.FileErr(errors.New("x"), "/f"), ShouldNotBeNil)
		So(errnew.NotFound.DirErr(errors.New("x"), "/d"), ShouldNotBeNil)
		So(errnew.NotFound.PayloadErr(errors.New("x")), ShouldNotBeNil)
		So(errnew.NotFound.RecordError(errors.New("x")), ShouldNotBeNil)
		So(errnew.NotFound.KeyError(errors.New("x")), ShouldNotBeNil)
	})
}
