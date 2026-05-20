package eithererrtests

import (
	"errors"
	"testing"

	
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/eithererr"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

// Test_Or verifies conditional wrapper selection.
func Test_Or(t *testing.T) {
	trueW := errnew.Type.Error(errtype.Generic, errors.New("true"))
	falseW := errnew.Type.Error(errtype.NotFound, errors.New("false"))

	Convey("Or returns trueWrapper when condition is true", t, func() {
		w := eithererr.Or(true, trueW, falseW)
		So(w, ShouldEqual, trueW)
	})

	Convey("Or returns falseWrapper when condition is false", t, func() {
		w := eithererr.Or(false, trueW, falseW)
		So(w, ShouldEqual, falseW)
	})
}

// Test_OrEmpty verifies conditional wrapper or nil.
func Test_OrEmpty(t *testing.T) {
	trueW := errnew.Type.Error(errtype.Generic, errors.New("true"))

	Convey("OrEmpty returns wrapper when condition is true", t, func() {
		w := eithererr.OrEmpty(true, trueW)
		So(w, ShouldEqual, trueW)
	})

	Convey("OrEmpty returns nil when condition is false", t, func() {
		w := eithererr.OrEmpty(false, trueW)
		So(w, ShouldBeNil)
	})
}

// Test_OrCollectionPtr verifies conditional collection selection.
func Test_OrCollectionPtr(t *testing.T) {
	trueC := errwrappers.Empty()
	trueC.AddTypeError(errtype.Generic, errors.New("true"))
	falseC := errwrappers.Empty()
	falseC.AddTypeError(errtype.NotFound, errors.New("false"))

	Convey("OrCollectionPtr returns true collection when condition is true", t, func() {
		c := eithererr.OrCollectionPtr(true, trueC, falseC)
		So(c, ShouldEqual, trueC)
	})

	Convey("OrCollectionPtr returns false collection when condition is false", t, func() {
		c := eithererr.OrCollectionPtr(false, trueC, falseC)
		So(c, ShouldEqual, falseC)
	})
}

// Test_OrEmptyCollectionPtr verifies conditional collection or empty.
func Test_OrEmptyCollectionPtr(t *testing.T) {
	trueC := errwrappers.Empty()
	trueC.AddTypeError(errtype.Generic, errors.New("true"))

	Convey("OrEmptyCollectionPtr returns collection when true", t, func() {
		c := eithererr.OrEmptyCollectionPtr(true, trueC)
		So(c, ShouldEqual, trueC)
	})

	Convey("OrEmptyCollectionPtr returns empty collection when false", t, func() {
		c := eithererr.OrEmptyCollectionPtr(false, trueC)
		So(c, ShouldNotBeNil)
		So(c.HasAnyError(), ShouldBeFalse)
	})
}

// Test_AnyFirstOrEmpty returns first wrapper with error.
func Test_AnyFirstOrEmpty(t *testing.T) {
	empty := errorwrapper.EmptyPtr()
	e1 := errnew.Type.Error(errtype.Generic, errors.New("e1"))
	e2 := errnew.Type.Error(errtype.NotFound, errors.New("e2"))

	Convey("empty args returns nil", t, func() {
		So(eithererr.AnyFirstOrEmpty(), ShouldBeNil)
	})

	Convey("all empty returns nil", t, func() {
		So(eithererr.AnyFirstOrEmpty(empty, empty), ShouldBeNil)
	})

	Convey("returns first non-empty wrapper", t, func() {
		So(eithererr.AnyFirstOrEmpty(empty, e1, e2), ShouldEqual, e1)
	})

	Convey("returns second if first is empty", t, func() {
		So(eithererr.AnyFirstOrEmpty(empty, e2), ShouldEqual, e2)
	})
}

// Test_AnyErrInfFirstOrEmpty returns first interface with error.
func Test_AnyErrInfFirstOrEmpty(t *testing.T) {
	emptyInf := errwrappers.Empty()
	collInf := errwrappers.Empty()
	collInf.AddTypeError(errtype.Generic, errors.New("coll"))
	w := errnew.Type.Error(errtype.NotFound, errors.New("w"))

	Convey("empty args returns nil", t, func() {
		So(eithererr.AnyErrInfFirstOrEmpty(), ShouldBeNil)
	})

	Convey("all empty returns nil", t, func() {
		So(eithererr.AnyErrInfFirstOrEmpty(emptyInf, emptyInf), ShouldBeNil)
	})

	Convey("returns first collection with error", t, func() {
		got := eithererr.AnyErrInfFirstOrEmpty(emptyInf, collInf)
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeTrue)
	})

	Convey("skips nil interfaces", t, func() {
		got := eithererr.AnyErrInfFirstOrEmpty(nil, collInf)
		So(got, ShouldNotBeNil)
	})
}

