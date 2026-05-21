package errfunctests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

// failingWrapperFunc returns a non-empty wrapper.
func failingWrapperFunc() *errorwrapper.Wrapper {
	return errnew.Messages.Single(errtype.InvalidInput, "boom")
}

// successWrapperFunc returns empty wrapper.
func successWrapperFunc() *errorwrapper.Wrapper {
	return errorwrapper.EmptyPtr()
}

// failingCollector pushes one error and reports failure.
func failingCollector(c *errwrappers.Collection) bool {
	c.AddTypeError(errtype.InvalidInput, errors.New("collector-fail"))
	return false
}

// successCollector reports success.
func successCollector(c *errwrappers.Collection) bool { return true }

// Test_LinuxAction_Structs — exercises the small action structs.
func Test_LinuxAction_Structs(t *testing.T) {
	Convey("Action structs hold their LinuxType and func value", t, func() {
		a := errfunc.LinuxErrorAction{
			LinuxType:       linuxtype.UbuntuServer,
			SimpleErrorFunc: func() error { return errors.New("x") },
		}
		So(a.SimpleErrorFunc(), ShouldNotBeNil)
		So(a.LinuxType.Name(), ShouldNotBeBlank)

		b := errfunc.LinuxWrapperAction{
			LinuxType:   linuxtype.UbuntuServer,
			WrapperFunc: failingWrapperFunc,
		}
		So(b.WrapperFunc(), ShouldNotBeNil)

		c := errfunc.LinuxErrorCollectorAction{
			LinuxType:     linuxtype.UbuntuServer,
			CollectorFunc: func(_ *errwrappers.Collection) {},
		}
		c.CollectorFunc(errwrappers.Empty())

		d := errfunc.LinuxIsSuccessPlusErrorCollectAction{
			LinuxType:              linuxtype.UbuntuServer,
			IsSuccessCollectorFunc: successCollector,
		}
		So(d.IsSuccessCollectorFunc(errwrappers.Empty()), ShouldBeTrue)

		e := errfunc.LinuxIsSuccessPlusProcessErrorCollectAction{
			LinuxType: linuxtype.UbuntuServer,
		}
		So(e.LinuxType.Name(), ShouldNotBeBlank)
	})
}

// Test_LinuxTypesToNameSlice — converts variants to names.
func Test_LinuxTypesToNameSlice(t *testing.T) {
	Convey("Returns empty slice for nil input", t, func() {
		So(errfunc.LinuxTypesToNameSlice(), ShouldBeEmpty)
	})
	Convey("Returns matching names for inputs", t, func() {
		out := errfunc.LinuxTypesToNameSlice(linuxtype.UbuntuServer, linuxtype.Unknown)
		So(len(out), ShouldEqual, 2)
		So(out[0], ShouldNotBeBlank)
	})
}

// Test_ConvertLinuxErrorCollectorActionToErrWrapperFunc — runs the converter.
func Test_ConvertLinuxErrorCollectorActionToErrWrapperFunc(t *testing.T) {
	Convey("returns nil when collector adds no items", t, func() {
		action := errfunc.LinuxErrorCollectorAction{
			LinuxType:     linuxtype.UbuntuServer,
			CollectorFunc: func(_ *errwrappers.Collection) {},
		}
		fn := errfunc.ConvertLinuxErrorCollectorActionToErrWrapperFunc(action)
		So(fn(), ShouldBeNil)
	})

	Convey("returns first wrapper when collector adds an item", t, func() {
		action := errfunc.LinuxErrorCollectorAction{
			LinuxType: linuxtype.UbuntuServer,
			CollectorFunc: func(c *errwrappers.Collection) {
				c.AddTypeError(errtype.InvalidInput, errors.New("oops"))
			},
		}
		fn := errfunc.ConvertLinuxErrorCollectorActionToErrWrapperFunc(action)
		w := fn()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})
}

// Test_ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc — runs the converter.
func Test_ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc(t *testing.T) {
	Convey("returns nil on success", t, func() {
		action := errfunc.LinuxIsSuccessPlusErrorCollectAction{
			LinuxType:              linuxtype.UbuntuServer,
			IsSuccessCollectorFunc: successCollector,
		}
		fn := errfunc.ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc(action)
		So(fn(), ShouldBeNil)
	})

	Convey("returns first wrapper on failure", t, func() {
		action := errfunc.LinuxIsSuccessPlusErrorCollectAction{
			LinuxType:              linuxtype.UbuntuServer,
			IsSuccessCollectorFunc: failingCollector,
		}
		fn := errfunc.ConvertLinuxIsSuccessPlusErrorCollectActionToErrWrapperFunc(action)
		w := fn()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})
}

