package main

import (
	"fmt"
	"time"
)

// Task structure
type Task struct {
	ID int
}

// Log format
const logFmt = "%-15s %2d at %s\n"

// Run Task method
func (rcv *Task) Run() {
	fmt.Printf(logFmt, "Started task", rcv.ID, time.Now())
	time.Sleep(time.Second)
	fmt.Printf(logFmt, "Finished task", rcv.ID, time.Now())
}

// Write access to global variables within goroutines is thread unsafe.
// Create a channel to communicate in gorounites and control number
// of already running goroutines.
func main() {
	// Array of Task structures
	var taskList = []Task{}

	// Maximum number of tasks
	maxTaskCounter := 50

	// Populate task array
	for i := 1; i <= maxTaskCounter; i++ {
		taskList = append(taskList, Task{i})
	}

	// Number of naximum concurrent tasks
	maxConcurrentTaskCounter := 10

	// Number of current running concurrent tasks
	taskCounterChan := make(chan struct{}, maxConcurrentTaskCounter)
	defer close(taskCounterChan)

	for _, task := range taskList {
		taskCounterChan <- struct{}{}
		fmt.Printf(logFmt, "Enqueue task", task.ID, time.Now())

		go func(task Task) {
			task.Run()
			<-taskCounterChan
		}(task)
	}

	// Wait for all tasks to complete.
	// Alternate soultion is to use sync.WaitGroup
	for len(taskCounterChan) > 0 {
	}
}
