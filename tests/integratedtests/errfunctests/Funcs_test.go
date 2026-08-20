package errfunctests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

// Test_ConvertErrorFuncToWrapper verifies the simple-error-func -> WrapperFunc bridge.
func Test_ConvertErrorFuncToWrapper(t *testing.T) {
	Convey("nil SimpleErrorFunc returns nil WrapperFunc", t, func() {
		wf := errfunc.ConvertErrorFuncToWrapper(errtype.Generic, nil)
		So(wf, ShouldBeNil)
	})

	Convey("func returning nil error yields nil wrapper from resulting WrapperFunc", t, func() {
		wf := errfunc.ConvertErrorFuncToWrapper(errtype.Generic, func() error { return nil })
		So(wf, ShouldNotBeNil)
		So(wf(), ShouldBeNil)
	})

	Convey("func returning error yields a non-empty wrapper", t, func() {
		wf := errfunc.ConvertErrorFuncToWrapper(errtype.InvalidInput, func() error {
			return errors.New("input fail")
		})
		So(wf, ShouldNotBeNil)
		w := wf()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Message(), ShouldContainSubstring, "input fail")
	})
}

// Test_ConvertErrorFuncToIsSuccessCollectorFunc verifies the collector variant.
func Test_ConvertErrorFuncToIsSuccessCollectorFunc(t *testing.T) {
	Convey("nil SimpleErrorFunc returns nil collector", t, func() {
		cf := errfunc.ConvertErrorFuncToIsSuccessCollectorFunc(errtype.Generic, nil)
		So(cf, ShouldBeNil)
	})

	Convey("collector returns true and leaves collection empty on success", t, func() {
		cf := errfunc.ConvertErrorFuncToIsSuccessCollectorFunc(errtype.Generic, func() error { return nil })
		So(cf, ShouldNotBeNil)
		c := errwrappers.Empty()
		ok := cf(c)
		So(ok, ShouldBeTrue)
		So(c.HasAnyItem(), ShouldBeFalse)
	})

	Convey("collector returns false and populates collection on error", t, func() {
		cf := errfunc.ConvertErrorFuncToIsSuccessCollectorFunc(errtype.CommandExecution, func() error {
			return errors.New("cmd fail")
		})
		So(cf, ShouldNotBeNil)
		c := errwrappers.Empty()
		ok := cf(c)
		So(ok, ShouldBeFalse)
		So(c.HasAnyItem(), ShouldBeTrue)
	})
}

// Test_ConvertWrapperFuncToIsSuccessCollectorFunc verifies wrapper -> collector conversion.
func Test_ConvertWrapperFuncToIsSuccessCollectorFunc(t *testing.T) {
	Convey("nil WrapperFunc returns nil collector", t, func() {
		cf := errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(nil)
		So(cf, ShouldBeNil)
	})

	Convey("empty wrapper yields true and no collection changes", t, func() {
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errorwrapper.EmptyPtr()
		})
		cf := errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(wf)
		So(cf, ShouldNotBeNil)
		c := errwrappers.Empty()
		ok := cf(c)
		So(ok, ShouldBeTrue)
		So(c.HasAnyItem(), ShouldBeFalse)
	})

	Convey("non-empty wrapper yields false and collection changes", t, func() {
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.NotFound, errors.New("missing"))
		})
		cf := errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(wf)
		So(cf, ShouldNotBeNil)
		c := errwrappers.Empty()
		ok := cf(c)
		So(ok, ShouldBeFalse)
		So(c.HasAnyItem(), ShouldBeTrue)
	})
}