// Test_MappedLinuxVersionWrapperFunctionExecutor — lookup + execute.
func Test_MappedLinuxVersionWrapperFunctionExecutor(t *testing.T) {
	Convey("GetWrapperFunctionByType returns the registered func", t, func() {
		exec := &errfunc.MappedLinuxVersionWrapperFunctionExecutor{
			Mapping: map[linuxtype.Variant]errfunc.WrapperFunc{
				linuxtype.UbuntuServer: successWrapperFunc,
			},
		}
		fn, w := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(w, ShouldBeNil)
		So(fn, ShouldNotBeNil)
		So(fn(), ShouldNotBeNil) // empty wrapper, but non-nil ptr
	})

	Convey("GetWrapperFunctionByType returns error wrapper for missing key", t, func() {
		exec := &errfunc.MappedLinuxVersionWrapperFunctionExecutor{
			Mapping: map[linuxtype.Variant]errfunc.WrapperFunc{},
		}
		fn, w := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(fn, ShouldBeNil)
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})

	Convey("GetWrapperFunctionsByTypes splits hits and misses", t, func() {
		exec := &errfunc.MappedLinuxVersionWrapperFunctionExecutor{
			Mapping: map[linuxtype.Variant]errfunc.WrapperFunc{
				linuxtype.UbuntuServer: successWrapperFunc,
			},
		}
		mapping, errs := exec.GetWrapperFunctionsByTypes(
			linuxtype.UbuntuServer,
			linuxtype.Unknown,
		)
		So(mapping, ShouldNotBeNil)
		So(errs, ShouldNotBeNil)
		So(errs.HasAnyError(), ShouldBeTrue) // missing one
	})

	Convey("ExecuteAllAvailableFunctions runs every registered func", t, func() {
		exec := &errfunc.MappedLinuxVersionWrapperFunctionExecutor{
			Mapping: map[linuxtype.Variant]errfunc.WrapperFunc{
				linuxtype.UbuntuServer: failingWrapperFunc,
			},
		}
		out := exec.ExecuteAllAvailableFunctions()
		So(out, ShouldNotBeNil)
		So(out.HasAnyError(), ShouldBeTrue)
	})

	Convey("ExecuteByLinuxTypes runs the requested types", t, func() {
		exec := &errfunc.MappedLinuxVersionWrapperFunctionExecutor{
			Mapping: map[linuxtype.Variant]errfunc.WrapperFunc{
				linuxtype.UbuntuServer: failingWrapperFunc,
			},
		}
		out := exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer)
		So(out, ShouldNotBeNil)
		So(out.HasAnyError(), ShouldBeTrue)
	})
}

// Test_MappedLinuxVersionIsSuccessCollectorFuncExecutor — same shape for collector funcs.
func Test_MappedLinuxVersionIsSuccessCollectorFuncExecutor(t *testing.T) {
	Convey("GetWrapperFunctionByType hit + miss", t, func() {
		exec := &errfunc.MappedLinuxVersionIsSuccessCollectorFuncExecutor{
			Mapping: map[linuxtype.Variant]errfunc.IsSuccessCollectorFunc{
				linuxtype.UbuntuServer: successCollector,
			},
		}
		fn, w := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(w, ShouldBeNil)
		So(fn, ShouldNotBeNil)

		_, w2 := exec.GetWrapperFunctionByType(linuxtype.Unknown)
		So(w2, ShouldNotBeNil)
		So(w2.HasError(), ShouldBeTrue)
	})

	Convey("GetFunctionsByTypes returns mapping and collects misses", t, func() {
		exec := &errfunc.MappedLinuxVersionIsSuccessCollectorFuncExecutor{
			Mapping: map[linuxtype.Variant]errfunc.IsSuccessCollectorFunc{
				linuxtype.UbuntuServer: successCollector,
			},
		}
		mapping, errs := exec.GetFunctionsByTypes(linuxtype.UbuntuServer, linuxtype.Unknown)
		So(mapping, ShouldNotBeNil)
		So(errs.HasAnyError(), ShouldBeTrue)
	})

	Convey("ExecuteAllAvailableFunctions / ExecuteByLinuxTypes run", t, func() {
		exec := &errfunc.MappedLinuxVersionIsSuccessCollectorFuncExecutor{
			Mapping: map[linuxtype.Variant]errfunc.IsSuccessCollectorFunc{
				linuxtype.UbuntuServer: failingCollector,
			},
		}
		So(exec.ExecuteAllAvailableFunctions(), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer), ShouldNotBeNil)
	})
}

