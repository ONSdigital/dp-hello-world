package event

import (
	"context"

	"github.com/ONSdigital/dp-hello-world-event/schema"
	kafka "github.com/ONSdigital/dp-kafka/v2"
	"github.com/ONSdigital/log.go/log"
)

//go:generate moq -out mock/handler.go -pkg mock . Handler

// Handler represents a handler for processing a single event.
type Handler interface {
	Handle(ctx context.Context, helloCalled *HelloCalled) error
}

// Consume converts messages to event instances, and pass the event to the provided handler.
func Consume(ctx context.Context, messageConsumer kafka.IConsumerGroup, handler Handler, numWorkers int) {

	// consume loop, to be executed by each worker
	var consume = func(workerID int) {
		for {
			select {
			case message, ok := <-messageConsumer.Channels().Upstream:
				if !ok {
					log.Event(ctx, "closing event consumer loop because upstream channel is closed", log.INFO, log.Data{"worker_id": workerID})
					return
				}
				messageCtx := context.Background()
				processMessage(messageCtx, message, handler)
				message.Release()
			case <-messageConsumer.Channels().Closer:
				log.Event(ctx, "closing event consumer loop because closer channel is closed", log.INFO, log.Data{"worker_id": workerID})
				return
			}
		}
	}

	// workers to consume messages in parallel
	for w := 1; w <= numWorkers; w++ {
		go consume(w)
	}
}

// processMessage unmarshals the provided kafka message into an event and calls the handler.
// After the message is handled, it is committed.
func processMessage(ctx context.Context, message kafka.Message, handler Handler) {

	// unmarshal - commit on failure (consuming the message again would result in the same error)
	event, err := unmarshal(message)
	if err != nil {
		log.Event(ctx, "failed to unmarshal event", log.ERROR, log.Error(err))
		message.Commit()
		return
	}

	log.Event(ctx, "event received", log.INFO, log.Data{"event": event})

	// handle - commit on failure (implement error handling to not commit if message needs to be consumed again)
	err = handler.Handle(ctx, event)
	if err != nil {
		log.Event(ctx, "failed to handle event", log.ERROR, log.Error(err))
		message.Commit()
		return
	}

	log.Event(ctx, "event processed - committing message", log.INFO, log.Data{"event": event})
	message.Commit()
	log.Event(ctx, "message committed", log.INFO, log.Data{"event": event})
}

// unmarshal converts a event instance to []byte.
func unmarshal(message kafka.Message) (*HelloCalled, error) {
	var event HelloCalled
	err := schema.HelloCalledEvent.Unmarshal(message.GetData(), &event)
	return &event, err
}
