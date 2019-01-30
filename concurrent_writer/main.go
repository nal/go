package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Log format
const logFmt = "%-15s %2d at %s\n"

// Task structure
type Task struct {
	ID int
}

// LogWriter struct
type LogWriter struct {
	filePath string
	fd       *os.File
}

// Create and write to file hander.
func (logW LogWriter) Write(buf *bytes.Buffer) {
	ctime := time.Now()
	hour := ctime.Hour()
	min := ctime.Minute()

	// File is closed or rotating
	if logW.fd == nil {
		logW.filePath = fmt.Sprintf("tmp_%02d%02d.txt", hour, min)

		f, err := os.OpenFile(logW.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Store new file handler
		logW.fd = f
	}

	logW.fd.Write(buf.Bytes())
	logW.fd.Sync()
	buf.Reset()
}

// Close filehandler
func (logW LogWriter) Close() {
	// File is opened
	if logW.fd != nil {
		if err := logW.fd.Close(); err != nil {
			log.Println(err)
		}
	}
}

// WriteLog Task method
func (rcv Task) WriteLog(logger *log.Logger, logChan chan<- struct{}, taskCounterChan <-chan struct{}) {
	logger.Printf(logFmt, "Started task", rcv.ID, time.Now())

	// Sleep for random time more than 60 seconds to test log file rotation.
	sleepSeconds := time.Duration(rand.Int31n(60)+20) * time.Second
	logger.Printf("%-15s %2d sleep for %v\n", "Task", rcv.ID, sleepSeconds)
	time.Sleep(sleepSeconds)

	logger.Printf(logFmt, "Finished task", rcv.ID, time.Now())

	// Block channel and notify we have filled a buffer
	logChan <- struct{}{}

	// Decrease number of worker goroutines
	<-taskCounterChan
}

// Solution: main function spawns worker goroutines with `task.WriteLog` method.
// These methods write data to `*bytes.Buffer` and blocks.
// In `for-select-case` block we read data from 4 channels:
// * flush file timer (to flush buffer to file and flush file to filesystem)
// * rotate file timer (to do the same as flush file timer plus close file handler)
// * log channel (write to buffer and write it to log file)
// * worker count channel (abort execution, all workers finished)
func main() {
	// Init empty LogWriter
	logW := LogWriter{filePath: "", fd: nil}
	defer logW.Close()

	// Array of Task structures
	var taskList = []Task{}

	// Maximum number of tasks
	maxTaskCounter := 10

	// Populate task array
	for i := 1; i <= maxTaskCounter; i++ {
		taskList = append(taskList, Task{i})
	}

	// Number of current running concurrent tasks
	taskCounterChan := make(chan struct{}, maxTaskCounter)
	defer close(taskCounterChan)

	// Channel to block writes to log file
	logChan := make(chan struct{})
	defer close(logChan)

	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	tickerFlush := time.NewTicker(1 * time.Second)
	defer tickerFlush.Stop()

	tickerRotate := time.NewTicker(60 * time.Second)
	defer tickerRotate.Stop()

	for _, task := range taskList {
		logger.Printf(logFmt, "Enqueue task", task.ID, time.Now())
		taskCounterChan <- struct{}{}
		go task.WriteLog(logger, logChan, taskCounterChan)
	}

	log.Println("All tasks enqueued...")

loop:
	for {
		select {

		case <-tickerFlush.C:
			// flush the buffer, write to log file
			logW.Write(&buf)
			log.Println("Flush buffer to file...") // DEBUG

		case <-tickerRotate.C:
			// flush the buffer, write to log file, close log file
			logW.Write(&buf)
			logW.Close()
			log.Println("Rotate file...") // DEBUG

		case <-logChan:
			logW.Write(&buf)
			log.Println("Task done...") // DEBUG

		default:
			// Wait for all goroutines to finish
			if len(taskCounterChan) == 0 {
				break loop
			}
		}
	}

	log.Println("All tasks finished...")
}
