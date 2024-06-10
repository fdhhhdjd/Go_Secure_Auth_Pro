package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("@every 1m", func() {
		fmt.Println("Job running every minute: Tai dev!")
	})

	if err != nil {
		fmt.Println("Error adding cron job to run every 1 minute:", err)
		return
	}

	var hourJobDone sync.Once

	_, err = c.AddFunc("0 0 * * * *", func() {
		hourJobDone.Do(func() {
			err := repo.UpdateVerificationBulk(global.DB)
			if err != nil {
				fmt.Println("Error updating verification records:", err)
			}
			fmt.Println("Job running every hour: Update Verification not used!")
		})
	})

	if err != nil {
		fmt.Println("Error adding cron job to run every hour:", err)
		return
	}

	c.Start()

	var wg sync.WaitGroup
	wg.Add(1)

	time.AfterFunc(1*time.Minute, func() {
		fmt.Println("Program exiting gracefully...")
		c.Stop()
		wg.Done()
	})

	wg.Wait()
}
