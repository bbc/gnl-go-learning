package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	Convey("GIVEN the test environment is prepared", t, func() {
		Convey("THEN we get the expected message", func() {
			Convey("WHEN we call main().", func() {
				So(getGreeting(), ShouldResemble, `Hello, World!`)
			})
		})
	})
}
