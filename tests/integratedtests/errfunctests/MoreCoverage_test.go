package errfunctests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

func Test_MoreCoverage_EnumToNameSlice(t *testing.T) {
	Convey("EnumToNameSlice handles nil and items", t, func() {
		So(errfunc.EnumToNameSlice(), ShouldHaveLength, 0)
		So(errfunc.EnumToNameSlice(errtype.IO, errtype.NotFound), ShouldHaveLength, 2)
	})
}

func Test_MoreCoverage_LinuxTypesToNameSlice(t *testing.T) {
	Convey("LinuxTypesToNameSlice handles nil and items", t, func() {
		So(errfunc.LinuxTypesToNameSlice(), ShouldHaveLength, 0)
		got := errfunc.LinuxTypesToNameSlice(linuxtype.UbuntuServer, linuxtype.Centos7)
		So(got, ShouldHaveLength, 2)
		So(got[0], ShouldNotBeEmpty)
	})
}

func Test_MoreCoverage_ConvertLinuxErrorCollectorAction(t *testing.T) {
	Convey("LinuxErrorCollectorAction conversion", t, func() {
		action := errfunc.LinuxErrorCollectorAction{
			LinuxType: linuxtype.UbuntuServer,
			CollectorFunc: func(c *errwrappers.Collection) {
				c.AddTypeError(errtype.IO, errors.New("bad"))
			},
		}
		fn := errfunc.ConvertLinuxErrorCollectorActionToErrWrapperFunc(action)
		So(fn, ShouldNotBeNil)
		w := fn()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})

	Convey("LinuxErrorCollectorAction with no collected errors returns nil", t, func() {
		action := errfunc.LinuxErrorCollectorAction{
			LinuxType:     linuxtype.UbuntuServer,
			CollectorFunc: func(c *errwrappers.Collection) {},
		}
		fn := errfunc.ConvertLinuxErrorCollectorActionToErrWrapperFunc(action)
		So(fn(), ShouldBeNil)
	})
}

func Test_MoreCoverage_ConvertLinuxIsSuccessPlusErrorCollectAction(t *testing.T) {
	Convey("Success returns nil wrapper", t, func() {
		action := errfunc.LinuxIsSuccessPlusErrorCollectAction{
			LinuxType: linuxtype.Centos7,
			IsSuccessCollectorFunc: func(c *errwrappers.Collection) bool {
				return true
			},
		}
		fn := errfunc.ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc(action)
		So(fn(), ShouldBeNil)
	})

	Convey("Failure returns first wrapper", t, func() {
		action := errfunc.LinuxIsSuccessPlusErrorCollectAction{
			LinuxType: linuxtype.Centos7,
			IsSuccessCollectorFunc: func(c *errwrappers.Collection) bool {
				c.AddTypeError(errtype.IO, errors.New("fail"))
				return false
			},
		}
		fn := errfunc.ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc(action)
		w := fn()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})
}

func Test_MoreCoverage_MappedLinuxVersionWrapperFunctionExecutor(t *testing.T) {
	Convey("MappedLinuxVersionWrapperFunctionExecutor accessors", t, func() {
		successFn := errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return nil })
		failFn := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.IO, errors.New("boom"))
		})

		exec := &errfunc.MappedLinuxVersionWrapperFunctionExecutor{
			Mapping: map[linuxtype.Variant]errfunc.WrapperFunc{
				linuxtype.UbuntuServer: successFn,
				linuxtype.Centos7:      failFn,
			},
		}

		fn, errWrap := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(fn, ShouldNotBeNil)
		So(errWrap, ShouldBeNil)

		fn2, errWrap2 := exec.GetWrapperFunctionByType(linuxtype.Android)
		So(fn2, ShouldBeNil)
		So(errWrap2, ShouldNotBeNil)

		mapping, ec := exec.GetWrapperFunctionsByTypes(linuxtype.UbuntuServer, linuxtype.Android)
		So(mapping, ShouldNotBeNil)
		So(ec.HasError(), ShouldBeTrue)

		So(exec.ExecuteAllAvailableFunctions(), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_LinuxVersionWrapperFunctionExecutor(t *testing.T) {
	Convey("LinuxVersionWrapperFunctionExecutor accessors", t, func() {
		exec := &errfunc.LinuxVersionWrapperFunctionExecutor{
			UbuntuServer: func() *errorwrapper.Wrapper { return nil },
		}
		m := exec.GetFunctionsMappingLock()
		So(m, ShouldNotBeNil)
		So(exec.GetFunctionsMapping(), ShouldNotBeNil)

		fn, errWrap := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(fn, ShouldNotBeNil)
		So(errWrap, ShouldBeNil)

		_, errWrap2 := exec.GetWrapperFunctionByType(linuxtype.Android)
		So(errWrap2, ShouldNotBeNil)

		fns, ec := exec.GetWrapperFunctionsByTypes(linuxtype.UbuntuServer)
		So(fns, ShouldNotBeNil)
		So(ec, ShouldNotBeNil)

		So(exec.ExecuteAllAvailableFunctionsLock(), ShouldNotBeNil)
		So(exec.ExecuteAllAvailableFunctions(), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypesLock(linuxtype.UbuntuServer), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer), ShouldNotBeNil)
	})
}
