package event

import (
	"context"

	"github.com/ONSdigital/dp-hello-world-event/config"
	"github.com/ONSdigital/dp-hello-world-event/schema"
	kafka "github.com/ONSdigital/dp-kafka/v4"
	"github.com/ONSdigital/log.go/v2/log"
)

//go:generate moq -out mock/handler.go -pkg mock . Handler

// TODO: remove or replace hello called logic with app specific
// Handler represents a handler for processing a single event.
type Handler interface {
	Handle(ctx context.Context, cfg *config.Config, helloCalled *HelloCalled) error
}

// Consume converts messages to event instances, and pass the event to the provided handler.
func Consume(ctx context.Context, messageConsumer kafka.IConsumerGroup, handler Handler, cfg *config.Config) {
	// consume loop, to be executed by each worker
	var consume = func(workerID int) {
		for {
			select {
			case message, ok := <-messageConsumer.Channels().Upstream:
				if !ok {
					log.Info(ctx, "closing event consumer loop because upstream channel is closed", log.Data{"worker_id": workerID})
					return
				}
				messageCtx := context.Background()
				processMessage(messageCtx, message, handler, cfg)
				message.Release()
			case <-messageConsumer.Channels().Closer:
				log.Info(ctx, "closing event consumer loop because closer channel is closed", log.Data{"worker_id": workerID})
				return
			}
		}
	}

	// workers to consume messages in parallel
	for w := 1; w <= cfg.KafkaConfig.NumWorkers; w++ {
		go consume(w)
	}
}

// processMessage unmarshals the provided kafka message into an event and calls the handler.
// After the message is handled, it is committed.
func processMessage(ctx context.Context, message kafka.Message, handler Handler, cfg *config.Config) {
	// unmarshal - commit on failure (consuming the message again would result in the same error)
	event, err := unmarshal(message)
	if err != nil {
		log.Error(ctx, "failed to unmarshal event", err)
		message.Commit()
		return
	}

	log.Info(ctx, "event received", log.Data{"event": event})

	// handle - commit on failure (implement error handling to not commit if message needs to be consumed again)
	err = handler.Handle(ctx, cfg, event)
	if err != nil {
		log.Error(ctx, "failed to handle event", err)
		message.Commit()
		return
	}

	log.Info(ctx, "event processed - committing message", log.Data{"event": event})
	message.Commit()
	log.Info(ctx, "message committed", log.Data{"event": event})
}

// unmarshal converts a event instance to []byte.
func unmarshal(message kafka.Message) (*HelloCalled, error) {
	var event HelloCalled
	err := schema.HelloCalledEvent.Unmarshal(message.GetData(), &event)
	return &event, err
}
