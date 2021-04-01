package event

import (
	"context"
	"fmt"
	"github.com/ONSdigital/log.go/log"
	"io/ioutil"
)

// HelloCalledHandler ...
type HelloCalledHandler struct {
}

// Handle takes a single event.
func (h *HelloCalledHandler) Handle(ctx context.Context, event *HelloCalled) (err error) {
	logData := log.Data{
		"event": event,
	}
	log.Event(ctx, "event handler called", log.INFO, logData)

	//TODO Replace with actual event handler logic…
	greeting := fmt.Sprintf("Hello, %s!", event.RecipientName)
	_ = ioutil.WriteFile("/tmp/helloworld.txt", []byte(greeting), 0644)

	logData["greeting"] = greeting
	log.Event(ctx, "hello world example handler called successfully", log.INFO, logData)

	log.Event(ctx, "event successfully handled", log.INFO, logData)
	return nil
}
