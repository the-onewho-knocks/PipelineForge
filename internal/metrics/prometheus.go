package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MessagesConsumed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "pipelineforge",
			Subsystem: "worker",
			Name:      "messages_consumed_total",
			Help:      "Total number of RabbitMQ messages consumed",
		},
	)

	DBInserts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "pipelineforge",
			Subsystem: "postgres",
			Name:      "inserts_total",
			Help:      "Total number of successful DB inserts",
		},
	)

	DBErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "pipelineforge",
			Subsystem: "postgres",
			Name:      "errors_total",
			Help:      "Total number of DB errors",
		},
	)

	MessageProcessingDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "pipelineforge",
			Subsystem: "worker",
			Name:      "message_processing_seconds",
			Help:      "Time taken to process a single message",
			Buckets:   prometheus.DefBuckets,
		},
	)
)

func Register() {
	prometheus.MustRegister(
		MessagesConsumed,
		DBInserts,
		DBErrors,
		MessageProcessingDuration,
	)
}
