package messaging

import (
	"log"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	amqp "github.com/rabbitmq/amqp091-go"
)

// ProducerSendMessage sends a message to the message queue.
// It opens a channel, enables publisher confirms, declares a queue,
// and publishes the message to the queue.
// If the message is successfully published, it logs the message as sent.
// If there is an error during any step, it logs the error and terminates the program.
func ProducerSendMessage(message string) {
	ch, err := global.MessageQueue.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Bật Publisher Confirms
	if err := ch.Confirm(false); err != nil {
		log.Fatalf("Failed to put channel in confirm mode: %s", err)
	}

	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	q, err := ch.QueueDeclare(
		constants.KeyAuthPro, // name
		true,                 // durable: If true, the queue will survive broker restarts.
		false,                // delete when unused: If true, the queue will be deleted when there are no consumers.
		false,                // exclusive: If true, the queue can only be used by the connection that declares it and will be deleted when the connection closes.
		false,                // no-wait: If true, the queue declaration will not wait for a confirmation from the server.
		nil,                  // arguments: Optional additional arguments (e.g., for message TTL, queue length limits).
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	err = ch.Publish(
		"",     // exchange: The name of the exchange to publish to. An empty string indicates the default exchange.
		q.Name, // routing key: The routing key for the message. Used to route messages to queues in certain exchange types.
		false,  // mandatory: If true, the server will return an unroutable message with a mandatory flag set.
		false,  // immediate: If true, the server will return an immediate response if no consumer is available.

		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	select {
	case confirmed := <-confirms:
		if confirmed.Ack {
			log.Println("Queue status:", constants.KeyAuthPro)
			log.Printf("[x] Sent %s", message)
		} else {
			log.Printf("Failed to confirm message: %s", message)
		}
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		log.Printf("Confirmation timeout for message: %s", message)
	}
}
