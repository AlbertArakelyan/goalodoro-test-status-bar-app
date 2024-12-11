package main

import (
	"sync"
	"time"

	"github.com/caseymrm/menuet"
)

var (
	stop        chan struct{}
	clockActive sync.Mutex // Mutex to ensure only one clock is active at a time
)

func switchClock(newItem string) {
	if stop != nil {
		close(stop) // Signal the currently running clock to stop
		stop = nil  // Reset stop channel
	}
	activeItem = newItem
	stop = make(chan struct{}) // Create a new stop channel
	go goalClock(newItem)      // Start new clock goroutine
}

func goalClock(str string) {
	clockActive.Lock() // Ensure only one clock can run
	defer clockActive.Unlock()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stop: // Stop signal received, exit the function
			return
		case t := <-ticker.C: // Update every second
			menuet.App().SetMenuState(&menuet.MenuState{
				Title: str + " " + t.Format(":05"), // TODO write a separate logic/for function for tracking goal
			})
		}
	}
}
