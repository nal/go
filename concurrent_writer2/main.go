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

// LogWriter interface implemens io.Writer interface
type LogWriter interface {
	Write(p []byte) (n int, err error)
	openFile()
	closeFile()
	rotateFile()
}

// LogFile struct
type LogFile struct {
	filePath string
	fd       *os.File
}

// Implementing io.Writer interface
// Write to underlying data stream.
func (logF *LogFile) Write(p []byte) (n int, err error) {
	// File is closed or rotating
	if logF.fd == nil {
		logF.openFile()
	}

	// Write to log file and return the number of bytes written from p (0 <= n <= len(p))
	// and any error encountered
	n, err = logF.fd.Write(p)
	if n < len(p) {
		return n, err
	}

	logF.fd.Sync()
	return n, nil
}

// openFile opens log file. Panic if operation failed.
func (logF *LogFile) openFile() {
	ctime := time.Now()
	hour := ctime.Hour()
	min := ctime.Minute()
	logF.filePath = fmt.Sprintf("tmp_%02d%02d.txt", hour, min)

	f, err := os.OpenFile(logF.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Open or rotate file %s: ", logF.filePath)

	// Store new file handler
	logF.fd = f
}

// closeFile closes log file.
func (logF *LogFile) closeFile() {
	// File is opened
	if logF.fd != nil {
		// Close file
		logF.fd.Sync()

		if err := logF.fd.Close(); err != nil {
			log.Println(err)
		} else {
			// Drop file handler or else file will not reopen!
			logF.fd = nil
		}
	}
}

// rotateFile rotates log file.
func (logF *LogFile) rotateFile() {
	logF.closeFile()
	logF.openFile()
}

// Run Task method. Do nothing useful, just fill buffer with debug data.
func (rcv *Task) Run(logChan chan<- bytes.Buffer, taskCounterChan <-chan struct{}) {
	var buf bytes.Buffer
	log.Printf(logFmt, "Started task", rcv.ID, time.Now())
	buf.Write([]byte(fmt.Sprintf(logFmt, "Started task", rcv.ID, time.Now())))

	// Sleep for random time more than 60 seconds to test log file rotation.
	sleepSeconds := time.Duration(rand.Int31n(60)+20) * time.Second

	log.Printf("%-15s %2d sleep for %v\n", "Started task", rcv.ID, sleepSeconds)
	buf.Write(([]byte(fmt.Sprintf("%-15s %2d sleep for %v\n", "Started task", rcv.ID, sleepSeconds))))

	time.Sleep(sleepSeconds)

	buf.Write(([]byte(fmt.Sprintf(logFmt, "Finished task", rcv.ID, time.Now()))))

	// Block channel and notify we have filled a buffer
	logChan <- buf

	// Decrease number of worker goroutines
	<-taskCounterChan
}

// Solution: main function spawns worker goroutines with `LogWriter.Write` method.
// These methods write data to `*bytes.Buffer` and blocks.
// In `for-select-case` block we read data from 3 channels:
// * rotate file timer (to do the same as flush file timer plus close file handler)
// * log channel (write to buffer and write it to log file)
// * worker count channel (abort execution, all workers finished)
func main() {
	// Init empty LogWriter
	logF := LogFile{filePath: "", fd: nil}
	defer logF.closeFile()

	// Array of Task structures
	var taskList = []*Task{}

	// Maximum number of tasks
	maxTaskCounter := 200

	// Populate task array
	for i := 1; i <= maxTaskCounter; i++ {
		taskList = append(taskList, &Task{i})
	}

	// Number of current running concurrent tasks
	taskCounterChan := make(chan struct{}, maxTaskCounter)
	defer close(taskCounterChan)

	// Channel to block writes to log file
	logChan := make(chan bytes.Buffer)
	defer close(logChan)

	// Shared buffer across goroutines
	// TODO: use "sync/atomic" package to "lock" filename across goroutines

	// Rotate log file ticker channel
	tickerRotate := time.NewTicker(60 * time.Second)
	defer tickerRotate.Stop()

	// Run tasks in goroutine
	for _, task := range taskList {
		taskCounterChan <- struct{}{}
		go task.Run(logChan, taskCounterChan)
	}

	log.Println("All tasks enqueued...")

loop:
	for {
		select {

		case <-tickerRotate.C:
			logF.rotateFile()
			log.Println("Rotate file...")

		case p := <-logChan:
			_, err := logF.Write(p.Bytes())
			if err != nil {
				log.Fatalf("Failed to write to log file: %s", err)

			} else {
				p.Reset()
				log.Println("Task done...")
			}

		default:
			// Wait for all goroutines to finish
			if len(taskCounterChan) == 0 {
				break loop
			}
		}
	}

	log.Println("All tasks finished...")
}
