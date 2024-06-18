package messaging

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumerMessages() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := global.MessageQueue.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		constants.KeyAuthPro, // queue name
		true,                 // durable: If true, the queue will survive broker restarts.
		false,                // delete when unused: If true, the queue will be deleted when there are no consumers.
		false,                // exclusive: If true, the queue can only be used by the connection that declares it and will be deleted when the connection closes.
		false,                // no-wait: If true, the queue declaration will not wait for a confirmation from the server.
		nil,                  // arguments: Optional additional arguments (e.g., for message TTL, queue length limits).
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // manual ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	var wg sync.WaitGroup
	messageBuffer := make(chan amqp.Delivery, 10) // Buffer size of 10

	// Start worker goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for d := range messageBuffer {
				if err := processMessage(ctx, d); err != nil {
					log.Printf("Worker %d: Error processing message: %s", workerID, err)
					// Requeue or handle the error
				}
				d.Ack(false)
			}
		}(i)
	}

	// Signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Main loop to receive messages
	go func() {
		for {
			select {
			case d, ok := <-msgs:
				if !ok {
					close(messageBuffer)
					return
				}
				messageBuffer <- d
			case sig := <-sigChan:
				log.Printf("Received signal: %s. Shutting down...", sig)
				cancel()
				close(messageBuffer)
				return
			}
		}
	}()

	wg.Wait() // Wait for all workers to finish
	log.Println("All workers have finished processing")
}

func processMessage(ctx context.Context, d amqp.Delivery) error {
	// Implement your message processing logic here with retry mechanism
	log.Printf("Received a message: %s", d.Body)
	// Simulating processing time
	time.Sleep(2 * time.Second)

	// Example: Implement retry logic
	for retry := 0; retry < 3; retry++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Simulate message processing
			if err := simulateProcessing(d.Body); err != nil {
				log.Printf("Processing failed, retrying (%d/3): %s", retry+1, err)
				time.Sleep(1 * time.Second)
			} else {
				return nil
			}
		}
	}
	return errors.New("message processing failed after retries")
}

func simulateProcessing(_ []byte) error {
	// Replace with actual processing logic
	return nil
}
