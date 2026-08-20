package refstests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/ref"
	"github.com/alimtvnetwork/errorwrapper-v4/refs"
	. "github.com/smartystreets/goconvey/convey"
)

// Test_Ref_New — basic Value construction stores name + value.
func Test_Ref_New(t *testing.T) {
	Convey("ref.New stores VarName and value", t, func() {
		v := ref.New("userId", 42)
		So(v.VarName(), ShouldEqual, "userId")
		So(v.ValueAny(), ShouldEqual, 42)
	})
}

// Test_Ref_Compile — Compile produces non-empty string containing the name.
func Test_Ref_Compile(t *testing.T) {
	Convey("ref.Value.Compile contains the variable name and value", t, func() {
		v := ref.New("email", "a@b.co")
		s := v.Compile()
		So(s, ShouldContainSubstring, "email")
		So(s, ShouldContainSubstring, "a@b.co")
	})
}

// Test_Ref_IsEqual — equal Values match, different ones don't.
func Test_Ref_IsEqual(t *testing.T) {
	Convey("ref.Value.IsEqual matches identical Values", t, func() {
		a := ref.New("k", "v")
		b := ref.New("k", "v")
		c := ref.New("k", "x")
		So(a.IsEqual(b), ShouldBeTrue)
		So(a.IsEqual(c), ShouldBeFalse)
	})
}

// Test_Refs_EmptyPtr — empty collection has zero items.
func Test_Refs_EmptyPtr(t *testing.T) {
	Convey("refs.EmptyPtr returns a collection with Count 0", t, func() {
		c := refs.EmptyPtr()
		So(c, ShouldNotBeNil)
		So(c.Count(), ShouldEqual, 0)
	})
}

// Test_Refs_AddVarVal — adding entries increments count.
func Test_Refs_AddVarVal(t *testing.T) {
	Convey("AddVarVal appends to collection", t, func() {
		c := refs.EmptyPtr()
		c.AddVarVal("name", "alice")
		c.AddVarVal("age", 30)
		So(c.Count(), ShouldEqual, 2)
		m := c.MapStringAny()
		So(m["name"], ShouldEqual, "alice")
		So(m["age"], ShouldEqual, 30)
	})
}

// Test_Refs_NewDirectItem — convenience single-item constructor.
func Test_Refs_NewDirectItem(t *testing.T) {
	Convey("NewDirectItem builds a one-item collection", t, func() {
		c := refs.NewDirectItem("token", "abc")
		So(c.Count(), ShouldEqual, 1)
		So(c.Compile(), ShouldContainSubstring, "token")
	})
}
