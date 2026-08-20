package errnewtests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
	"github.com/alimtvnetwork/errorwrapper-v4/refs"
)

func Test_MoreCoverage4_Path_Remaining(t *testing.T) {
	Convey("Path FromTo / Messages / Marshal family", t, func() {
		err := errors.New("e")

		So(errnew.Path.FromToError(errtype.IO, err, "/a", "/b").HasError(), ShouldBeTrue)
		So(errnew.Path.FromToMessage(errtype.IO, "msg", "/a", "/b").HasError(), ShouldBeTrue)
		So(errnew.Path.FromToErrorMessage(errtype.IO, err, "msg", "/a", "/b").HasError(), ShouldBeTrue)

		So(errnew.Path.Messages(errtype.IO, "/a", "m1", "m2").HasError(), ShouldBeTrue)
		So(errnew.Path.MessagesUsingStackSkip(0, errtype.IO, "/a", "m1").HasError(), ShouldBeTrue)

		So(errnew.Path.ErrorMessages(errtype.IO, err, "/a", "m1").HasError(), ShouldBeTrue)
		So(errnew.Path.ErrorUsingStackSkip(0, errtype.IO, err, "/a").HasError(), ShouldBeTrue)

		So(errnew.Path.EmptyContent("/a").HasError(), ShouldBeTrue)

		So(errnew.Path.Marshal(nil, "/a"), ShouldBeNil)
		So(errnew.Path.Marshal(err, "/a").HasError(), ShouldBeTrue)
		So(errnew.Path.Unmarshal(nil, "/a"), ShouldBeNil)
		So(errnew.Path.Unmarshal(err, "/a").HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage4_Error_TypeFuncs(t *testing.T) {
	Convey("Error TypeFunc / TypeAnyFunctions / TypeAllFunctions", t, func() {
		errFn := func() error { return errors.New("e") }
		okFn := func() error { return nil }

		So(errnew.Error.TypeFunc(errtype.IO, nil), ShouldBeNil)
		So(errnew.Error.TypeFunc(errtype.IO, okFn), ShouldBeNil)
		So(errnew.Error.TypeFunc(errtype.IO, errFn).HasError(), ShouldBeTrue)
		So(errnew.Error.TypeFuncStackSkip(0, errtype.IO, errFn).HasError(), ShouldBeTrue)

		So(errnew.Error.TypeAnyFunctions(errtype.IO), ShouldBeNil)
		So(errnew.Error.TypeAnyFunctions(errtype.IO, okFn, okFn), ShouldBeNil)
		So(errnew.Error.TypeAnyFunctions(errtype.IO, okFn, errFn).HasError(), ShouldBeTrue)
		So(errnew.Error.TypeAnyFunctionsStackSkip(0, errtype.IO, errFn).HasError(), ShouldBeTrue)

		So(errnew.Error.TypeAllFunctions(errtype.IO), ShouldBeNil)
		So(errnew.Error.TypeAllFunctions(errtype.IO, okFn, okFn), ShouldBeNil)
		So(errnew.Error.TypeAllFunctions(errtype.IO, errFn, errFn).HasError(), ShouldBeTrue)
		So(errnew.Error.TypeAllFunctionsStackSkip(0, errtype.IO, errFn, errFn).HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage4_Ref_Remaining(t *testing.T) {
	Convey("Ref remaining creators", t, func() {
		err := errors.New("e")
		r1 := ref.New("k1", 1)
		r2 := ref.New("k2", 2)

		So(errnew.Ref.MsgWithMany(errtype.IO, "", r1), ShouldBeNil)
		So(errnew.Ref.MsgWithMany(errtype.IO, "msg", r1, r2).HasError(), ShouldBeTrue)

		So(errnew.Ref.ErrorWithOne(errtype.IO, nil, r1), ShouldBeNil)
		So(errnew.Ref.ErrorWithOne(errtype.IO, err, r1).HasError(), ShouldBeTrue)
		So(errnew.Ref.ErrorWithOneUsingStackSkip(0, errtype.IO, err, r1).HasError(), ShouldBeTrue)

		So(errnew.Ref.MsgWithOne(errtype.IO, "", r1), ShouldBeNil)
		So(errnew.Ref.MsgWithOne(errtype.IO, "msg", r1).HasError(), ShouldBeTrue)
		So(errnew.Ref.MsgWithOneUsingStackSkip(0, errtype.IO, "msg", r1).HasError(), ShouldBeTrue)

		So(errnew.Ref.ErrorWithMany(errtype.IO, nil, r1, r2), ShouldBeNil)
		So(errnew.Ref.ErrorWithMany(errtype.IO, err, r1, r2).HasError(), ShouldBeTrue)
		So(errnew.Ref.ErrorWithManyUsingStackSkip(0, errtype.IO, err, r1).HasError(), ShouldBeTrue)

		So(errnew.Ref.TwoWithMsg(errtype.IO, "msg", "a", 1, "b", 2).HasError(), ShouldBeTrue)
		So(errnew.Ref.TwoWithMsgUsingStackSkip(0, errtype.IO, "msg", "a", 1, "b", 2).HasError(), ShouldBeTrue)
		So(errnew.Ref.TwoWithError(errtype.IO, err, "a", 1, "b", 2).HasError(), ShouldBeTrue)
		So(errnew.Ref.TwoWithErrorUsingStackSkip(0, errtype.IO, err, "a", 1, "b", 2).HasError(), ShouldBeTrue)

		So(errnew.Ref.ManyWithMsg(errtype.IO, "msg", r1, r2).HasError(), ShouldBeTrue)
		So(errnew.Ref.ManyWithMsgUsingStackSkip(0, errtype.IO, "msg", r1).HasError(), ShouldBeTrue)
		So(errnew.Ref.Many(errtype.IO, r1, r2).HasError(), ShouldBeTrue)
		So(errnew.Ref.ManyUsingStackSkip(0, errtype.IO, r1).HasError(), ShouldBeTrue)

		So(errnew.Ref.ManyWithError(errtype.IO, nil, r1), ShouldBeNil)
		So(errnew.Ref.ManyWithError(errtype.IO, err, r1, r2).HasError(), ShouldBeTrue)
		So(errnew.Ref.ManyWithErrorUsingStackSkip(0, errtype.IO, err, r1).HasError(), ShouldBeTrue)

		wrap := errnew.Type.Error(errtype.IO, err)
		So(errnew.Ref.ManyUsingWrapper(nil, r1), ShouldBeNil)
		So(errnew.Ref.ManyUsingWrapper(wrap, r1).HasError(), ShouldBeTrue)
		So(errnew.Ref.ManyUsingWrapperUsingStackSkip(0, wrap, r1).HasError(), ShouldBeTrue)

		So(errnew.Ref.Message(errtype.IO, "k", 1, "msg").HasError(), ShouldBeTrue)
		So(errnew.Ref.MessageUsingStackSkip(0, errtype.IO, "k", 1, "msg").HasError(), ShouldBeTrue)

		So(errnew.Ref.Messages(errtype.IO, "k", 1, "m1", "m2").HasError(), ShouldBeTrue)
		So(errnew.Ref.MessagesUsingJoiner(errtype.IO, "k", 1, " | ", "m1", "m2").HasError(), ShouldBeTrue)
		So(errnew.Ref.MessagesUsingJoinerStackSkip(0, errtype.IO, "k", 1, " | ", "m1").HasError(), ShouldBeTrue)

		So(errnew.Ref.WithMessagesJoiner(errtype.IO, "k", 1, " | ", "m1", 2).HasError(), ShouldBeTrue)
		So(errnew.Ref.WithMessagesJoinerStackSkip(0, errtype.IO, "k", 1, " | ", "m1").HasError(), ShouldBeTrue)

		col := refs.New(2)
		col.Add("a", 1)
		So(errnew.Ref.CollectionWithMessagesJoiner(errtype.IO, col, " | ", "m1", 2).HasError(), ShouldBeTrue)
		So(errnew.Ref.CollectionWithMessagesJoinerUsingStackSkip(0, errtype.IO, col, " | ", "m1").HasError(), ShouldBeTrue)
	})
}
