package errfuncstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/enum-v10/linuxtype"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v3/errfuncs"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

// ---------- WrappersMap ----------

func Test_WrappersMap_BasicContracts_MC_MC(t *testing.T) {
	Convey("NewWrappersMap creates an empty map", t, func() {
		m := errfuncs.NewWrappersMap(2)
		So(m, ShouldNotBeNil)
		So(m.Length(), ShouldEqual, 0)
		So(m.Count(), ShouldEqual, 0)
		So(m.IsEmpty(), ShouldBeTrue)
		So(m.HasAnyItem(), ShouldBeFalse)
		So(m.LastIndex(), ShouldEqual, -1)
		So(m.HasIndex(0), ShouldBeFalse)
		So(m.Items(), ShouldNotBeNil)
	})

	Convey("Add appends non-nil and skips nil", t, func() {
		m := errfuncs.NewWrappersMap(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("e"))
		})

		m.Add("a", wf)
		So(m.Length(), ShouldEqual, 1)
		So(m.IsKeyExist("a"), ShouldBeTrue)
		So(m.IsKeyExist("missing"), ShouldBeFalse)
		So(m.Get("a"), ShouldNotBeNil)
		So(m.HasIndex(0), ShouldBeTrue)

		m.Add("nil-skipped", nil)
		So(m.Length(), ShouldEqual, 1)
	})

	Convey("IsAllKeysExist and GetItemsByKeys", t, func() {
		m := errfuncs.NewWrappersMap(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() })
		m.Add("a", wf)
		m.Add("b", wf)

		So(m.IsAllKeysExist("a", "b"), ShouldBeTrue)
		So(m.IsAllKeysExist("a", "missing"), ShouldBeFalse)

		items := m.GetItemsByKeys("a", "b", "missing")
		So(len(items), ShouldEqual, 2)

		empty := m.GetItemsByKeys()
		So(len(empty), ShouldEqual, 0)
	})

	Convey("ExecuteAll / String aggregate results", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("err", errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("boom"))
		}))

		coll := m.ExecuteAllByDefault()
		So(coll, ShouldNotBeNil)
		So(coll.HasAnyItem(), ShouldBeTrue)

		w := m.ExecuteAll()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)

		So(m.String(), ShouldNotBeBlank)
	})

	Convey("ExecuteAllUsingCollection on empty map returns input untouched", t, func() {
		m := errfuncs.NewWrappersMap(2)
		c := errwrappers.Empty()
		out := m.ExecuteAllUsingCollection(c)
		So(out, ShouldEqual, c)
		So(out.HasAnyItem(), ShouldBeFalse)
	})

	Convey("ExecuteByKeys reports missing-key errors", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("present", errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("p"))
		}))
		coll := m.ExecuteByKeys("present", "missing")
		So(coll, ShouldNotBeNil)
		So(coll.HasAnyItem(), ShouldBeTrue)
	})

	Convey("ExecuteByKeysAndCollection with nil collection allocates a new one", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("k", errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() }))
		coll := m.ExecuteByKeysAndCollection(nil, "k")
		So(coll, ShouldNotBeNil)
	})

	Convey("Linux helpers add and execute by linuxtype", t, func() {
		m := errfuncs.NewWrappersMap(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("linux"))
		})

		m.AddLinuxWrapperActions(linuxtype.UbuntuServer, wf)
		So(m.IsAllLinuxTypesExist(linuxtype.UbuntuServer), ShouldBeTrue)
		So(m.IsAllLinuxTypesExist(linuxtype.Centos7), ShouldBeFalse)
		So(m.IsAllLinuxTypesExist(nil...), ShouldBeTrue)

		coll := m.ExecuteByLinuxTypes(linuxtype.UbuntuServer)
		So(coll, ShouldNotBeNil)
		So(coll.HasAnyItem(), ShouldBeTrue)

		So(m.ExecuteByLinuxTypes(nil...).HasAnyItem(), ShouldBeFalse)
	})

	Convey("AddLinuxWrapperActionsIf respects guard flag", t, func() {
		m := errfuncs.NewWrappersMap(2)
		action := errfunc.LinuxWrapperAction{
			LinuxType:   linuxtype.UbuntuServer,
			WrapperFunc: func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() },
		}

		m.AddLinuxWrapperActionsIf(false, action)
		So(m.Length(), ShouldEqual, 0)

		m.AddLinuxWrapperActionsIf(true, action)
		So(m.Length(), ShouldEqual, 1)
	})

	Convey("AddLinuxTypeCmdOnce skips nil", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.AddLinuxTypeCmdOnce(linuxtype.UbuntuServer, nil)
		So(m.Length(), ShouldEqual, 0)
	})

	Convey("AsBasicSliceContractsBinder is non-nil", t, func() {
		m := errfuncs.NewWrappersMap(1)
		So(m.AsBasicSliceContractsBinder(), ShouldNotBeNil)
	})
}

