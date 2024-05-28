package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

func freePort(port int) {
	cmd := exec.Command("fuser", "-k", fmt.Sprintf("%d/tcp", port))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("Failed to free port %d: %v", port, err)
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	var cmd *exec.Cmd
	var mu sync.Mutex
	var debounceTimer *time.Timer

	buildAndRestart := func() {
		mu.Lock()
		defer mu.Unlock()

		// Build the application
		buildCmd := exec.Command("go", "build", "-o", "./tmp/main", "./cmd/server/main.go")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			log.Println("Build failed:", err)
			return
		}

		// Free the port
		freePort(8000)

		// Restart the application
		if cmd != nil && cmd.Process != nil {
			if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
				log.Printf("Failed to terminate process: %v", err)
			}
			_ = cmd.Wait()
		}
		cmd = exec.Command("./tmp/main")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		log.Println("Application restarted")
	}

	go func() {
		// Call buildAndRestart immediately on start
		buildAndRestart()

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if (event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create) && filepath.Ext(event.Name) == ".go" {
					log.Println("Modified file:", event.Name)
					if debounceTimer != nil {
						debounceTimer.Stop()
					}
					debounceTimer = time.AfterFunc(500*time.Millisecond, buildAndRestart)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			switch info.Name() {
			case "vendor", "tmp", "node_modules", ".git":
				return filepath.SkipDir
			default:
				return watcher.Add(path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	<-done
}
