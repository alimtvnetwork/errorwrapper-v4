package errstrtests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errstr"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrStr_Result_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errstr.Result
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.HasIssuesOrWhitespace(), ShouldBeTrue)
		So(r.Int(), ShouldEqual, 0)
		So(r.Byte(), ShouldEqual, 0)
		So(r.Bool(), ShouldBeFalse)
		So(r.SafeString(), ShouldEqual, "")
		So(r.SafeBytes(), ShouldResemble, []byte{})
		So(r.IsEqual("x"), ShouldBeFalse)
		So(r.IsEqualIgnoreCase("X"), ShouldBeFalse)
		So(func() { r.Dispose() }, ShouldNotPanic)
		So(func() { r.ValidValue() }, ShouldNotPanic)
		So(func() { r.SimpleStringOnce(true) }, ShouldNotPanic)
		So(func() { r.SimpleStringOnceInit() }, ShouldNotPanic)
		So(r.SplitLines(), ShouldResemble, []string{})
		So(func() { r.SplitLinesSimpleSlice() }, ShouldNotPanic)
	})

	Convey("Empty string without error", t, func() {
		r := &errstr.Result{Value: "", ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.IsEmptyOrWhitespace(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasIssuesOrWhitespace(), ShouldBeTrue)
	})

	Convey("Whitespace string without error", t, func() {
		r := &errstr.Result{Value: "   ", ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.IsEmptyOrWhitespace(), ShouldBeTrue)
		So(r.HasIssuesOrWhitespace(), ShouldBeTrue)
	})

	Convey("Non-empty string without error", t, func() {
		r := &errstr.Result{Value: "hello", ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.IsEmptyOrWhitespace(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasIssuesOrWhitespace(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.String(), ShouldEqual, "hello")
		So(r.SafeString(), ShouldEqual, "hello")
		So(r.Bytes(), ShouldResemble, []byte("hello"))
		So(r.SafeBytes(), ShouldResemble, []byte("hello"))
		So(r.Bool(), ShouldBeFalse)
		So(r.IsEqual("hello"), ShouldBeTrue)
		So(r.IsEqual("world"), ShouldBeFalse)
		So(r.IsEqualIgnoreCase("HELLO"), ShouldBeTrue)
		So(r.IsEqualIgnoreCase("WORLD"), ShouldBeFalse)
	})

	Convey("Numeric string", t, func() {
		r := &errstr.Result{Value: "42", ErrorWrapper: nil}
		So(r.Int(), ShouldEqual, 42)
		So(r.Byte(), ShouldEqual, 42)
	})

	Convey("SplitLines", t, func() {
		r := &errstr.Result{Value: "a\nb\nc", ErrorWrapper: nil}
		So(r.SplitLines(), ShouldResemble, []string{"a", "b", "c"})
	})

	Convey("Value with error", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := &errstr.Result{Value: "x", ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
		So(r.SafeBytes(), ShouldResemble, []byte{})
	})

	Convey("IsEqualResult", t, func() {
		a := &errstr.Result{Value: "same", ErrorWrapper: nil}
		b := &errstr.Result{Value: "same", ErrorWrapper: nil}
		c := &errstr.Result{Value: "diff", ErrorWrapper: nil}
		So(a.IsEqualResult(b), ShouldBeTrue)
		So(a.IsEqualResult(c), ShouldBeFalse)
	})

	Convey("Dispose clears value", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "x")
		r := &errstr.Result{Value: "test", ErrorWrapper: w}
		r.Dispose()
		So(r.Value, ShouldEqual, "")
	})
}

func Test_ErrStr_Result_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errstr.Result{Value: "abc", ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}

