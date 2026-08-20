package errnewtests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
	"github.com/alimtvnetwork/errorwrapper-v4/refs"
)

func Test_MoreCoverage5_Refs_Remaining(t *testing.T) {
	Convey("errnew.Refs uncovered methods", t, func() {
		err := errors.New("boom")
		rv := ref.New("k", "v")
		coll := refs.NewUsingValues(rv)

		// nil-guard paths
		So(errnew.Refs.Error(errtype.IO, nil, rv), ShouldBeNil)
		So(errnew.Refs.ErrorUsingStackSkip(0, errtype.IO, nil, rv), ShouldBeNil)
		So(errnew.Refs.ErrorWithMessage(errtype.IO, nil, "m", rv), ShouldBeNil)
		So(errnew.Refs.ErrorWithMessageUsingStackSkip(0, errtype.IO, nil, "m", rv), ShouldBeNil)
		So(errnew.Refs.Msg(errtype.IO, "", rv), ShouldBeNil)
		So(errnew.Refs.ErrorWithOne(errtype.IO, nil, rv), ShouldBeNil)
		So(errnew.Refs.ErrorWithOneUsingStackSkip(0, errtype.IO, nil, rv), ShouldBeNil)
		So(errnew.Refs.MsgWithOne(errtype.IO, "", rv), ShouldBeNil)
		So(errnew.Refs.MsgWithOneUsingStackSkip(0, errtype.IO, "", rv), ShouldBeNil)
		So(errnew.Refs.ErrorMessages(errtype.IO, nil, coll, "m"), ShouldBeNil)
		So(errnew.Refs.ErrorMessagesUsingStackSkip(0, errtype.IO, nil, coll, "m"), ShouldBeNil)
		So(errnew.Refs.OneErrorMessages(0, errtype.IO, nil, rv, "m"), ShouldBeNil)
		So(errnew.Refs.OneErrorMessagesUsingStackSkip(0, errtype.IO, nil, rv, "m"), ShouldBeNil)
		So(errnew.Refs.MergeWrapper(nil, rv), ShouldBeNil)
		So(errnew.Refs.MergeWrapperUsingStackSkip(0, nil, rv), ShouldBeNil)

		// populated paths
		So(errnew.Refs.Error(errtype.IO, err, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.ErrorUsingStackSkip(0, errtype.IO, err, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.ErrorWithMessage(errtype.IO, err, "m", rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.ErrorWithMessageUsingStackSkip(0, errtype.IO, err, "m", rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.Msg(errtype.IO, "m", rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.MsgUsingStackSkip(0, errtype.IO, "m", rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.Quick(errtype.IO, "m", "loc1", "loc2").HasError(), ShouldBeTrue)
		So(errnew.Refs.Type(errtype.IO, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.UsingStackSkip(0, errtype.IO, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.Many(errtype.IO, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.ManyUsingStackSkip(0, errtype.IO, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.ErrorWithOne(errtype.IO, err, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.ErrorWithOneUsingStackSkip(0, errtype.IO, err, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.MsgWithOne(errtype.IO, "m", rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.MsgWithOneUsingStackSkip(0, errtype.IO, "m", rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.OnlyOne(errtype.IO, "k", "v").HasError(), ShouldBeTrue)
		So(errnew.Refs.OnlyOneUsingStackSkip(0, errtype.IO, "k", "v").HasError(), ShouldBeTrue)

		base := errnew.Type.Error(errtype.IO, err)
		So(errnew.Refs.MergeWrapper(base, rv).HasError(), ShouldBeTrue)
		So(errnew.Refs.MergeWrapperUsingStackSkip(0, base, rv).HasError(), ShouldBeTrue)

		So(errnew.Refs.Messages(errtype.IO, coll, "m1", "m2").HasError(), ShouldBeTrue)
		So(errnew.Refs.ErrorMessages(errtype.IO, err, coll, "m").HasError(), ShouldBeTrue)
		So(errnew.Refs.ErrorMessagesUsingStackSkip(0, errtype.IO, err, coll, "m").HasError(), ShouldBeTrue)
		So(errnew.Refs.MessagesUsingStackSkip(0, errtype.IO, coll, "m").HasError(), ShouldBeTrue)
		So(errnew.Refs.OneErrorMessages(0, errtype.IO, err, rv, "m").HasError(), ShouldBeTrue)
		So(errnew.Refs.OneErrorMessagesUsingStackSkip(0, errtype.IO, err, rv, "m").HasError(), ShouldBeTrue)
		So(errnew.Refs.MessagesUsingJoiner(errtype.IO, coll, ",", "a", "b").HasError(), ShouldBeTrue)
		So(errnew.Refs.MessagesUsingJoinerStackSkip(0, errtype.IO, coll, ",", "a", "b").HasError(), ShouldBeTrue)
		So(errnew.Refs.TypeQuick(errtype.IO, "k", "v").HasError(), ShouldBeTrue)
		So(errnew.Refs.TypeQuickUsingStackSkip(0, errtype.IO, "k", "v").HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage5_Null(t *testing.T) {
	Convey("errnew.Null variants", t, func() {
		err := errors.New("e")
		var nilPtr *int

		// nil object => Null error wrapper produced
		So(errnew.Null.OrWrapper(nil, nilPtr).HasError(), ShouldBeTrue)
		// existing error wrapper passthrough
		existing := errnew.Type.Error(errtype.IO, err)
		So(errnew.Null.OrWrapper(existing, nilPtr).HasError(), ShouldBeTrue)

		So(errnew.Null.OrError(errtype.IO, nil, nilPtr).HasError(), ShouldBeTrue)
		So(errnew.Null.OrError(errtype.IO, err, nilPtr).HasError(), ShouldBeTrue)

		So(errnew.Null.OrErrorFunc(nilPtr, errtype.IO, nil), ShouldBeNil)
		So(errnew.Null.OrErrorFunc(nilPtr, errtype.IO, func() error { return nil }), ShouldBeNil)
		So(errnew.Null.OrErrorFunc(nilPtr, errtype.IO, func() error { return err }).HasError(), ShouldBeTrue)

		So(errnew.Null.OrErrorWrapperFunc(nilPtr, nil).HasError(), ShouldBeTrue)
		So(errnew.Null.OrErrorWrapperFunc(nilPtr, func() *errorwrapper.Wrapper { return nil }), ShouldBeNil)

		So(errnew.Null.OrWrapperOrError(nilPtr, existing, errtype.IO, nil).HasError(), ShouldBeTrue)
		So(errnew.Null.OrWrapperOrError(nilPtr, nil, errtype.IO, err).HasError(), ShouldBeTrue)
		So(errnew.Null.OrWrapperOrError(nilPtr, nil, errtype.IO, nil).HasError(), ShouldBeTrue)

		So(errnew.Null.WithError(errtype.IO, nil, nilPtr).HasError(), ShouldBeTrue)
		So(errnew.Null.WithError(errtype.IO, err, nilPtr).HasError(), ShouldBeTrue)

		So(errnew.Null.Simple(nilPtr).HasError(), ShouldBeTrue)
		So(errnew.Null.ManyWithMessage("msg", nilPtr, "ok").HasError(), ShouldBeTrue)
		So(errnew.Null.ManyWithMessage("msg"), ShouldBeNil)

		So(errnew.Null.ManyByChecking(), ShouldBeNil)
		So(errnew.Null.ManyByChecking("ok"), ShouldBeNil)
		So(errnew.Null.ManyByChecking(nilPtr).HasError(), ShouldBeTrue)
		So(errnew.Null.ManyByCheckingUsingStackSkip(0), ShouldBeNil)
		So(errnew.Null.ManyByCheckingUsingStackSkip(0, nilPtr).HasError(), ShouldBeTrue)

		So(errnew.Null.WithRefs("m", nilPtr, ref.New("k", "v")).HasError(), ShouldBeTrue)

		So(errnew.Null.Message(""), ShouldBeNil)
		So(errnew.Null.Message("m").HasError(), ShouldBeTrue)
		So(errnew.Null.Error(nil), ShouldBeNil)
		So(errnew.Null.Error(err).HasError(), ShouldBeTrue)
		So(errnew.Null.ErrorWithMessage("m", nil), ShouldBeNil)
		So(errnew.Null.ErrorWithMessage("m", err).HasError(), ShouldBeTrue)
		So(errnew.Null.WithMessage("m", nilPtr).HasError(), ShouldBeTrue)
		So(errnew.Null.UsingStackSkip(0, "m", nilPtr).HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage5_FromTo(t *testing.T) {
	Convey("errnew.FromTo variants", t, func() {
		So(errnew.FromTo.Message(errtype.IO, "m", "a", "b").HasError(), ShouldBeTrue)
		So(errnew.FromTo.MessageStackSkip(0, errtype.IO, "m", "a", "b").HasError(), ShouldBeTrue)
		So(errnew.FromTo.Create(errtype.IO, "a", "b").HasError(), ShouldBeTrue)
		So(errnew.FromTo.CreateUsingStackSkip(0, errtype.IO, "a", "b").HasError(), ShouldBeTrue)

		So(errnew.FromTo.Messages(errtype.IO, false, "a", "b", "m1").HasError(), ShouldBeTrue)
		So(errnew.FromTo.Messages(errtype.IO, true, "a", "b", "m1").HasError(), ShouldBeTrue)
		So(errnew.FromTo.MessagesUsingStackSkip(0, errtype.IO, false, "a", "b", "m1").HasError(), ShouldBeTrue)
		So(errnew.FromTo.MessagesUsingStackSkip(0, errtype.IO, true, "a", "b", "m1").HasError(), ShouldBeTrue)

		So(errnew.FromTo.WithMetaMessages(errtype.IO, false, "a", "am", "b", "bm", "m").HasError(), ShouldBeTrue)
		So(errnew.FromTo.WithMetaMessages(errtype.IO, true, "a", "am", "b", "bm", "m").HasError(), ShouldBeTrue)
		So(errnew.FromTo.WithMetaMessagesUsingStackSkip(0, errtype.IO, false, "a", "am", "b", "bm", "m").HasError(), ShouldBeTrue)
		So(errnew.FromTo.WithMetaMessagesUsingStackSkip(0, errtype.IO, true, "a", "am", "b", "bm", "m").HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage5_Merge(t *testing.T) {
	Convey("errnew.Merge variants", t, func() {
		err := errors.New("e")
		a := errnew.Type.Error(errtype.IO, err)
		b := errnew.Type.Error(errtype.IO, errors.New("e2"))
		rv := ref.New("k", "v")

		So(errnew.Merge.New(nil, nil), ShouldBeNil)
		So(errnew.Merge.New(a, nil).HasError(), ShouldBeTrue)
		So(errnew.Merge.New(nil, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.New(a, b).HasError(), ShouldBeTrue)

		So(errnew.Merge.UsingNewType(errtype.IO, nil, nil), ShouldBeNil)
		So(errnew.Merge.UsingNewType(errtype.IO, a, nil).HasError(), ShouldBeTrue)
		So(errnew.Merge.UsingNewType(errtype.IO, nil, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.UsingNewType(errtype.IO, a, b).HasError(), ShouldBeTrue)

		So(errnew.Merge.UsingStackSkip(0, nil, nil), ShouldBeNil)
		So(errnew.Merge.UsingStackSkip(0, a, nil).HasError(), ShouldBeTrue)
		So(errnew.Merge.UsingStackSkip(0, nil, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.UsingStackSkip(0, a, b).HasError(), ShouldBeTrue)

		So(errnew.Merge.Three(nil, nil, nil), ShouldBeNil)
		So(errnew.Merge.Three(a, b, nil).HasError(), ShouldBeTrue)

		So(errnew.Merge.TwoWithRefs(errtype.IO, nil, nil), ShouldBeNil)
		So(errnew.Merge.TwoWithRefs(errtype.IO, a, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.TwoWithRefs(errtype.IO, a, b, rv).HasError(), ShouldBeTrue)
		So(errnew.Merge.TwoWithRefsUsingStackSkip(0, errtype.IO, nil, nil), ShouldBeNil)
		So(errnew.Merge.TwoWithRefsUsingStackSkip(0, errtype.IO, a, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.TwoWithRefsUsingStackSkip(0, errtype.IO, a, b, rv).HasError(), ShouldBeTrue)

		So(errnew.Merge.AddRefs(nil, rv), ShouldBeNil)
		So(errnew.Merge.AddRefs(a, rv).HasError(), ShouldBeTrue)
		So(errnew.Merge.AddRefsUsingStackSkip(0, nil, rv), ShouldBeNil)
		So(errnew.Merge.AddRefsUsingStackSkip(0, a, rv).HasError(), ShouldBeTrue)

		So(errnew.Merge.Many(errtype.IO), ShouldBeNil)
		So(errnew.Merge.Many(errtype.IO, a, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.ManyStackSkip(0, errtype.IO), ShouldBeNil)
		So(errnew.Merge.ManyStackSkip(0, errtype.IO, a, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.ManyUsingJoinerStackSkip(0, errtype.IO, ","), ShouldBeNil)
		So(errnew.Merge.ManyUsingJoinerStackSkip(0, errtype.IO, ",", a, b).HasError(), ShouldBeTrue)
		So(errnew.Merge.ManyAdditionalRefsUsingJoinerStackSkip(0, errtype.IO, ",", refs.NewUsingValues(rv)), ShouldBeNil)
		So(errnew.Merge.ManyAdditionalRefsUsingJoinerStackSkip(0, errtype.IO, ",", refs.NewUsingValues(rv), a, b).HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage5_DeserializeTo(t *testing.T) {
	Convey("errnew.DeserializeTo variants", t, func() {
		// nil/empty inputs => nil
		So(errnew.DeserializeTo.JsonResultToAnySkipOnNull(nil, nil), ShouldBeNil)
		convWp, parsedWp := errnew.DeserializeTo.BytesToWrapper(nil)
		So(convWp, ShouldBeNil)
		So(parsedWp, ShouldBeNil)
		So(errnew.DeserializeTo.BytesToUnmarshal(nil, nil), ShouldBeNil)
		So(errnew.DeserializeTo.BytesToAnyPtr(nil, nil), ShouldBeNil)

		convWp, parsedWp = errnew.DeserializeTo.JsonErrToWrapper(nil)
		So(convWp, ShouldBeNil)
		So(parsedWp, ShouldBeNil)
		convWp, parsedWp = errnew.DeserializeTo.JsonResultErrToWrapper(nil)
		So(convWp, ShouldBeNil)
		So(parsedWp, ShouldBeNil)
		convWp, parsedWp = errnew.DeserializeTo.JsonResultToWrapper(nil)
		So(convWp, ShouldBeNil)
		So(parsedWp, ShouldBeNil)

		// failing paths with junk bytes
		junk := []byte("not-json")
		var target int
		So(errnew.DeserializeTo.BytesToUnmarshal(junk, &target).HasError(), ShouldBeTrue)
		So(errnew.DeserializeTo.BytesToAnyPtr(junk, &target).HasError(), ShouldBeTrue)

		convWp, parsedWp = errnew.DeserializeTo.BytesToWrapper(junk)
		So(convWp, ShouldNotBeNil)
		So(parsedWp.HasError(), ShouldBeTrue)
		convWp, parsedWp = errnew.DeserializeTo.BytesToWrapperUsingStackSkip(0, junk)
		So(parsedWp.HasError(), ShouldBeTrue)
		_ = convWp

		convWp, parsedWp = errnew.DeserializeTo.JsonStringToWrapper(false, "")
		So(convWp, ShouldBeNil)
		So(parsedWp.HasError(), ShouldBeTrue)
		convWp, parsedWp = errnew.DeserializeTo.JsonStringToWrapper(true, "")
		So(convWp, ShouldBeNil)
		So(parsedWp, ShouldBeNil)
		convWp, parsedWp = errnew.DeserializeTo.JsonStringToWrapper(false, "not-json")
		So(parsedWp.HasError(), ShouldBeTrue)
		_ = convWp

		convWp, parsedWp = errnew.DeserializeTo.JsonErrToWrapper(errors.New("not-json"))
		So(parsedWp.HasError(), ShouldBeTrue)
		_ = convWp
		convWp, parsedWp = errnew.DeserializeTo.JsonResultErrToWrapper(errors.New("not-json"))
		So(parsedWp.HasError(), ShouldBeTrue)
		_ = convWp

		// JsonResultToAnyOnErrAddMsg with nil json result => deserialize fails
		So(errnew.DeserializeTo.JsonResultToAnyOnErrAddMsg("ctx: ", nil, &target).HasError(), ShouldBeTrue)
		// JsonResultToAnyOption both branches with nil
		So(errnew.DeserializeTo.JsonResultToAnyOption(true, nil, &target), ShouldBeNil)
		So(errnew.DeserializeTo.JsonResultToAnyOption(false, nil, &target).HasError(), ShouldBeTrue)
		// JsonResultToAny with nil json result => fails
		So(errnew.DeserializeTo.JsonResultToAny(nil, &target).HasError(), ShouldBeTrue)
	})
}
