package errnewtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
	. "github.com/smartystreets/goconvey/convey"
)

// MoreCoverage_test.go — exercises a wide surface of the errnew creator
// families to lift coverage. Most assertions only check non-nil + Type
// because the constructors are simple plumbing layers.

func Test_MoreCoverage_Helpers(t *testing.T) {
	Convey("Empty / OnEmpty / NotImpl helpers", t, func() {
		So(errnew.Empty(), ShouldNotBeNil)
		So(errnew.OnEmpty(nil), ShouldNotBeNil)
		non := errorwrapper.NewMsgDisplayErrorNoReference(codestack.Skip1, errtype.IO, "x")
		So(errnew.OnEmpty(non), ShouldEqual, non)

		So(errnew.NotImpl("/foo").Type(), ShouldEqual, errtype.NotImplemented)
		So(errnew.NotImplPtrUsingStackSkip(0, "/bar").Type(), ShouldEqual, errtype.NotImplemented)

		So(errnew.WasExpecting(errtype.InvalidInput, "title", 1, 2), ShouldNotBeNil)
		So(errnew.WasExpectingUsingStackSkip(0, errtype.InvalidInput, "title", 1, 2), ShouldNotBeNil)

		So(errnew.NotSupportedOption("varX", 42, "msg").Type(), ShouldEqual, errtype.NotSupportedOption)
		So(errnew.NotSupportedOptionUsingStackSkip(0, "varX", 42, "msg").Type(), ShouldEqual, errtype.NotSupportedOption)

		So(errnew.OutOfRange(0, 10, "1..9", "value").Type(), ShouldEqual, errtype.OutOfRangeValue)
	})
}

func Test_MoreCoverage_Messages(t *testing.T) {
	Convey("Messages creator family", t, func() {
		r := ref.New("k", "v")
		So(errnew.Messages.WithRef(errtype.IO, r, "m1").HasError(), ShouldBeTrue)
		So(errnew.Messages.WithRefUsingStackSkip(0, errtype.IO, r, "m1").HasError(), ShouldBeTrue)
		So(errnew.Messages.WithOnlyRefs(errtype.IO, r), ShouldNotBeNil)
		So(errnew.Messages.Single(errtype.IO, "x").Type(), ShouldEqual, errtype.IO)
		So(errnew.Messages.SingleUsingStackSkip(0, errtype.IO, "x").Type(), ShouldEqual, errtype.IO)
		So(errnew.Messages.Many(errtype.IO, "a", "b").Type(), ShouldEqual, errtype.IO)
		So(errnew.Messages.ManyUsingStackSkip(0, errtype.IO, "a", "b").Type(), ShouldEqual, errtype.IO)
		So(errnew.Messages.Create(errtype.IO, "a", "b"), ShouldNotBeNil)

		err := errors.New("boom")
		So(errnew.Messages.ErrorWithRef(errtype.IO, err, r, "ctx").HasError(), ShouldBeTrue)
		So(errnew.Messages.ErrorWithRef(errtype.IO, nil, r, "ctx"), ShouldBeNil)
		So(errnew.Messages.ErrorWithRefUsingStackSkip(0, errtype.IO, err, r, "ctx").HasError(), ShouldBeTrue)
		So(errnew.Messages.ErrorWithManyUsingStackSkip(0, errtype.IO, err, "ctx").HasError(), ShouldBeTrue)
		So(errnew.Messages.ErrorWithMany(errtype.IO, nil, "ctx"), ShouldBeNil)
	})
}

