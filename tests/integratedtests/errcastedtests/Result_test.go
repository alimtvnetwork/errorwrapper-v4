package errcastedtests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errcasted"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrCasted_Result_Basics(t *testing.T) {
	Convey("FailedTypeCast", t, func() {
		r := errcasted.FailedTypeCast(42, "string", "cannot cast")
		So(r.IsCastedProperly, ShouldBeFalse)
		So(r.Wrapper.HasError(), ShouldBeTrue)
	})

	Convey("New with wrapper", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "ok")
		r := errcasted.New(w)
		So(r.IsCastedProperly, ShouldBeTrue)
		So(r.Wrapper.HasError(), ShouldBeTrue)
	})

	Convey("Empty", t, func() {
		r := errcasted.Empty()
		So(r.IsCastedProperly, ShouldBeFalse)
		So(r.Wrapper.IsEmpty(), ShouldBeTrue)
	})

	Convey("ToResultPtr", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "ok")
		r := errcasted.New(w)
		rp := r.ToResultPtr()
		So(rp, ShouldNotBeNil)
		So(rp.IsCastedProperly, ShouldBeTrue)
		So(rp.Wrapper.HasError(), ShouldBeTrue)
	})
}

func Test_ErrCasted_ResultPtr_Basics(t *testing.T) {
	Convey("FailedTypeCastPtr", t, func() {
		rp := errcasted.FailedTypeCastPtr(42, "string", "cannot cast")
		So(rp.IsCastedProperly, ShouldBeFalse)
		So(rp.Wrapper.HasError(), ShouldBeTrue)
	})

	Convey("EmptyPtr", t, func() {
		rp := errcasted.EmptyPtr()
		So(rp.IsCastedProperly, ShouldBeFalse)
		So(rp.Wrapper, ShouldBeNil)
	})

	Convey("NewPtr with wrapper", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "ok")
		rp := errcasted.NewPtr(w)
		So(rp.IsCastedProperly, ShouldBeTrue)
		So(rp.Wrapper.HasError(), ShouldBeTrue)
	})

	Convey("ToResult", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "ok")
		rp := errcasted.NewPtr(w)
		r := rp.ToResult()
		So(r.IsCastedProperly, ShouldBeTrue)
		So(r.Wrapper.HasError(), ShouldBeTrue)
	})
}
