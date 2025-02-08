package utils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func PublishMessage(db_mq *amqp.Connection , queueName string, message string) error {
	if db_mq == nil {
		return fmt.Errorf("❌ RabbitMQ connection is not initialized")
	}

	// Open a channel
	ch, err := db_mq.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queue (ensure it exists)
	_, err = ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	// Publish the message
	err = ch.Publish(
		"",        // Exchange (default)
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	log.Printf("Published message to queue [%s]: %s", queueName, message)
	return nil
}

func ListenToQueue(db_mq *amqp.Connection, queueName string, handlerFunc func(string)) error {
	if db_mq == nil {
		return fmt.Errorf("rabbitMQ connection is not initialized")
	}

	// Open a channel
	ch, err := db_mq.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queue (must match the producer)
	_, err = ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		queueName, // Queue name
		"",        // Consumer tag
		true,      // Auto-ack
		false,     // Exclusive
		false,     // No-local
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %v", err)
	}

	// Process messages
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("received message from queue [%s]: %s", queueName, msg.Body)
			handlerFunc(string(msg.Body)) // Call the custom handler function
		}
	}()

	log.Printf("listening for messages on queue [%s]...", queueName)
	<-forever
	return nil
}