func Test_MoreCoverage_Message(t *testing.T) {
	Convey("Message creator family", t, func() {
		So(errnew.Message.Create(errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.CreateUsingStackSkip(0, errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.Type(errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.TypeUsingStackSkip(0, errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.New(errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.NewUsingStackSkip(0, errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.Default(errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.Many(errtype.IO, "a", "b").Type(), ShouldEqual, errtype.IO)
		So(errnew.Message.ManyUsingStackSkip(0, errtype.IO, "a", "b").Type(), ShouldEqual, errtype.IO)

		err := errors.New("e")
		So(errnew.Message.ErrorWithMany(errtype.IO, err, "ctx").HasError(), ShouldBeTrue)
		So(errnew.Message.ErrorWithMany(errtype.IO, nil, "ctx"), ShouldBeNil)
		So(errnew.Message.ErrorWithManyUsingStackSkip(0, errtype.IO, err, "ctx").HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage_Type(t *testing.T) {
	Convey("Type creator family", t, func() {
		So(errnew.Type.Default(errtype.IO).Type(), ShouldEqual, errtype.IO)
		So(errnew.Type.DefaultUsingStackSkip(0, errtype.IO).Type(), ShouldEqual, errtype.IO)
		So(errnew.Type.UsingStackSkip(0, errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Type.Create(errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Type.New(errtype.IO, errors.New("m")).Type(), ShouldEqual, errtype.IO)
		So(errnew.Type.Message(errtype.IO, "m").Type(), ShouldEqual, errtype.IO)
		So(errnew.Type.Messages(errtype.IO, "a", "b").Type(), ShouldEqual, errtype.IO)
		So(errnew.Type.MessagesUsingStackSkip(0, errtype.IO, "a", "b").Type(), ShouldEqual, errtype.IO)

		err := errors.New("e")
		So(errnew.Type.Error(errtype.IO, err).HasError(), ShouldBeTrue)
		So(errnew.Type.Error(errtype.IO, nil), ShouldBeNil)
		So(errnew.Type.ErrorUsingStackSkip(0, errtype.IO, err).HasError(), ShouldBeTrue)
		emptyRefs := refs.Empty()
		So(errnew.Type.ErrorWithMessage(errtype.IO, err, "ctx", emptyRefs).HasError(), ShouldBeTrue)
		So(errnew.Type.ErrorWithMessageUsingStackSkip(0, errtype.IO, err, "ctx", emptyRefs).HasError(), ShouldBeTrue)
		So(errnew.Type.ErrorWithMessages(errtype.IO, err, "a", "b").HasError(), ShouldBeTrue)
		So(errnew.Type.ErrorWithMessagesUsingStackSkip(0, errtype.IO, err, "a", "b").HasError(), ShouldBeTrue)

		So(errnew.Type.Marshal(err).HasError(), ShouldBeTrue)
		So(errnew.Type.Unmarshal(err).HasError(), ShouldBeTrue)

		r := ref.New("k", "v")
		So(errnew.Type.Refs(errtype.IO, r), ShouldNotBeNil)
		So(errnew.Type.References(errtype.IO, r), ShouldNotBeNil)
		So(errnew.Type.DirectRefs(errtype.IO, "m", "k", "v"), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Error(t *testing.T) {
	Convey("Error creator family", t, func() {
		err := errors.New("e")
		So(errnew.Error.TypeOnly(errtype.IO).Type(), ShouldEqual, errtype.IO)
		So(errnew.Error.TypeOnlyUsingStackSkip(0, errtype.IO).Type(), ShouldEqual, errtype.IO)
		So(errnew.Error.NoType(err).HasError(), ShouldBeTrue)
		So(errnew.Error.NoTypeUsingStackSkip(0, err).HasError(), ShouldBeTrue)
		So(errnew.Error.Default(errtype.IO, err), ShouldNotBeNil)
		So(errnew.Error.UsingStackSkip(0, errtype.IO, err).HasError(), ShouldBeTrue)
		So(errnew.Error.DefaultUsingStackSkip(0, errtype.IO, err), ShouldNotBeNil)
		So(errnew.Error.Create(errtype.IO, err).HasError(), ShouldBeTrue)
		So(errnew.Error.Type(errtype.IO, err).HasError(), ShouldBeTrue)
		So(errnew.Error.TypeUsingStackSkip(0, errtype.IO, err).HasError(), ShouldBeTrue)
		So(errnew.Error.ErrorWithMsg(err, "msg"), ShouldNotBeNil)
		So(errnew.Error.TypeMsg(errtype.IO, err, "msg").HasError(), ShouldBeTrue)
		So(errnew.Error.TypeMsgUsingStackSkip(0, errtype.IO, err, "msg").HasError(), ShouldBeTrue)
		So(errnew.Error.TypeMsgRef(errtype.IO, err, "msg", "k", "v").HasError(), ShouldBeTrue)
		So(errnew.Error.TypeMsgRefUsingStackSkip(0, errtype.IO, err, "msg", "k", "v").HasError(), ShouldBeTrue)
		So(errnew.Error.TypeMessages(errtype.IO, err, "a", "b").HasError(), ShouldBeTrue)
		So(errnew.Error.TypeMessagesUsingStackSkip(0, errtype.IO, err, "a", "b").HasError(), ShouldBeTrue)
		So(errnew.Error.TypeWithMessagesUsingStackSkip(0, errtype.IO, err, "a", "b").HasError(), ShouldBeTrue)

		r := ref.New("k", "v")
		So(errnew.Error.TypeWithRefs(errtype.IO, err, r).HasError(), ShouldBeTrue)
		So(errnew.Error.TypeWithRefsUsingStackSkip(0, errtype.IO, err, r).HasError(), ShouldBeTrue)
		So(errnew.Error.TypeWithMessageRefs(errtype.IO, err, "msg", r).HasError(), ShouldBeTrue)

		// nil-error short circuits
		So(errnew.Error.NoType(nil), ShouldBeNil)
		So(errnew.Error.TypeMsg(errtype.IO, nil, "msg"), ShouldBeNil)
	})
}

func Test_MoreCoverage_NotFound(t *testing.T) {
	Convey("NotFound creator family", t, func() {
		So(errnew.NotFound.Reference("ref").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.MessageRef("m", "ref").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.MessageRefName("m", "var", "v").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.Missing("m", "r1", "r2").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.Invalid("m", "r").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.InvalidData("m", "r").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.InvalidStatus("m", "r"), ShouldNotBeNil)
		So(errnew.NotFound.InvalidBytes("m", []byte("r")), ShouldNotBeNil)
		So(errnew.NotFound.All(errtype.IO, "m", "r").Type(), ShouldEqual, errtype.IO)
		So(errnew.NotFound.Message("m").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.Message(""), ShouldBeNil)
		So(errnew.NotFound.MessageError("m", errors.New("e")).HasError(), ShouldBeTrue)
		So(errnew.NotFound.MessageError("m", nil), ShouldBeNil)
		So(errnew.NotFound.MessageReference("m", "r").Type(), ShouldEqual, errtype.NotFound)
		So(errnew.NotFound.MessageErrorReference("m", errors.New("e"), "r").HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage_Path(t *testing.T) {
	Convey("Path creator family", t, func() {
		err := errors.New("e")
		So(errnew.Path.Dir(err, "/a").HasError(), ShouldBeTrue)
		So(errnew.Path.Dir(nil, "/a"), ShouldBeNil)
		So(errnew.Path.NotDir("/a"), ShouldNotBeNil)
		So(errnew.Path.NotFile("/a"), ShouldNotBeNil)
		So(errnew.Path.Invalid("/a"), ShouldNotBeNil)
		So(errnew.Path.InvalidUsingStackSkip(0, "/a"), ShouldNotBeNil)
		So(errnew.Path.InvalidMany("/a", "/b"), ShouldNotBeNil)
		So(errnew.Path.InvalidMany(), ShouldBeNil)
		So(errnew.Path.InvalidManyUsingStackSkip(0, "/a"), ShouldNotBeNil)
		So(errnew.Path.Empty(), ShouldNotBeNil)
		So(errnew.Path.File(err, "/a"), ShouldNotBeNil)
		So(errnew.Path.FileContentIssue("msg", "/a"), ShouldNotBeNil)
		So(errnew.Path.Type(errtype.IO, "/a").Type(), ShouldEqual, errtype.IO)
		So(errnew.Path.TypeMsg(errtype.IO, "/a", "msg"), ShouldNotBeNil)
		So(errnew.Path.TypeMsgManyPaths(errtype.IO, "msg", "/a", "/b"), ShouldNotBeNil)
		So(errnew.Path.TypeUsingStackSkip(0, errtype.IO, "/a"), ShouldNotBeNil)
		So(errnew.Path.Error(err, "/a").HasError(), ShouldBeTrue)
		So(errnew.Path.Messages(errtype.IO, "/a", "m1", "m2"), ShouldNotBeNil)
		So(errnew.Path.ErrorMessages(errtype.IO, err, "/a", "m1"), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Merge(t *testing.T) {
	Convey("Merge creator family", t, func() {
		a := errnew.Messages.Single(errtype.IO, "a")
		b := errnew.Messages.Single(errtype.IO, "b")
		c := errnew.Messages.Single(errtype.IO, "c")
		So(errnew.Merge.New(a, b), ShouldNotBeNil)
		So(errnew.Merge.New(nil, nil), ShouldBeNil)
		So(errnew.Merge.Three(a, b, c), ShouldNotBeNil)
		So(errnew.Merge.Many(errtype.IO, a, b, c), ShouldNotBeNil)
		So(errnew.Merge.Many(errtype.IO), ShouldBeNil)
		So(errnew.Merge.ManyStackSkip(0, errtype.IO, a, b), ShouldNotBeNil)
		So(errnew.Merge.UsingStackSkip(0, a, b), ShouldNotBeNil)
		So(errnew.Merge.UsingNewType(errtype.IO, a, b), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_FromTo(t *testing.T) {
	Convey("FromTo creator family", t, func() {
		So(errnew.FromTo.Create(errtype.IO, "from", "to").Type(), ShouldEqual, errtype.IO)
		So(errnew.FromTo.CreateUsingStackSkip(0, errtype.IO, "from", "to"), ShouldNotBeNil)
		So(errnew.FromTo.Message(errtype.IO, "msg", "from", "to"), ShouldNotBeNil)
		So(errnew.FromTo.MessageStackSkip(0, errtype.IO, "msg", "from", "to"), ShouldNotBeNil)
		So(errnew.FromTo.Messages(errtype.IO, true, "from", "to", "m1", "m2"), ShouldNotBeNil)
		So(errnew.FromTo.MessagesUsingStackSkip(0, errtype.IO, true, "from", "to", "m1"), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Null(t *testing.T) {
	Convey("Null creator family", t, func() {
		So(errnew.Null.Simple(nil), ShouldNotBeNil)
		So(errnew.Null.Simple("not-null"), ShouldBeNil)
		So(errnew.Null.UsingStackSkip(0, "ctx", nil), ShouldNotBeNil)
		So(errnew.Null.Message("hello"), ShouldNotBeNil)
		So(errnew.Null.Message(""), ShouldBeNil)
		So(errnew.Null.WithMessage("ctx", nil), ShouldNotBeNil)
		So(errnew.Null.Error(errors.New("e")), ShouldNotBeNil)
		So(errnew.Null.Error(nil), ShouldBeNil)
		So(errnew.Null.ErrorWithMessage("ctx", errors.New("e")), ShouldNotBeNil)
		So(errnew.Null.OrWrapper(errnew.Messages.Single(errtype.IO, "fallback"), nil), ShouldNotBeNil)
		So(errnew.Null.OrWrapper(nil, "notnull"), ShouldBeNil)
		So(errnew.Null.OrError(nil, errors.New("e")), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_MessageWithRef(t *testing.T) {
	Convey("MessageWithRef family", t, func() {
		So(errnew.MessageWithRef.Default(errtype.IO, "m", "r").Type(), ShouldEqual, errtype.IO)
		So(errnew.MessageWithRef.DefaultVarName(errtype.IO, "m", "var", "v"), ShouldNotBeNil)
		So(errnew.MessageWithRef.Error(errtype.IO, errors.New("e"), "r"), ShouldNotBeNil)
		So(errnew.MessageWithRef.ErrorVarName(errtype.IO, errors.New("e"), "var", "v"), ShouldNotBeNil)
		So(errnew.MessageWithRef.ErrorRefs(errtype.IO, errors.New("e"), "m", ref.New("k", "v")), ShouldNotBeNil)
		So(errnew.MessageWithRef.References(errtype.IO, "m", ref.New("k", "v")), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Range(t *testing.T) {
	Convey("Range creator family", t, func() {
		So(errnew.Range.Within(5, 1, 10).Type(), ShouldEqual, errtype.OutOfRangeValue)
		So(errnew.Range.OutOf(5, 1, 10), ShouldNotBeNil)
		So(errnew.Range.MessageOutOf("msg", "ref"), ShouldNotBeNil)
		So(errnew.Range.Error(errors.New("e")), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Fmt(t *testing.T) {
	Convey("Fmt formatter family", t, func() {
		So(errnew.Fmt.Default(errtype.IO, "msg %d", 7).Type(), ShouldEqual, errtype.IO)
		So(errnew.Fmt.Format(errtype.IO, "msg %s", "x"), ShouldNotBeNil)
		So(errnew.Fmt.MessageRef(errtype.IO, "m %d", 1), ShouldNotBeNil)
		So(errnew.Fmt.MessageRefs(errtype.IO, "m %d", 1), ShouldNotBeNil)
		So(errnew.Fmt.ErrorRefs(errtype.IO, errors.New("e"), "m %d", 1), ShouldNotBeNil)
		So(errnew.Fmt.Error(errtype.IO, errors.New("e"), "m %d", 1), ShouldNotBeNil)
		So(errnew.Fmt.Message(errtype.IO, "m %d", 1), ShouldNotBeNil)
		So(errnew.Fmt.MessageError(errtype.IO, errors.New("e"), "m %s", "x"), ShouldNotBeNil)
		So(errnew.Fmt.If(true, errtype.IO, "m %s", "x"), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_SrcDst(t *testing.T) {
	Convey("SrcDst creator family", t, func() {
		So(errnew.SrcDst.Create(errtype.IO, "s", "d").Type(), ShouldEqual, errtype.IO)
		So(errnew.SrcDst.CreateUsingStackSkip(0, errtype.IO, "s", "d"), ShouldNotBeNil)
		So(errnew.SrcDst.Message(errtype.IO, "s", "d", "msg"), ShouldNotBeNil)
		So(errnew.SrcDst.Messages(errtype.IO, "s", "d", "m1", "m2"), ShouldNotBeNil)
		So(errnew.SrcDst.MessagesUsingStackSkip(0, errtype.IO, "s", "d", "m1"), ShouldNotBeNil)
		So(errnew.SrcDst.Error(errtype.IO, errors.New("e"), "s", "d"), ShouldNotBeNil)
		So(errnew.SrcDst.ErrorUsingStackSkip(0, errtype.IO, errors.New("e"), "s", "d"), ShouldNotBeNil)
		So(errnew.SrcDst.ErrorWithMessage(errtype.IO, errors.New("e"), "s", "d", "ctx"), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Refs(t *testing.T) {
	Convey("Refs creator family", t, func() {
		r := ref.New("k", "v")
		err := errors.New("e")
		So(errnew.Refs.Error(errtype.IO, err, r).HasError(), ShouldBeTrue)
		So(errnew.Refs.Error(errtype.IO, nil, r), ShouldBeNil)
		So(errnew.Refs.ErrorUsingStackSkip(0, errtype.IO, err, r), ShouldNotBeNil)
		So(errnew.Refs.ErrorWithMessage(errtype.IO, err, "ctx", r), ShouldNotBeNil)
		So(errnew.Refs.ErrorWithMessageUsingStackSkip(0, errtype.IO, err, "ctx", r), ShouldNotBeNil)
		So(errnew.Refs.Msg(errtype.IO, "m", r), ShouldNotBeNil)
		So(errnew.Refs.MsgUsingStackSkip(0, errtype.IO, "m", r), ShouldNotBeNil)
		So(errnew.Refs.Type(errtype.IO, r), ShouldNotBeNil)
		So(errnew.Refs.UsingStackSkip(0, errtype.IO, r), ShouldNotBeNil)
		So(errnew.Refs.Many(errtype.IO, r, r), ShouldNotBeNil)
		So(errnew.Refs.ManyUsingStackSkip(0, errtype.IO, r), ShouldNotBeNil)
		So(errnew.Refs.OnlyOne(errtype.IO, r), ShouldNotBeNil)
		So(errnew.Refs.OnlyOneUsingStackSkip(0, errtype.IO, r), ShouldNotBeNil)
		coll := refs.New(1).Add("k", "v")
		So(errnew.Refs.Messages(errtype.IO, coll, "a", "b"), ShouldNotBeNil)
		So(errnew.Refs.MessagesUsingStackSkip(0, errtype.IO, coll, "a"), ShouldNotBeNil)
		So(errnew.Refs.MessagesUsingJoiner(errtype.IO, coll, " | ", "a", "b"), ShouldNotBeNil)
		So(errnew.Refs.MessagesUsingJoinerStackSkip(0, errtype.IO, coll, " | ", "a"), ShouldNotBeNil)
		So(errnew.Refs.TypeQuick(errtype.IO, "k", "v"), ShouldNotBeNil)
		So(errnew.Refs.TypeQuickUsingStackSkip(0, errtype.IO, "k", "v"), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Ref(t *testing.T) {
	Convey("Ref creator family (subset)", t, func() {
		err := errors.New("e")
		So(errnew.Ref.ErrorOne(errtype.IO, err, "k", "v").HasError(), ShouldBeTrue)
		So(errnew.Ref.ErrorOne(errtype.IO, nil, "k", "v"), ShouldBeNil)
		So(errnew.Ref.ErrorOneUsingStackSkip(0, errtype.IO, err, "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.ErrorWithMessage(errtype.IO, err, "msg", "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.ErrorWithMessageDefault(err, "msg", "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.Default(errtype.IO, "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.OnlyOne(errtype.IO, "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.OnlyOneUsingStackSkip(0, errtype.IO, "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.MsgOne(errtype.IO, "msg", "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.MsgOneUsingStackSkip(0, errtype.IO, "msg", "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.TypeQuick(errtype.IO, "k", "v"), ShouldNotBeNil)
		So(errnew.Ref.TypeQuickUsingStackSkip(0, errtype.IO, "k", "v"), ShouldNotBeNil)
	})
}
