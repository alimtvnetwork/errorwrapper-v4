package errwoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwonce"
	. "github.com/smartystreets/goconvey/convey"
)

func TestErrwonce_Instance(t *testing.T) {
	Convey("Given an Instance with a lazy initializer that returns an error wrapper", t, func() {
		initCallCount := 0
		initFunc := func() *errorwrapper.Wrapper {
			initCallCount++
			return errnew.Type.Default(errtype.Invalid)
		}
		instance := errwonce.New(initFunc)

		Convey("Value should trigger initialization exactly once", func() {
			So(instance.Value(), ShouldNotBeNil)
			So(initCallCount, ShouldEqual, 1)
			// second call should not increment
			_ = instance.Value()
			So(initCallCount, ShouldEqual, 1)
		})

		Convey("HasError should be true and IsSuccess should be false", func() {
			So(instance.HasError(), ShouldBeTrue)
			So(instance.IsSuccess(), ShouldBeFalse)
			So(instance.IsFailed(), ShouldBeTrue)
		})

		Convey("IsNull and IsNullOrEmpty should be false", func() {
			So(instance.IsNull(), ShouldBeFalse)
			So(instance.IsNullOrEmpty(), ShouldBeFalse)
		})

		Convey("IsEmpty should return false (delegate to Value)", func() {
			So(instance.IsEmpty(), ShouldBeFalse)
		})

		Convey("Message should return the error string", func() {
			msg := instance.Message()
			So(msg, ShouldNotBeEmpty)
			So(msg, ShouldContainSubstring, "Invalid")
		})

		Convey("IsMessageEqual should match correctly", func() {
			So(instance.IsMessageEqual(instance.Message()), ShouldBeTrue)
			So(instance.IsMessageEqual("not this"), ShouldBeFalse)
		})

		Convey("String and FullString should contain error details", func() {
			So(instance.String(), ShouldContainSubstring, "Invalid")
			So(instance.FullString(), ShouldContainSubstring, "Invalid")
		})

		Convey("ErrorWrapper should return the underlying wrapper", func() {
			So(instance.ErrorWrapper(), ShouldNotBeNil)
			So(instance.ErrorWrapper().HasError(), ShouldBeTrue)
		})

		Convey("HasReferences should delegate to Value", func() {
			So(instance.HasReferences(), ShouldBeFalse)
		})

		Convey("HandleError should panic when there is an error", func() {
			So(func() { instance.HandleError() }, ShouldPanic)
		})

		Convey("HandleErrorWith should panic with concatenated message", func() {
			So(func() { instance.HandleErrorWith("extra") }, ShouldPanic)
		})

		Convey("ConcatNewString should append messages", func() {
			result := instance.ConcatNewString("extra", "info")
			So(result, ShouldContainSubstring, "Invalid")
			So(result, ShouldContainSubstring, "extra")
			So(result, ShouldContainSubstring, "info")
		})
	})

	Convey("Given an Instance with a lazy initializer that returns nil", t, func() {
		initCallCount := 0
		initFunc := func() *errorwrapper.Wrapper {
			initCallCount++
			return nil
		}
		instance := errwonce.New(initFunc)

		Convey("IsNull should be true", func() {
			So(instance.IsNull(), ShouldBeTrue)
		})

		Convey("IsNullOrEmpty should be true", func() {
			So(instance.IsNullOrEmpty(), ShouldBeTrue)
		})

		Convey("HasError should be false", func() {
			So(instance.HasError(), ShouldBeFalse)
		})

		Convey("IsSuccess should be true", func() {
			So(instance.IsSuccess(), ShouldBeTrue)
		})

		Convey("IsFailed should be false", func() {
			So(instance.IsFailed(), ShouldBeFalse)
		})

		Convey("Message should be empty", func() {
			So(instance.Message(), ShouldBeEmpty)
		})

		Convey("IsMessageEqual should return false when null", func() {
			So(instance.IsMessageEqual("anything"), ShouldBeFalse)
		})

		Convey("HandleError should not panic when no error", func() {
			So(func() { instance.HandleError() }, ShouldNotPanic)
		})

		Convey("HandleErrorWith should not panic when no error", func() {
			So(func() { instance.HandleErrorWith("extra") }, ShouldNotPanic)
		})

		Convey("ConcatNewString should return just the additional messages", func() {
			result := instance.ConcatNewString("only", "this")
			So(result, ShouldEqual, "only, this")
		})
	})

	Convey("Given NewPtr", t, func() {
		initFunc := func() *errorwrapper.Wrapper {
			return errnew.NotFound.Simple("item")
		}
		ptr := errwonce.NewPtr(initFunc)

		Convey("pointer instance should work correctly", func() {
			So(ptr, ShouldNotBeNil)
			So(ptr.HasError(), ShouldBeTrue)
		})
	})

	Convey("Given NewPtrUsingErrFunc with an error-returning function", t, func() {
		Convey("when function returns error, Instance should have error", func() {
			errFunc := func() error { return errors.New("boom") }
			ptr := errwonce.NewPtrUsingErrFunc(errtype.Unknown, errFunc)
			So(ptr.HasError(), ShouldBeTrue)
			So(ptr.Message(), ShouldContainSubstring, "boom")
		})

		Convey("when function returns nil, Instance should have no error", func() {
			errFunc := func() error { return nil }
			ptr := errwonce.NewPtrUsingErrFunc(errtype.Unknown, errFunc)
			So(ptr.HasError(), ShouldBeFalse)
		})
	})

	Convey("Given an Instance with references", t, func() {
		initFunc := func() *errorwrapper.Wrapper {
			return errnew.Refs.Quick(errtype.PathMismatch, "path", "loc1", "loc2")
		}
		instance := errwonce.New(initFunc)

		Convey("HasReferences should be true", func() {
			So(instance.HasReferences(), ShouldBeTrue)
		})

		Convey("ConcatNewWrapper should create a new wrapper with messages", func() {
			newWrapper := instance.ConcatNewWrapper(0, "more")
			So(newWrapper, ShouldNotBeNil)
			So(newWrapper.HasError(), ShouldBeTrue)
		})

		Convey("ConcatNewWrapperUsingError should create a new wrapper from error", func() {
			newWrapper := instance.ConcatNewWrapperUsingError(0, errors.New("extra"))
			So(newWrapper, ShouldNotBeNil)
			So(newWrapper.HasError(), ShouldBeTrue)
		})

		Convey("ConcatNewWrapperUsingAnother should create a new wrapper from another wrapper", func() {
			other := errnew.Type.Default(errtype.Invalid)
			newWrapper := instance.ConcatNewWrapperUsingAnother(0, other)
			So(newWrapper, ShouldNotBeNil)
			So(newWrapper.HasError(), ShouldBeTrue)
		})

		Convey("ConcatNewErrors should create a new wrapper from errors", func() {
			newWrapper := instance.ConcatNewErrors(0, errors.New("e1"), errors.New("e2"))
			So(newWrapper, ShouldNotBeNil)
			So(newWrapper.HasError(), ShouldBeTrue)
		})

		Convey("ConcatNew should return an error with concatenated string", func() {
			errResult := instance.ConcatNew("tail")
			So(errResult, ShouldNotBeNil)
			So(errResult.Error(), ShouldContainSubstring, "tail")
		})
	})

	Convey("JSON serialization", t, func() {
		Convey("MarshalJSON for non-empty instance should produce JSON bytes", func() {
			instance := errwonce.New(func() *errorwrapper.Wrapper {
				return errnew.Type.Default(errtype.InvalidOption)
			})
			bytes, err := instance.MarshalJSON()
			So(err, ShouldBeNil)
			So(bytes, ShouldNotBeNil)
		})

		Convey("MarshalJSON for null instance should produce empty string JSON", func() {
			instance := errwonce.New(func() *errorwrapper.Wrapper { return nil })
			bytes, err := instance.MarshalJSON()
			So(err, ShouldBeNil)
			So(string(bytes), ShouldEqual, `""`)
		})

		Convey("UnmarshalJSON should populate the instance", func() {
			instance := errwonce.NewPtr(func() *errorwrapper.Wrapper { return nil })
			wrapper := errnew.Type.Default(errtype.NotFound)
			jsonBytes, _ := wrapper.Serialize()
			err := instance.UnmarshalJSON(jsonBytes)
			So(err, ShouldBeNil)
			So(instance.HasError(), ShouldBeTrue)
			So(instance.Message(), ShouldContainSubstring, "NotFound")
		})
	})
}
