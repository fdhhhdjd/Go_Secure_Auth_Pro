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
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
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
