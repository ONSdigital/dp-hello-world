package event_test

import (
	"github.com/ONSdigital/dp-hello-world-event/event"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func TestHelloCalledHandler_Handle(t *testing.T) {

	Convey("Given a successful event handler, when Handle is triggered it writes the message to a file", t, func() {

		eventHandler := event.HelloCalledHandler{}
		err := eventHandler.Handle(testCtx, &testEvent)
		So(err, ShouldBeNil)
		data, err := ioutil.ReadFile("/tmp/helloworld.txt")
		So(string(data), ShouldEqual, "Hello, " + testEvent.RecipientName + "!")
	})

}
