package errfuncstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v3/errfuncs"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

// ---------- Wrappers ----------

func Test_Wrappers_BasicContracts(t *testing.T) {
	Convey("NewWrappers creates an empty slice container", t, func() {
		w := errfuncs.NewWrappers(2)
		So(w, ShouldNotBeNil)
		So(w.Length(), ShouldEqual, 0)
		So(w.IsEmpty(), ShouldBeTrue)
		So(w.HasAnyItem(), ShouldBeFalse)
	})

	Convey("Add appends non-nil wrapper funcs", t, func() {
		w := errfuncs.NewWrappers(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("a"))
		})
		w.Add(wf)
		So(w.Length(), ShouldEqual, 1)
		So(w.HasAnyItem(), ShouldBeTrue)
		So(w.HasIndex(0), ShouldBeTrue)
		So(w.HasIndex(1), ShouldBeFalse)
	})

	Convey("Add skips nil wrapper func", t, func() {
		w := errfuncs.NewWrappers(2)
		w.Add(nil)
		So(w.Length(), ShouldEqual, 0)
	})

	Convey("GetAt returns the item at index", t, func() {
		w := errfuncs.NewWrappers(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() })
		w.Add(wf)
		So(w.GetAt(0), ShouldEqual, wf)
	})

	Convey("GetSafeAt returns nil for out-of-bounds", t, func() {
		w := errfuncs.NewWrappers(2)
		So(w.GetSafeAt(0), ShouldBeNil)
	})

	Convey("AddsIf adds only when true", t, func() {
		w := errfuncs.NewWrappers(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() })
		w.AddsIf(false, wf)
		So(w.Length(), ShouldEqual, 0)
		w.AddsIf(true, wf)
		So(w.Length(), ShouldEqual, 1)
	})

	Convey("SimpleErrorFunctionsAdds converts and appends", t, func() {
		w := errfuncs.NewWrappers(2)
		w.SimpleErrorFunctionsAdds(errtype.InvalidInput,
			func() error { return errors.New("simple") },
		)
		So(w.Length(), ShouldEqual, 1)
		coll := w.ExecuteAllCollection()
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteAll aggregates wrappers into a single Wrapper", t, func() {
		w := errfuncs.NewWrappers(2)
		w.Add(errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("e1"))
		}))
		w.Add(errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.NotFound, errors.New("e2"))
		}))
		result := w.ExecuteAll()
		So(result, ShouldNotBeNil)
		So(result.HasError(), ShouldBeTrue)
		msg := result.Message()
		So(msg, ShouldContainSubstring, "e1")
		So(msg, ShouldContainSubstring, "e2")
	})

	Convey("String serializes accumulated errors", t, func() {
		w := errfuncs.NewWrappers(2)
		w.Add(errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("x"))
		}))
		So(w.String(), ShouldNotBeBlank)
	})
}

// ---------- Collectors ----------

func Test_Collectors_BasicContracts(t *testing.T) {
	Convey("NewCollectors creates an empty collector slice", t, func() {
		c := errfuncs.NewCollectors(2)
		So(c, ShouldNotBeNil)
		So(c.Length(), ShouldEqual, 0)
		So(c.IsEmpty(), ShouldBeTrue)
	})

	Convey("Add appends non-nil collector funcs", t, func() {
		c := errfuncs.NewCollectors(2)
		cf := errfunc.CollectorFunc(func(ec *errwrappers.Collection) {
			ec.AddTypeError(errtype.Generic, errors.New("collected"))
		})
		c.Add(cf)
		So(c.Length(), ShouldEqual, 1)
	})

	Convey("Add skips nil collector func", t, func() {
		c := errfuncs.NewCollectors(2)
		c.Add(nil)
		So(c.Length(), ShouldEqual, 0)
	})

	Convey("ExecuteAllCollection mutates the passed collection", t, func() {
		c := errfuncs.NewCollectors(2)
		c.Add(errfunc.CollectorFunc(func(ec *errwrappers.Collection) {
			ec.AddTypeError(errtype.Generic, errors.New("one"))
		}))
		coll := errwrappers.Empty()
		c.ExecuteAllCollection(coll)
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteAllCollectionWithNewEmpty returns a new populated collection", t, func() {
		c := errfuncs.NewCollectors(2)
		c.Add(errfunc.CollectorFunc(func(ec *errwrappers.Collection) {
			ec.AddTypeError(errtype.InvalidInput, errors.New("two"))
		}))
		coll := c.ExecuteAllCollectionWithNewEmpty()
		So(coll, ShouldNotBeNil)
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteAll returns aggregated wrapper", t, func() {
		c := errfuncs.NewCollectors(2)
		c.Add(errfunc.CollectorFunc(func(ec *errwrappers.Collection) {
			ec.AddTypeError(errtype.Generic, errors.New("agg"))
		}))
		w := c.ExecuteAll()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})

	Convey("String serializes collector output", t, func() {
		c := errfuncs.NewCollectors(2)
		c.Add(errfunc.CollectorFunc(func(ec *errwrappers.Collection) {
			ec.AddTypeError(errtype.Generic, errors.New("s"))
		}))
		So(c.String(), ShouldNotBeBlank)
	})
}

// ---------- IsSuccessCollectors ----------

