package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"pipelineforge/internal/metrics"
	"pipelineforge/internal/queue"
)

type Consumer struct {
	rmq     *queue.RabbitMQ
	handler *Handler
}

func NewConsumer(rmq *queue.RabbitMQ, handler *Handler) *Consumer {
	return &Consumer{rmq: rmq, handler: handler}
}

func (c *Consumer) Start(ctx context.Context) error {
	msgs, err := c.rmq.Channel.Consume(
		c.rmq.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			start := time.Now()
			metrics.MessagesConsumed.Inc()

			var msg queue.TrendingRepoMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				_ = d.Nack(false, false)
				continue
			}

			if err := c.handler.Handle(ctx, msg); err != nil {
				metrics.DBErrors.Inc()
				log.Println("db error:", err)
				_ = d.Nack(false, true)
				continue
			}

			metrics.DBInserts.Inc()
			metrics.MessageProcessingDuration.Observe(time.Since(start).Seconds())
			_ = d.Ack(false)
		}
	}()

	return nil
}
