package server

import (
	"fmt"
	"log"
	"path/filepath"
	"spagen/config"
	"time"

	"github.com/fsnotify/fsnotify"
)

func WatchFiles(restart chan string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	config := config.Get()

	err = watcher.Add("posts")
	if err != nil {
		log.Fatalf("Error adding watcher for posts: %v", err)
	}
	for _, pattern := range config.WatchFiles {
		p := filepath.Dir(pattern)
		err = watcher.Add(p)
		if err != nil {
			log.Fatalf("Error adding watcher for %s: %v", p, err)
		}

		fmt.Println("Watching for changes in:", pattern)
	}

	debounce := make(chan struct{}, 1)
	go func() {
		var timer *time.Timer
		for range debounce {
			if timer != nil {
				timer.Stop()
			}
			timer = time.NewTimer(300 * time.Millisecond)
			<-timer.C
			restart <- "change"
		}
	}()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			fmt.Printf("File %s: OP: %s\n", event.Name, event.Op)
			select {
			case debounce <- struct{}{}:
			default:
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}