func Test_ErrStr_Results_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errstr.Results
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 0)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.IsValid(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.SafeString(), ShouldEqual, "")
		So(r.SafeValues(), ShouldResemble, []string{})
		So(func() { r.Dispose() }, ShouldNotPanic)
		So(func() { r.Clear() }, ShouldNotPanic)
		So(func() { r.List() }, ShouldResemble, []string{})
		So(func() { r.Lines() }, ShouldResemble, []string{})
		So(func() { r.Items() }, ShouldResemble, []string{})
		So(func() { r.Strings() }, ShouldResemble, []string{})
		So(func() { r.Hashset() }, ShouldNotPanic)
		So(func() { r.UniqueMap() }, ShouldNotPanic)
		So(func() { r.StringCollection() }, ShouldNotPanic)
		So(func() { r.SimpleSlice() }, ShouldNotPanic)
	})

	Convey("Empty slice", t, func() {
		r := &errstr.Results{Values: []string{}, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Non-empty slice without error", t, func() {
		r := &errstr.Results{Values: []string{"a", "b"}, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 2)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.SafeValues(), ShouldResemble, []string{"a", "b"})
		So(r.List(), ShouldResemble, []string{"a", "b"})
		So(r.Lines(), ShouldResemble, []string{"a", "b"})
		So(r.Items(), ShouldResemble, []string{"a", "b"})
		So(r.Strings(), ShouldResemble, []string{"a", "b"})
	})

	Convey("Clear resets values", t, func() {
		r := &errstr.Results{Values: []string{"a", "b"}, ErrorWrapper: nil}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values and wrapper", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "x")
		r := &errstr.Results{Values: []string{"a"}, ErrorWrapper: w}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})

	Convey("IsEqualResult default compare", t, func() {
		a := &errstr.Results{Values: []string{"a", "b"}, ErrorWrapper: nil}
		b := &errstr.Results{Values: []string{"a", "b"}, ErrorWrapper: nil}
		c := &errstr.Results{Values: []string{"b", "a"}, ErrorWrapper: nil}
		So(a.IsEqualResultDefault(b), ShouldBeTrue)
		So(a.IsEqualResultDefault(c), ShouldBeFalse)
		So(a.IsIgnoreOrderEqualResult(c), ShouldBeTrue)
	})
}

func Test_ErrStr_Result2_Basics(t *testing.T) {
	Convey("Result2 holds two values", t, func() {
		r := &errstr.Result2{
			Result: errstr.Result{Value: "first", ErrorWrapper: nil},
			Value2: "second",
		}
		So(r.Value, ShouldEqual, "first")
		So(r.Value2, ShouldEqual, "second")
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
	})
}

func Test_ErrStr_ResultWithApplicable_Basics(t *testing.T) {
	Convey("ResultWithApplicable", t, func() {
		r := &errstr.ResultWithApplicable{
			Result:       errstr.Result{Value: "test", ErrorWrapper: nil},
			IsApplicable: true,
		}
		So(r.Value, ShouldEqual, "test")
		So(r.IsApplicable, ShouldBeTrue)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}

func Test_ErrStr_ResultWithApplicable2_Basics(t *testing.T) {
	Convey("ResultWithApplicable2", t, func() {
		r := &errstr.ResultWithApplicable2{
			Result2:      errstr.Result2{Result: errstr.Result{Value: "v1", ErrorWrapper: nil}, Value2: "v2"},
			IsApplicable: true,
		}
		So(r.Value, ShouldEqual, "v1")
		So(r.Value2, ShouldEqual, "v2")
		So(r.IsApplicable, ShouldBeTrue)
	})
}

func Test_ErrStr_ResultsWithErrorCollection_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errstr.ResultsWithErrorCollection
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsEmptyItems(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 0)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.IsValid(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.SafeString(), ShouldEqual, "")
		So(func() { r.Dispose() }, ShouldNotPanic)
		So(func() { r.Clear() }, ShouldNotPanic)
	})

	Convey("Non-empty without error", t, func() {
		r := &errstr.ResultsWithErrorCollection{
			Values:        []string{"a", "b"},
			ErrorWrappers: errwrappers.EmptyCollection(),
		}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 2)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.SafeValues(), ShouldResemble, []string{"a", "b"})
	})

	Convey("Clear resets values", t, func() {
		r := &errstr.ResultsWithErrorCollection{
			Values:        []string{"a", "b"},
			ErrorWrappers: errwrappers.EmptyCollection(),
		}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values", t, func() {
		r := &errstr.ResultsWithErrorCollection{
			Values:        []string{"a"},
			ErrorWrappers: errwrappers.EmptyCollection(),
		}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})
}
