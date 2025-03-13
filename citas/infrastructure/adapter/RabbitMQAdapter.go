package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/lyzsolar/ApiConsumer/citas/domain/entities"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQAdapter struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial("")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("Failed to open a channel: %v", err)
	}

	queue, err := ch.QueueDeclare(
		"cita",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("Failed to declare a queue: %v", err)
	}

	return &RabbitMQAdapter{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}, nil
}

func (r *RabbitMQAdapter) PublishEvent(queueName string, cita entities.Cita) error {
	body, err := json.Marshal(cita)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	log.Printf("Enviando mensaje a la cola: %s", queueName)

	err = r.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Mensaje enviado correctamente: %s", body)
	return nil
}

func (r *RabbitMQAdapter) Send(cita entities.Cita) error {
	return r.PublishEvent(r.queue.Name, cita)
}

func (r *RabbitMQAdapter) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