// ---------- IsSuccessCollectorsMap ----------

func Test_IsSuccessCollectorsMap_BasicContracts_MC_MC(t *testing.T) {
	Convey("NewIsSuccessCollectorsMap creates an empty map", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		So(m, ShouldNotBeNil)
		So(m.Length(), ShouldEqual, 0)
		So(m.Count(), ShouldEqual, 0)
		So(m.IsEmpty(), ShouldBeTrue)
		So(m.HasAnyItem(), ShouldBeFalse)
		So(m.LastIndex(), ShouldEqual, -1)
		So(m.HasIndex(0), ShouldBeFalse)
		So(m.Items(), ShouldNotBeNil)
	})

	Convey("Add and lookup", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		f := errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("isc"))
			return false
		})

		m.Add("a", f)
		So(m.Length(), ShouldEqual, 1)
		So(m.IsKeyExist("a"), ShouldBeTrue)
		So(m.Get("a"), ShouldNotBeNil)

		m.Add("skipped", nil)
		So(m.Length(), ShouldEqual, 1)

		So(m.IsAllKeysExist("a"), ShouldBeTrue)
		So(m.IsAllKeysExist("a", "missing"), ShouldBeFalse)

		items := m.GetItemsByKeys("a", "missing")
		So(len(items), ShouldEqual, 1)
		So(len(m.GetItemsByKeys()), ShouldEqual, 0)
	})

	Convey("ExecuteAll / String aggregate", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		m.Add("a", errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("x"))
			return false
		}))

		coll := m.ExecuteAllByDefault()
		So(coll, ShouldNotBeNil)
		So(coll.HasAnyItem(), ShouldBeTrue)

		w := m.ExecuteAll()
		So(w, ShouldNotBeNil)

		So(m.String(), ShouldNotBeBlank)
	})

	Convey("ExecuteAllUsingCollection on empty map returns input untouched", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		c := errwrappers.Empty()
		out := m.ExecuteAllUsingCollection(c)
		So(out, ShouldEqual, c)
		So(out.HasAnyItem(), ShouldBeFalse)
	})

	Convey("Linux helpers", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		f := errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool { return true })
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("x"))
		})

		m.AddLinuxIsSuccessCollectorFunc(linuxtype.UbuntuServer, f)
		m.AddLinuxWrapperFunc(linuxtype.Centos7, wf)
		m.AddLinuxWrapperAction(errfunc.LinuxWrapperAction{
			LinuxType:   linuxtype.Android,
			WrapperFunc: wf,
		})

		So(m.IsAllLinuxTypesExist(linuxtype.UbuntuServer, linuxtype.Centos7, linuxtype.Android), ShouldBeTrue)
		So(m.IsAllLinuxTypesExist(nil...), ShouldBeTrue)

		coll := m.ExecuteByLinuxTypes(linuxtype.UbuntuServer)
		So(coll, ShouldNotBeNil)

		So(m.ExecuteByLinuxTypes(nil...).HasAnyItem(), ShouldBeFalse)
	})

	Convey("AddLinuxIsSuccessPlusErrorCollectActionsIf respects guard flag", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		action := errfunc.LinuxIsSuccessPlusErrorCollectAction{
			LinuxType:              linuxtype.UbuntuServer,
			IsSuccessCollectorFunc: func(ec *errwrappers.Collection) bool { return true },
		}

		m.AddLinuxIsSuccessPlusErrorCollectActionsIf(false, action)
		So(m.Length(), ShouldEqual, 0)

		m.AddLinuxIsSuccessPlusErrorCollectActionsIf(true, action)
		So(m.Length(), ShouldEqual, 1)
	})

	Convey("AddLinuxTypeCmdOnce skips nil", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		m.AddLinuxTypeCmdOnce(linuxtype.UbuntuServer, nil)
		So(m.Length(), ShouldEqual, 0)
	})

	Convey("AsBasicSliceContractsBinder is non-nil", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(1)
		So(m.AsBasicSliceContractsBinder(), ShouldNotBeNil)
	})
}