// Test_AnyBasicErrFirstOrEmpty returns first basic error with error.
func Test_AnyBasicErrFirstOrEmpty(t *testing.T) {
	emptyW := errorwrapper.EmptyPtr()
	w := errnew.Type.Error(errtype.Generic, errors.New("w"))

	Convey("empty args returns nil", t, func() {
		So(eithererr.AnyBasicErrFirstOrEmpty(), ShouldBeNil)
	})

	Convey("all empty returns nil", t, func() {
		So(eithererr.AnyBasicErrFirstOrEmpty(emptyW, emptyW), ShouldBeNil)
	})

	Convey("returns first non-empty basic error", t, func() {
		got := eithererr.AnyBasicErrFirstOrEmpty(emptyW, w)
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeTrue)
	})

	Convey("skips nil basic errors", t, func() {
		got := eithererr.AnyBasicErrFirstOrEmpty(nil, w)
		So(got, ShouldNotBeNil)
	})
}

// Test_ExecAnyFunc returns first function result with error.
func Test_ExecAnyFunc(t *testing.T) {
	Convey("empty functions returns nil", t, func() {
		So(eithererr.ExecAnyFunc(), ShouldBeNil)
	})

	Convey("all functions return empty => nil", t, func() {
		got := eithererr.ExecAnyFunc(
			func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() },
			func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() },
		)
		So(got, ShouldBeNil)
	})

	Convey("returns first non-empty wrapper", t, func() {
		got := eithererr.ExecAnyFunc(
			func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() },
			func() *errorwrapper.Wrapper { return errnew.Type.Error(errtype.Generic, errors.New("f2")) },
			func() *errorwrapper.Wrapper { return errnew.Type.Error(errtype.NotFound, errors.New("f3")) },
		)
		So(got, ShouldNotBeNil)
		So(got.Message(), ShouldContainSubstring, "f2")
	})
}

// Test_ExecAllFunc returns all wrappers with errors.
func Test_ExecAllFunc(t *testing.T) {
	Convey("empty functions returns nil", t, func() {
		So(eithererr.ExecAllFunc(), ShouldBeNil)
	})

	Convey("collects only non-empty wrappers", t, func() {
		got := eithererr.ExecAllFunc(
			func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() },
			func() *errorwrapper.Wrapper { return errnew.Type.Error(errtype.Generic, errors.New("e1")) },
			func() *errorwrapper.Wrapper { return errorwrapper.EmptyPtr() },
			func() *errorwrapper.Wrapper { return errnew.Type.Error(errtype.NotFound, errors.New("e2")) },
		)
		So(len(got), ShouldEqual, 2)
		So(got[0].Message(), ShouldContainSubstring, "e1")
		So(got[1].Message(), ShouldContainSubstring, "e2")
	})
}

// Test_WrapperOrFunc short-circuits or defers to fallback func.
func Test_WrapperOrFunc(t *testing.T) {
	Convey("returns wrapper when it has error", t, func() {
		w := errnew.Type.Error(errtype.Generic, errors.New("has err"))
		got := eithererr.WrapperOrFunc(w, func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.NotFound, errors.New("fallback"))
		})
		So(got, ShouldEqual, w)
	})

	Convey("returns wrapper even when empty (non-nil)", t, func() {
		w := errorwrapper.EmptyPtr()
		got := eithererr.WrapperOrFunc(w, func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.NotFound, errors.New("fallback"))
		})
		So(got, ShouldEqual, w)
	})

	Convey("returns fallback when wrapper is nil", t, func() {
		fallback := errnew.Type.Error(errtype.NotFound, errors.New("fallback"))
		got := eithererr.WrapperOrFunc(nil, func() *errorwrapper.Wrapper { return fallback })
		So(got, ShouldEqual, fallback)
	})
}