func Test_IsSuccessCollectors_BasicContracts(t *testing.T) {
	Convey("NewIsSuccessCollectors creates an empty slice", t, func() {
		isc := errfuncs.NewIsSuccessCollectors(2)
		So(isc, ShouldNotBeNil)
		So(isc.Length(), ShouldEqual, 0)
		So(isc.IsEmpty(), ShouldBeTrue)
	})

	Convey("Add appends non-nil funcs", t, func() {
		isc := errfuncs.NewIsSuccessCollectors(2)
		f := errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool { return true })
		isc.Add(f)
		So(isc.Length(), ShouldEqual, 1)
	})

	Convey("Add skips nil funcs", t, func() {
		isc := errfuncs.NewIsSuccessCollectors(2)
		isc.Add(nil)
		So(isc.Length(), ShouldEqual, 0)
	})

	Convey("SimpleErrorFunctionsAdds converts and appends", t, func() {
		isc := errfuncs.NewIsSuccessCollectors(2)
		isc.SimpleErrorFunctionsAdds(errtype.CommandExecution,
			func() error { return errors.New("fail") },
		)
		So(isc.Length(), ShouldEqual, 1)
		coll := isc.ExecuteAllCollection()
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteAllCollection aggregates into collection", t, func() {
		isc := errfuncs.NewIsSuccessCollectors(2)
		isc.Add(errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("isc"))
			return false
		}))
		coll := isc.ExecuteAllCollection()
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteAll returns aggregated wrapper", t, func() {
		isc := errfuncs.NewIsSuccessCollectors(2)
		isc.Add(errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("w"))
			return false
		}))
		w := isc.ExecuteAll()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})

	Convey("String serializes output", t, func() {
		isc := errfuncs.NewIsSuccessCollectors(2)
		isc.Add(errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("str"))
			return false
		}))
		So(isc.String(), ShouldNotBeBlank)
	})
}

// ---------- WrappersMap ----------

func Test_WrappersMap_BasicContracts(t *testing.T) {
	Convey("NewWrappersMap creates an empty map", t, func() {
		m := errfuncs.NewWrappersMap(2)
		So(m, ShouldNotBeNil)
		So(m.Length(), ShouldEqual, 0)
		So(m.IsEmpty(), ShouldBeTrue)
	})

	Convey("Add and Get round-trip", t, func() {
		m := errfuncs.NewWrappersMap(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() })
		m.Add("k1", wf)
		So(m.Length(), ShouldEqual, 1)
		So(m.IsKeyExist("k1"), ShouldBeTrue)
		So(m.Get("k1"), ShouldEqual, wf)
		So(m.Get("missing"), ShouldBeNil)
	})

	Convey("Add skips nil wrapper func", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("k1", nil)
		So(m.Length(), ShouldEqual, 0)
	})

	Convey("IsAllKeysExist validates multiple keys", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("a", errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() }))
		m.Add("b", errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() }))
		So(m.IsAllKeysExist("a", "b"), ShouldBeTrue)
		So(m.IsAllKeysExist("a", "c"), ShouldBeFalse)
	})

	Convey("GetItemsByKeys returns matching funcs", t, func() {
		m := errfuncs.NewWrappersMap(2)
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() })
		m.Add("a", wf)
		items := m.GetItemsByKeys("a", "missing")
		So(len(items), ShouldEqual, 1)
	})

	Convey("ExecuteByKeys runs selected keys", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("ok", errfunc.WrapperFunc(func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() }))
		m.Add("bad", errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("bad"))
		}))
		coll := m.ExecuteByKeys("ok", "bad")
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteByKeys reports missing keys", t, func() {
		m := errfuncs.NewWrappersMap(2)
		coll := m.ExecuteByKeys("missing")
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteAll aggregates every entry", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("e1", errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("e1"))
		}))
		m.Add("e2", errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.NotFound, errors.New("e2"))
		}))
		w := m.ExecuteAll()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})

	Convey("String serializes map output", t, func() {
		m := errfuncs.NewWrappersMap(2)
		m.Add("x", errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("x"))
		}))
		So(m.String(), ShouldNotBeBlank)
	})
}

// ---------- IsSuccessCollectorsMap ----------

func Test_IsSuccessCollectorsMap_BasicContracts(t *testing.T) {
	Convey("NewIsSuccessCollectorsMap creates an empty map", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		So(m, ShouldNotBeNil)
		So(m.Length(), ShouldEqual, 0)
		So(m.IsEmpty(), ShouldBeTrue)
	})

	Convey("Add and Get round-trip", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		f := errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool { return true })
		m.Add("k1", f)
		So(m.Length(), ShouldEqual, 1)
		So(m.IsKeyExist("k1"), ShouldBeTrue)
		So(m.Get("k1"), ShouldEqual, f)
	})

	Convey("Add skips nil func", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		m.Add("k1", nil)
		So(m.Length(), ShouldEqual, 0)
	})

	Convey("IsAllKeysExist validates keys", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		f := errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool { return true })
		m.Add("a", f).Add("b", f)
		So(m.IsAllKeysExist("a", "b"), ShouldBeTrue)
		So(m.IsAllKeysExist("a", "z"), ShouldBeFalse)
	})

	Convey("ExecuteByKeys runs selected keys", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		m.Add("good", errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool { return true }))
		m.Add("bad", errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("bad"))
			return false
		}))
		coll := m.ExecuteByKeys("good", "bad")
		So(coll.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ExecuteAll aggregates every entry", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		m.Add("e1", errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("e1"))
			return false
		}))
		m.Add("e2", errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.NotFound, errors.New("e2"))
			return false
		}))
		w := m.ExecuteAll()
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})

	Convey("String serializes map output", t, func() {
		m := errfuncs.NewIsSuccessCollectorsMap(2)
		m.Add("x", errfunc.IsSuccessCollectorFunc(func(ec *errwrappers.Collection) bool {
			ec.AddTypeError(errtype.Generic, errors.New("x"))
			return false
		}))
		So(m.String(), ShouldNotBeBlank)
	})
}
