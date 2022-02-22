package event_test

import (
	"testing"

	"github.com/ONSdigital/dp-hello-world-event/event"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloCalledHandler_Handle(t *testing.T) {

	Convey("Given a successful event handler, when Handle is triggered", t, func() {

		eventHandler := event.HelloCalledHandler{}
		err := eventHandler.Handle(testCtx, &testEvent)
		So(err, ShouldBeNil)
	})

}
