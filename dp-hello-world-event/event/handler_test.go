package event_test

import (
	"os"
	"testing"

	"github.com/ONSdigital/dp-hello-world-event/config"
	"github.com/ONSdigital/dp-hello-world-event/event"
	. "github.com/smartystreets/goconvey/convey"
)

// TODO: remove hello called example test
func TestHelloCalledHandler_Handle(t *testing.T) {
	Convey("Given a successful event handler, when Handle is triggered", t, func() {
		eventHandler := &event.HelloCalledHandler{}
		filePath := "/tmp/helloworld.txt"
		os.Remove(filePath)
		err := eventHandler.Handle(testCtx, &config.Config{OutputFilePath: filePath}, &testEvent)
		So(err, ShouldBeNil)
	})

	Convey("handler returns an error when cannot write to file", t, func() {
		eventHandler := &event.HelloCalledHandler{}
		filePath := ""
		err := eventHandler.Handle(testCtx, &config.Config{OutputFilePath: filePath}, &testEvent)
		So(err, ShouldNotBeNil)
	})
}
