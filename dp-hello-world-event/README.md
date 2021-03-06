dp-hello-world-event
================

ONS service that processes an example event

### Getting started

* Run `make debug`

The service runs in the background consuming messages from Kafka.
An example event can be created using the helper script, `make produce`.

### Dependencies

* Requires running…
  * [kafka](https://github.com/ONSdigital/dp/blob/master/guides/INSTALLING.md#prerequisites)
* No further dependencies other than those defined in `go.mod`

### Configuration

| Environment variable         | Default                           | Description
| ---------------------------- | --------------------------------- | -----------
| BIND_ADDR                    | localhost:8125                    | The host and port to bind to
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5s                                | The graceful shutdown timeout in seconds (`time.Duration` format)
| HEALTHCHECK_INTERVAL         | 30s                               | Time between self-healthchecks (`time.Duration` format)
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s                               | Time to wait until an unhealthy dependent propagates its state to make this app unhealthy (`time.Duration` format)
| KAFKA_ADDR                   | "localhost:9092"                  | The address of Kafka (accepts list)
| KAFKA_OFFSET_OLDEST          | true                              | Start processing Kafka messages in order from the oldest in the queue
| KAFKA_NUM_WORKERS            | 1                                 | The maximum number of parallel kafka consumers
| HELLO_CALLED_GROUP           | dp-hello-world-event              | The consumer group this application to consume ImageUploaded messages
| HELLO_CALLED_TOPIC           | hello-called                      | The name of the topic to consume messages from

### Healthcheck

 The `/health` endpoint returns the current status of the service. Dependent services are health checked on an interval defined by the `HEALTHCHECK_INTERVAL` environment variable.

 On a development machine a request to the health check endpoint can be made by:

 `curl localhost:8125/health`

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2020, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

