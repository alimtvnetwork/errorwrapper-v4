package reflectinternaltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/internal/reflectinternal"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetElementType(t *testing.T) {
	Convey("GetElementType peels pointer + slice layers down to base type", t, func() {
		So(reflectinternal.GetElementType(nil), ShouldBeNil)

		s := "hi"
		So(reflectinternal.GetElementType(&s), ShouldEqual, reflectinternal.StringType)

		ints := []int{1, 2}
		So(reflectinternal.GetElementType(ints), ShouldEqual, reflectinternal.IntegerType)

		ptrSlice := &[]string{"a"}
		So(reflectinternal.GetElementType(ptrSlice), ShouldEqual, reflectinternal.StringType)
	})

	Convey("GetElementTypeMaxTry returns nil when budget exhausts before reaching a base type", t, func() {
		ints := &[]int{1}
		// 0 tries: peels one pointer layer then runs out of budget → nil
		So(reflectinternal.GetElementTypeMaxTry(ints, 0), ShouldBeNil)
	})
}

func Test_GetElementTypesMaxTry(t *testing.T) {
	Convey("GetElementTypesMaxTry records every visited type", t, func() {
		ptrSlice := &[]int{1}
		final, visited := reflectinternal.GetElementTypesMaxTry(ptrSlice, 4)
		So(final, ShouldEqual, reflectinternal.IntegerType)
		So(len(visited), ShouldBeGreaterThanOrEqualTo, 3) // *[]int, []int, int

		f, v := reflectinternal.GetElementTypesMaxTry(nil, 4)
		So(f, ShouldBeNil)
		So(v, ShouldBeNil)
	})
}

func Test_TypeNameHelpers(t *testing.T) {
	Convey("GetTypeName + IsType", t, func() {
		So(reflectinternal.GetTypeName("hi"), ShouldEqual, "string")
		So(reflectinternal.IsType(42, "int"), ShouldBeTrue)
		So(reflectinternal.IsType(42, "string"), ShouldBeFalse)
	})

	Convey("IsTypeSame", t, func() {
		So(reflectinternal.IsTypeSame(reflectinternal.StringType, reflectinternal.StringType), ShouldBeTrue)
		So(reflectinternal.IsTypeSame(reflectinternal.StringType, reflectinternal.IntegerType), ShouldBeFalse)
	})
}

func Test_GetPointerInfo(t *testing.T) {
	Convey("GetPointerInfo flags pointers and exposes the address", t, func() {
		x := 7
		info := reflectinternal.GetPointerInfo(&x)
		So(info.IsPointer, ShouldBeTrue)
		So(info.Pointer, ShouldNotBeNil)

		info2 := reflectinternal.GetPointerInfo(x)
		So(info2.IsPointer, ShouldBeFalse)
		So(info2.Pointer, ShouldBeNil)
	})
}

func Test_IsBytesOrBytesPointer(t *testing.T) {
	Convey("Detects []byte and *[]byte", t, func() {
		b := []byte("hi")
		ok, p := reflectinternal.IsBytesOrBytesPointer(b)
		So(ok, ShouldBeTrue)
		So(p, ShouldNotBeNil)
		So(string(*p), ShouldEqual, "hi")

		ok2, p2 := reflectinternal.IsBytesOrBytesPointer(&b)
		So(ok2, ShouldBeTrue)
		So(p2, ShouldNotBeNil)

		ok3, _ := reflectinternal.IsBytesOrBytesPointer("hi")
		So(ok3, ShouldBeFalse)
	})
}

func Test_IsStringOrStringPointer(t *testing.T) {
	Convey("Detects string and *string", t, func() {
		ok, p := reflectinternal.IsStringOrStringPointer("hi")
		So(ok, ShouldBeTrue)
		So(*p, ShouldEqual, "hi")

		s := "yo"
		ok2, p2 := reflectinternal.IsStringOrStringPointer(&s)
		So(ok2, ShouldBeTrue)
		So(*p2, ShouldEqual, "yo")

		ok3, _ := reflectinternal.IsStringOrStringPointer(42)
		So(ok3, ShouldBeFalse)
	})

	Convey("IsString returns the deref'd value or empty", t, func() {
		ok, str := reflectinternal.IsString("hello")
		So(ok, ShouldBeTrue)
		So(str, ShouldEqual, "hello")

		ok2, str2 := reflectinternal.IsString(42)
		So(ok2, ShouldBeFalse)
		So(str2, ShouldEqual, "")
	})
}

func Test_IsStringsOrStringsPointer(t *testing.T) {
	Convey("Detects []string and *[]string", t, func() {
		ss := []string{"a", "b"}
		ok, p := reflectinternal.IsStringsOrStringsPointer(ss)
		So(ok, ShouldBeTrue)
		So(p, ShouldNotBeNil)
		So((*p)[0], ShouldEqual, "a")

		ok2, p2 := reflectinternal.IsStringsOrStringsPointer(&ss)
		So(ok2, ShouldBeTrue)
		So(p2, ShouldNotBeNil)

		ok3, _ := reflectinternal.IsStringsOrStringsPointer("a")
		So(ok3, ShouldBeFalse)
	})
}

