package errfunctests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

func Test_MoreCoverage2_ConvertErrorFuncToWrapper(t *testing.T) {
	Convey("ConvertErrorFuncToWrapper", t, func() {
		So(errfunc.ConvertErrorFuncToWrapper(errtype.IO, nil), ShouldBeNil)

		nilFn := errfunc.ConvertErrorFuncToWrapper(errtype.IO, func() error { return nil })
		So(nilFn(), ShouldBeNil)

		errFn := errfunc.ConvertErrorFuncToWrapper(errtype.IO, func() error { return errors.New("boom") })
		w := errFn()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage2_ConvertErrorFuncToIsSuccessCollectorFunc(t *testing.T) {
	Convey("ConvertErrorFuncToIsSuccessCollectorFunc", t, func() {
		So(errfunc.ConvertErrorFuncToIsSuccessCollectorFunc(errtype.IO, nil), ShouldBeNil)

		okFn := errfunc.ConvertErrorFuncToIsSuccessCollectorFunc(
			errtype.IO,
			func() error { return nil })
		c := errwrappers.Empty()
		So(okFn(c), ShouldBeTrue)
		So(c.HasError(), ShouldBeFalse)

		failFn := errfunc.ConvertErrorFuncToIsSuccessCollectorFunc(
			errtype.IO,
			func() error { return errors.New("x") })
		c2 := errwrappers.Empty()
		So(failFn(c2), ShouldBeFalse)
		So(c2.HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage2_ConvertWrapperFuncToIsSuccessCollectorFunc(t *testing.T) {
	Convey("ConvertWrapperFuncToIsSuccessCollectorFunc", t, func() {
		So(errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(nil), ShouldBeNil)

		okFn := errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(
			errfunc.ConvertErrorFuncToWrapper(errtype.IO, func() error { return nil }))
		c := errwrappers.Empty()
		So(okFn(c), ShouldBeTrue)

		failFn := errfunc.ConvertWrapperFuncToIsSuccessCollectorFunc(
			errfunc.ConvertErrorFuncToWrapper(errtype.IO, func() error { return errors.New("y") }))
		c2 := errwrappers.Empty()
		So(failFn(c2), ShouldBeFalse)
		So(c2.HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage2_LinuxVersionIsSuccessCollectorFuncExecutor(t *testing.T) {
	Convey("LinuxVersionIsSuccessCollectorFuncExecutor accessors", t, func() {
		exec := &errfunc.LinuxVersionIsSuccessCollectorFuncExecutor{
			UbuntuServer: func(c *errwrappers.Collection) bool { return true },
			Centos7: func(c *errwrappers.Collection) bool {
				c.AddTypeError(errtype.IO, errors.New("c7"))
				return false
			},
		}

		m := exec.GetFunctionsMappingLock()
		So(m, ShouldNotBeNil)
		So(exec.GetFunctionsMapping(), ShouldNotBeNil)

		fn, ew := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(fn, ShouldNotBeNil)
		So(ew, ShouldBeNil)

		_, ew2 := exec.GetWrapperFunctionByType(linuxtype.Android)
		So(ew2, ShouldNotBeNil)

		mp, ec := exec.GetFunctionsByTypes(linuxtype.UbuntuServer, linuxtype.Android)
		So(mp, ShouldNotBeNil)
		So(ec.HasError(), ShouldBeTrue)

		So(exec.ExecuteAllAvailableFunctionsLock(), ShouldNotBeNil)
		So(exec.ExecuteAllAvailableFunctions(), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypesLock(linuxtype.UbuntuServer), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer, linuxtype.Centos7), ShouldNotBeNil)
		// Missing-type path: GetFunctionsByTypes surfaces an error → early return
		ecBad := exec.ExecuteByLinuxTypes(linuxtype.Android)
		So(ecBad.HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage2_MappedLinuxVersionIsSuccessCollectorFuncExecutor(t *testing.T) {
	Convey("MappedLinuxVersionIsSuccessCollectorFuncExecutor accessors", t, func() {
		exec := &errfunc.MappedLinuxVersionIsSuccessCollectorFuncExecutor{
			Mapping: map[linuxtype.Variant]errfunc.IsSuccessCollectorFunc{
				linuxtype.UbuntuServer: func(c *errwrappers.Collection) bool { return true },
				linuxtype.Centos7: func(c *errwrappers.Collection) bool {
					c.AddTypeError(errtype.IO, errors.New("c7"))
					return false
				},
			},
		}

		fn, ew := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(fn, ShouldNotBeNil)
		So(ew, ShouldBeNil)

		_, ew2 := exec.GetWrapperFunctionByType(linuxtype.Android)
		So(ew2, ShouldNotBeNil)

		mp, ec := exec.GetFunctionsByTypes(linuxtype.UbuntuServer, linuxtype.Android)
		So(mp, ShouldNotBeNil)
		So(ec.HasError(), ShouldBeTrue)

		So(exec.ExecuteAllAvailableFunctions(), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer, linuxtype.Centos7), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.Android).HasError(), ShouldBeTrue)
	})
}
