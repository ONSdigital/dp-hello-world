package event_test

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/ONSdigital/dp-hello-world-event/config"

	"github.com/ONSdigital/dp-hello-world-event/event"
	"github.com/ONSdigital/dp-hello-world-event/event/mock"
	"github.com/ONSdigital/dp-hello-world-event/schema"
	kafka "github.com/ONSdigital/dp-kafka/v4"
	"github.com/ONSdigital/dp-kafka/v4/kafkatest"
	. "github.com/smartystreets/goconvey/convey"
)

var testCtx = context.Background()

var errHandler = errors.New("handler Error")

var testEvent = event.HelloCalled{
	RecipientName: "World",
}

// TODO: remove or replace hello called logic with app specific
func TestConsume(t *testing.T) {
	Convey("Given kafka consumer and event handler mocks", t, func() {
		cgChannels := &kafka.ConsumerGroupChannels{Upstream: make(chan kafka.Message, 2)}
		mockConsumer := &kafkatest.IConsumerGroupMock{
			ChannelsFunc: func() *kafka.ConsumerGroupChannels { return cgChannels },
		}

		handlerWg := &sync.WaitGroup{}
		mockEventHandler := &mock.HandlerMock{
			HandleFunc: func(ctx context.Context, config *config.Config, event *event.HelloCalled) error {
				defer handlerWg.Done()
				return nil
			},
		}

		Convey("And a kafka message with the valid schema being sent to the Upstream channel", func() {
			message, err := kafkatest.NewMessage(marshal(testEvent), 0)
			if err != nil {
				t.Errorf("unable to create new message")
			}

			mockConsumer.Channels().Upstream <- message

			Convey("When consume message is called", func() {
				handlerWg.Add(1)
				event.Consume(testCtx, mockConsumer, mockEventHandler, &config.Config{KafkaConfig: config.KafkaConfig{NumWorkers: 1}})
				handlerWg.Wait()

				Convey("An event is sent to the mockEventHandler ", func() {
					So(len(mockEventHandler.HandleCalls()), ShouldEqual, 1)
					So(*mockEventHandler.HandleCalls()[0].HelloCalled, ShouldResemble, testEvent)
				})

				Convey("The message is committed and the consumer is released", func() {
					<-message.UpstreamDone()
					So(len(message.CommitCalls()), ShouldEqual, 1)
					So(len(message.ReleaseCalls()), ShouldEqual, 1)
				})
			})
		})

		Convey("And two kafka messages, one with a valid schema and one with an invalid schema", func() {
			validMessage, err := kafkatest.NewMessage(marshal(testEvent), 1)
			if err != nil {
				t.Errorf("unable to create new message")
			}

			invalidMessage, err := kafkatest.NewMessage([]byte("invalid schema"), 0)
			if err != nil {
				t.Errorf("unable to create new message")
			}

			mockConsumer.Channels().Upstream <- invalidMessage
			mockConsumer.Channels().Upstream <- validMessage

			Convey("When consume messages is called", func() {
				handlerWg.Add(1)
				event.Consume(testCtx, mockConsumer, mockEventHandler, &config.Config{KafkaConfig: config.KafkaConfig{NumWorkers: 1}})
				handlerWg.Wait()

				Convey("Only the valid event is sent to the mockEventHandler ", func() {
					So(len(mockEventHandler.HandleCalls()), ShouldEqual, 1)
					So(*mockEventHandler.HandleCalls()[0].HelloCalled, ShouldResemble, testEvent)
				})

				Convey("Only the valid message is committed, but the consumer is released for both messages", func() {
					<-validMessage.UpstreamDone()
					<-invalidMessage.UpstreamDone()
					So(len(validMessage.CommitCalls()), ShouldEqual, 1)
					So(len(invalidMessage.CommitCalls()), ShouldEqual, 1)
					So(len(validMessage.ReleaseCalls()), ShouldEqual, 1)
					So(len(invalidMessage.ReleaseCalls()), ShouldEqual, 1)
				})
			})
		})

		Convey("With a failing handler and a kafka message with the valid schema being sent to the Upstream channel", func() {
			mockEventHandler.HandleFunc = func(ctx context.Context, config *config.Config, event *event.HelloCalled) error {
				defer handlerWg.Done()
				return errHandler
			}

			message, err := kafkatest.NewMessage(marshal(testEvent), 0)
			if err != nil {
				t.Errorf("unable to create new message")
			}

			mockConsumer.Channels().Upstream <- message

			Convey("When consume message is called", func() {
				handlerWg.Add(1)
				event.Consume(testCtx, mockConsumer, mockEventHandler, &config.Config{KafkaConfig: config.KafkaConfig{NumWorkers: 1}})
				handlerWg.Wait()

				Convey("An event is sent to the mockEventHandler ", func() {
					So(len(mockEventHandler.HandleCalls()), ShouldEqual, 1)
					So(*mockEventHandler.HandleCalls()[0].HelloCalled, ShouldResemble, testEvent)
				})

				Convey("The message is committed and the consumer is released", func() {
					<-message.UpstreamDone()
					So(len(message.CommitCalls()), ShouldEqual, 1)
					So(len(message.ReleaseCalls()), ShouldEqual, 1)
				})
			})
		})
	})
}

// marshal helper method to marshal a event into a []byte
func marshal(e event.HelloCalled) []byte {
	bytes, err := schema.HelloCalledEvent.Marshal(e)
	So(err, ShouldBeNil)
	return bytes
}
