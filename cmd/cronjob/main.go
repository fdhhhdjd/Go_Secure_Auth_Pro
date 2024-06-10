package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	// Add a cron job that runs once after 1 minute
	_, err := c.AddFunc("@every 1m", func() {
		fmt.Println("Hello, World!")
		// Stop the cron scheduler after the first run
		c.Stop()
	})

	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// Start the cron scheduler
	c.Start()

	// Schedule a task to stop the program after 2 minutes
	time.AfterFunc(2*time.Minute, func() {
		fmt.Println("Program exiting gracefully...")
	})

	// Keep the program running until the cron job executes and the program exits
	<-time.After(3 * time.Minute)
}
