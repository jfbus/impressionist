package filter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	Convey("Scenario: testing dimension parsing", t, func() {

		Convey("Given: a valid dimention", func() {
			str := "100x200"
			Convey("Then: it should parse, reporting no error", func() {
				a, b, err := ParseDimensions(str)
				So(a, ShouldEqual, 100)
				So(b, ShouldEqual, 200)
				So(err, ShouldBeNil)
			})
		})

		Convey("Given: a dimension with negative value", func() {
			str := "100x-200"
			Convey("Then: it should return an error", func() {
				_, _, err := ParseDimensions(str)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Given: a valid rect", func() {
			str := "100x200-10x20"
			Convey("Then: it should parse, reporting no error", func() {
				a, b, c, d, err := ParseRect(str)
				So(a, ShouldEqual, 100)
				So(b, ShouldEqual, 200)
				So(c, ShouldEqual, 10)
				So(d, ShouldEqual, 20)
				So(err, ShouldBeNil)
			})
		})

		Convey("Given: a rect with negative value", func() {
			str := "100x200-10x-20"
			Convey("Then: it should return an error", func() {
				_, _, _, _, err := ParseRect(str)
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Scenario: testing modifier parsing", t, func() {

		Convey("Given: a list of possible modifiers of '+-'", func() {
			Convey("Then: 100x200+ should find '+'", func() {
				str := "100x200+"
				a, b := ParseModifier(str, "+-")
				So(a, ShouldEqual, '+')
				So(b, ShouldEqual, "100x200")
			})
			Convey("Then: 100x200- should find '-'", func() {
				str := "100x200-"
				a, b := ParseModifier(str, "+-")
				So(a, ShouldEqual, '-')
				So(b, ShouldEqual, "100x200")
			})
			Convey("Then: 100x200! should find nothing", func() {
				str := "100x200!"
				a, b := ParseModifier(str, "+-")
				So(a, ShouldEqual, 0)
				So(b, ShouldEqual, "100x200!")
			})
		})

	})
}
