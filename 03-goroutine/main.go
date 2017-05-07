package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	err := doit()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func doit() error {
	storyIds, err := QueryNewStoryIds()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(1 * time.Second)
	errors := make(chan error)

	var waitGroup sync.WaitGroup
	for _, storyId := range storyIds[:4] {
		<-ticker.C
		waitGroup.Add(1)
		go func(storyId uint64) {
			defer waitGroup.Done()
			story, err := QueryStory(storyId)
			if err != nil {
				errors <- err
			}
			fmt.Println(story)
		}(storyId)
	}
	waitGroup.Wait()

	select {
	case err, _ := <-errors:
		// Ignore other errors.
		return err
	default:
	}

	return nil
}
