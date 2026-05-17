package errconvtests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errconv"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

func TestErrconv_GetPtr(t *testing.T) {
	Convey("Given nil input", func() {
		result := errconv.GetPtr(nil)

		Convey("should return empty ResultPtr", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeFalse)
			So(result.Wrapper, ShouldBeNil)
		})
	})

	Convey("Given a *errorwrapper.Wrapper", func() {
		wrapper := errnew.Type.Default(errtype.Invalid)
		result := errconv.GetPtr(wrapper)

		Convey("should cast properly and return the wrapper", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeTrue)
			So(result.Wrapper, ShouldNotBeNil)
			So(result.Wrapper.HasError(), ShouldBeTrue)
		})
	})

	Convey("Given a nil *errorwrapper.Wrapper pointer", func() {
		var wrapper *errorwrapper.Wrapper
		result := errconv.GetPtr(wrapper)

		Convey("should return empty ResultPtr", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeFalse)
		})
	})

	Convey("Given an errorwrapper.Wrapper value", func() {
		wrapper := *errnew.Type.Default(errtype.NotFound)
		result := errconv.GetPtr(wrapper)

		Convey("should cast properly and return pointer to wrapper", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeTrue)
			So(result.Wrapper, ShouldNotBeNil)
			So(result.Wrapper.HasError(), ShouldBeTrue)
		})
	})

	Convey("Given a BasicErrWrapper interface value (Wrapper)", func() {
		var basicInf errcoreinf.BasicErrWrapper = errnew.Type.Default(errtype.Null)
		result := errconv.GetPtr(basicInf)

		Convey("should cast properly wrapping it in a new Wrapper", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeTrue)
			So(result.Wrapper, ShouldNotBeNil)
			So(result.Wrapper.HasError(), ShouldBeTrue)
		})
	})

	Convey("Given a BaseErrorOrCollectionWrapper interface value (Collection)", func() {
		collection := errwrappers.Empty().
			AddWrapperPtr(errnew.Type.Default(errtype.InvalidOption))
		var baseInf errcoreinf.BaseErrorOrCollectionWrapper = collection
		result := errconv.GetPtr(baseInf)

		Convey("should cast properly", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeTrue)
			So(result.Wrapper, ShouldNotBeNil)
			So(result.Wrapper.HasError(), ShouldBeTrue)
		})
	})

	Convey("Given an unrelated type", func() {
		result := errconv.GetPtr("just a string")

		Convey("should return empty ResultPtr", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeFalse)
		})
	})

	Convey("Given a random struct", func() {
		type myStruct struct{ X int }
		result := errconv.GetPtr(myStruct{X: 42})

		Convey("should return empty ResultPtr", func() {
			So(result, ShouldNotBeNil)
			So(result.IsCastedProperly, ShouldBeFalse)
		})
	})
}

func TestErrconv_Get(t *testing.T) {
	Convey("Given a wrapper value", func() {
		wrapper := errnew.Type.Default(errtype.PathMismatch)
		result := errconv.Get(wrapper)

		Convey("Get should return a properly casted Result", func() {
			So(result.IsCastedProperly, ShouldBeTrue)
			So(result.Wrapper, ShouldNotBeNil)
			So(result.Wrapper.HasError(), ShouldBeTrue)
		})
	})

	Convey("Given nil input", func() {
		result := errconv.Get(nil)

		Convey("Get should return failed cast Result", func() {
			So(result.IsCastedProperly, ShouldBeFalse)
			So(result.Wrapper, ShouldBeNil)
		})
	})

	Convey("Given an error value", func() {
		result := errconv.Get(errors.New("plain error"))

		Convey("should not cast properly (unrecognized type)", func() {
			So(result.IsCastedProperly, ShouldBeFalse)
		})
	})
}
