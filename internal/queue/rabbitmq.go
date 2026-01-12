package queue

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

//connection to rabbitMQ queueing is done here
type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
	Queue   amqp091.Queue
}

func NewRabbitMQ(url, queueName string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq dial: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel open: %w", err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("queue declare: %w", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
		Queue:   q,
	}, nil
}

func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		_ = r.Channel.Close()
	}
	if r.Conn != nil {
		_ = r.Conn.Close()
	}
}
