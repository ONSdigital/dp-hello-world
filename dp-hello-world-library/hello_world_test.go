package helloworld

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewHelloWorld(t *testing.T) {
	Convey("Given a new hello world message", t, func() {
		message := "hello new world!"
		Convey("When NewHelloWorld func is called", func() {
			helloWorld := NewHelloWorld(message)
			Convey("Then HelloWorld object is returned containing new message", func() {
				So(helloWorld, ShouldNotBeNil)
				So(helloWorld, ShouldResemble, &HelloWorld{Message: message})
			})
		})
	})
}
