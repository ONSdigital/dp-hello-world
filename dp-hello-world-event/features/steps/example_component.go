package steps

import (
	"context"
	"net/http"
	"os"

	componenttest "github.com/ONSdigital/dp-component-test"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-hello-world-event/config"
	"github.com/ONSdigital/dp-hello-world-event/service"
	"github.com/ONSdigital/dp-hello-world-event/service/mock"
	kafka "github.com/ONSdigital/dp-kafka/v4"
	"github.com/ONSdigital/dp-kafka/v4/kafkatest"
	dphttp "github.com/ONSdigital/dp-net/v2/http"
)

type Component struct {
	componenttest.ErrorFeature
	serviceList   *service.ExternalServiceList
	KafkaConsumer kafka.IConsumerGroup
	errorChan     chan error
	svc           *service.Service
	cfg           *config.Config
}

func NewComponent() *Component {
	cfg, err := config.Get()
	if err != nil {
		return nil
	}

	c := &Component{errorChan: make(chan error)}

	minBrokers := 1

	consumer, err := kafkatest.NewConsumer(context.Background(),
		&kafka.ConsumerGroupConfig{
			BrokerAddrs:       cfg.KafkaConfig.Brokers,
			Topic:             cfg.KafkaConfig.HelloCalledTopic,
			GroupName:         cfg.KafkaConfig.HelloCalledGroup,
			MinBrokersHealthy: &minBrokers,
			KafkaVersion:      &cfg.KafkaConfig.Version,
		}, nil,
	)
	if err != nil {
		return nil
	}

	consumer.Mock.CheckerFunc = funcCheck
	c.KafkaConsumer = consumer.Mock

	c.cfg = cfg

	initMock := &mock.InitialiserMock{
		DoGetKafkaConsumerFunc: c.DoGetConsumer,
		DoGetHealthCheckFunc:   c.DoGetHealthCheck,
		DoGetHTTPServerFunc:    c.DoGetHTTPServer,
	}

	c.serviceList = service.NewServiceList(initMock)

	return c
}

func (c *Component) Close() {
	os.Remove(c.cfg.OutputFilePath)
}

func (c *Component) Reset() {
	os.Remove(c.cfg.OutputFilePath)
}

func (c *Component) DoGetHealthCheck(cfg *config.Config, buildTime, gitCommit, version string) (service.HealthChecker, error) {
	return &mock.HealthCheckerMock{
		AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
		StartFunc:    func(ctx context.Context) {},
		StopFunc:     func() {},
	}, nil
}

func (c *Component) DoGetHTTPServer(bindAddr string, router http.Handler) service.HTTPServer {
	return dphttp.NewServer(bindAddr, router)
}

func (c *Component) DoGetConsumer(ctx context.Context, kafkaCfg *config.KafkaConfig) (kafkaConsumer kafka.IConsumerGroup, err error) {
	return c.KafkaConsumer, nil
}

func funcCheck(ctx context.Context, state *healthcheck.CheckState) error {
	return nil
}
