package event_test

import (
	"github.com/ONSdigital/dp-hello-world-event/event"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestImageUploadedHandler_Handle(t *testing.T) {

	Convey("Given S3 and Vault client mocks", t, func() {

		Convey("And a successful event handler, when Handle is triggered", func() {

			eventHandler := event.HelloCalledHandler{}
			err := eventHandler.Handle(testCtx, &testEvent)
			So(err, ShouldBeNil)
		})

	})

}
