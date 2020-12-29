package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-hello-world-event
type Config struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	Brokers                    []string      `envconfig:"KAFKA_ADDR"                     json:"-"`
	HelloCalledGroup           string        `envconfig:"HELLO_CALLED_GROUP"`
	HelloCalledTopic           string        `envconfig:"HELLO_CALLED_TOPIC"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg := &Config{
		BindAddr:                   "localhost:8125",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		Brokers:                    []string{"localhost:9092"},
		HelloCalledGroup:           "dp-hello-world-event",
		HelloCalledTopic:           "hello-called",
	}

	return cfg, envconfig.Process("", cfg)
}
