package main

import (
	"fmt"

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

	// Keep the program running to ensure the cron job executes
	// The cron job will run once after 1 minute, then the scheduler will stop
	select {}
}