func Test_IsIntegers_And_Integer(t *testing.T) {
	Convey("IsIntegersOrIntegersPointer", t, func() {
		ints := []int{1, 2}
		ok, p := reflectinternal.IsIntegersOrIntegersPointer(ints)
		So(ok, ShouldBeTrue)
		So((*p)[1], ShouldEqual, 2)

		ok2, p2 := reflectinternal.IsIntegersOrIntegersPointer(&ints)
		So(ok2, ShouldBeTrue)
		So(p2, ShouldNotBeNil)

		ok3, _ := reflectinternal.IsIntegersOrIntegersPointer("nope")
		So(ok3, ShouldBeFalse)
	})

	Convey("IsIntegerOrIntegerPointer + IsInteger", t, func() {
		x := 9
		ok, p := reflectinternal.IsIntegerOrIntegerPointer(&x)
		So(ok, ShouldBeTrue)
		So(*p, ShouldEqual, 9)

		ok2, v := reflectinternal.IsInteger(42)
		So(ok2, ShouldBeTrue)
		So(v, ShouldEqual, 42)

		ok3, v3 := reflectinternal.IsInteger("nope")
		So(ok3, ShouldBeFalse)
		So(v3, ShouldEqual, 0)
	})
}

func Test_IsBoolean(t *testing.T) {
	Convey("IsBoolean handles value, pointer, and non-bool", t, func() {
		ok, v := reflectinternal.IsBoolean(true)
		So(ok, ShouldBeTrue)
		So(v, ShouldBeTrue)

		b := false
		ok2, v2 := reflectinternal.IsBoolean(&b)
		So(ok2, ShouldBeTrue)
		So(v2, ShouldBeFalse)

		ok3, _ := reflectinternal.IsBoolean("nope")
		So(ok3, ShouldBeFalse)
	})

	Convey("IsBooleanPointer returns a usable *bool", t, func() {
		ok, p := reflectinternal.IsBooleanPointer(true)
		So(ok, ShouldBeTrue)
		So(*p, ShouldBeTrue)

		ok2, _ := reflectinternal.IsBooleanPointer("nope")
		So(ok2, ShouldBeFalse)
	})
}

func Test_IsFloat64sOrFloat64sPointer(t *testing.T) {
	Convey("Detects []float64 and *[]float64", t, func() {
		fs := []float64{1.5, 2.5}
		ok, p := reflectinternal.IsFloat64sOrFloat64sPointer(fs)
		So(ok, ShouldBeTrue)
		So((*p)[0], ShouldEqual, 1.5)

		ok2, p2 := reflectinternal.IsFloat64sOrFloat64sPointer(&fs)
		So(ok2, ShouldBeTrue)
		So(p2, ShouldNotBeNil)

		ok3, _ := reflectinternal.IsFloat64sOrFloat64sPointer(42)
		So(ok3, ShouldBeFalse)
	})
}

func Test_NewScanReport(t *testing.T) {
	Convey("NewScanReport for a pointer-to-string", t, func() {
		s := "hi"
		r := reflectinternal.NewScanReport(&s, 4)
		So(r.IsPointer, ShouldBeTrue)
		So(r.IsString, ShouldBeTrue)
		So(r.IsInt, ShouldBeFalse)
		So(r.TypeName, ShouldEqual, "*string")
		So(r.IndirectReflectionType, ShouldNotBeNil)
	})

	Convey("NewScanReport for a plain int", t, func() {
		r := reflectinternal.NewScanReport(7, 4)
		So(r.IsPointer, ShouldBeFalse)
		So(r.IsInt, ShouldBeTrue)
		So(r.IndirectReflectionType, ShouldBeNil)
		So(r.TypeName, ShouldEqual, "int")
	})

	Convey("NewScanReport flags slice / map / func / bool", t, func() {
		So(reflectinternal.NewScanReport([]int{1}, 4).IsSlice, ShouldBeTrue)
		So(reflectinternal.NewScanReport(map[string]int{}, 4).IsMap, ShouldBeTrue)
		So(reflectinternal.NewScanReport(func() {}, 4).IsFunc, ShouldBeTrue)
		So(reflectinternal.NewScanReport(true, 4).IsBool, ShouldBeTrue)
	})
}

func Test_GetFieldValue(t *testing.T) {
	Convey("GetFieldValue returns the underlying interface{} for a struct field", t, func() {
		type S struct {
			Name string
			Age  int
		}
		v := reflect.ValueOf(S{Name: "bob", Age: 42})
		So(reflectinternal.GetFieldValue(v.Field(0)), ShouldEqual, "bob")
		So(reflectinternal.GetFieldValue(v.Field(1)), ShouldEqual, 42)
	})
}
