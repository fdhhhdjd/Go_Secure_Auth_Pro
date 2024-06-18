package initialization

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ConnectRabbitMQ establishes a connection to RabbitMQ using the provided DSN.
// It retries the connection for a maximum number of times with a delay between retries.
// If the connection is successful, it returns the RabbitMQ connection object and nil error.
// If the connection fails after the maximum number of attempts, it returns nil connection
// and an error indicating the failure.
func ConnectRabbitMQ(dsn string) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error

	// Maximum number of retries
	maxRetries := 5
	// Delay between retries
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(dsn)
		if err == nil {
			log.Println("CONNECTED SUCCESS RABBITMQ 🐇!")
			return conn, nil
		}

		log.Printf("Failed to connect to RabbitMQ: %s", err)
		log.Printf("Retrying in %v...", retryDelay)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("failed to connect to RabbitMQ after %d attempts: %w", maxRetries, err)
}
