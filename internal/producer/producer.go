package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"pipelineforge/internal/queue"

	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	rmq *queue.RabbitMQ
}

func New(rmq *queue.RabbitMQ) *Producer {
	return &Producer{rmq: rmq}
}

func (p *Producer) Publish(ctx context.Context, msg queue.TrendingRepoMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	return p.rmq.Channel.PublishWithContext(
		ctx,
		"", // default exchange
		p.rmq.Queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp091.Persistent, // survives broker restart
			Body:         body,
		},
	)
}