// Test_LinuxVersionWrapperFunctionExecutor — reflect-based executor.
func Test_LinuxVersionWrapperFunctionExecutor(t *testing.T) {
	Convey("GetFunctionsMapping picks up populated fields", t, func() {
		exec := &errfunc.LinuxVersionWrapperFunctionExecutor{
			UbuntuServer: failingWrapperFunc,
			Centos7:      successWrapperFunc,
		}
		mapping := exec.GetFunctionsMapping()
		So(mapping, ShouldNotBeNil)
		So(len(mapping), ShouldBeGreaterThanOrEqualTo, 2)

		// second call returns cached
		mapping2 := exec.GetFunctionsMappingLock()
		So(mapping2, ShouldNotBeNil)
	})

	Convey("GetWrapperFunctionByType hit + miss", t, func() {
		exec := &errfunc.LinuxVersionWrapperFunctionExecutor{
			UbuntuServer: failingWrapperFunc,
		}
		fn, w := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(w, ShouldBeNil)
		So(fn, ShouldNotBeNil)

		exec2 := &errfunc.LinuxVersionWrapperFunctionExecutor{}
		_, w2 := exec2.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(w2, ShouldNotBeNil)
		So(w2.HasError(), ShouldBeTrue)
	})

	Convey("GetWrapperFunctionsByTypes splits hits and misses", t, func() {
		exec := &errfunc.LinuxVersionWrapperFunctionExecutor{
			UbuntuServer: failingWrapperFunc,
		}
		mapping, errs := exec.GetWrapperFunctionsByTypes(
			linuxtype.UbuntuServer,
			linuxtype.Unknown,
		)
		So(mapping, ShouldNotBeNil)
		So(errs, ShouldNotBeNil)
	})

	Convey("ExecuteAllAvailableFunctions + ExecuteByLinuxTypes run", t, func() {
		exec := &errfunc.LinuxVersionWrapperFunctionExecutor{
			UbuntuServer: failingWrapperFunc,
		}
		So(exec.ExecuteAllAvailableFunctions(), ShouldNotBeNil)
		So(exec.ExecuteAllAvailableFunctionsLock(), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypesLock(linuxtype.UbuntuServer), ShouldNotBeNil)
	})
}

// Test_LinuxVersionIsSuccessCollectorFuncExecutor — reflect-based executor.
func Test_LinuxVersionIsSuccessCollectorFuncExecutor(t *testing.T) {
	Convey("GetFunctionsMapping picks up populated fields", t, func() {
		exec := &errfunc.LinuxVersionIsSuccessCollectorFuncExecutor{
			UbuntuServer: successCollector,
			Centos7:      failingCollector,
		}
		mapping := exec.GetFunctionsMapping()
		So(mapping, ShouldNotBeNil)
		So(len(mapping), ShouldBeGreaterThanOrEqualTo, 2)
		So(exec.GetFunctionsMappingLock(), ShouldNotBeNil)
	})

	Convey("GetWrapperFunctionByType hit + miss", t, func() {
		exec := &errfunc.LinuxVersionIsSuccessCollectorFuncExecutor{
			UbuntuServer: successCollector,
		}
		fn, w := exec.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(w, ShouldBeNil)
		So(fn, ShouldNotBeNil)

		empty := &errfunc.LinuxVersionIsSuccessCollectorFuncExecutor{}
		_, w2 := empty.GetWrapperFunctionByType(linuxtype.UbuntuServer)
		So(w2, ShouldNotBeNil)
	})

	Convey("GetFunctionsByTypes returns mapping and collects misses", t, func() {
		exec := &errfunc.LinuxVersionIsSuccessCollectorFuncExecutor{
			UbuntuServer: successCollector,
		}
		mapping, errs := exec.GetFunctionsByTypes(
			linuxtype.UbuntuServer,
			linuxtype.Unknown,
		)
		So(mapping, ShouldNotBeNil)
		So(errs, ShouldNotBeNil)
	})

	Convey("ExecuteAllAvailableFunctions + ExecuteByLinuxTypes run", t, func() {
		exec := &errfunc.LinuxVersionIsSuccessCollectorFuncExecutor{
			UbuntuServer: failingCollector,
		}
		So(exec.ExecuteAllAvailableFunctions(), ShouldNotBeNil)
		So(exec.ExecuteAllAvailableFunctionsLock(), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypes(linuxtype.UbuntuServer), ShouldNotBeNil)
		So(exec.ExecuteByLinuxTypesLock(linuxtype.UbuntuServer), ShouldNotBeNil)
	})
}
