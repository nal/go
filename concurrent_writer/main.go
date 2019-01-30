package main

import (
	"bytes"
	"fmt"
	"log"
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
// func (myW *MyWriter) Write(w io.Writer) {
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
func (rcv Task) WriteLog(logger *log.Logger, logChan chan<- struct{}) {
	logger.Printf(logFmt, "Started task", rcv.ID, time.Now())
	// Sleep for random time to test log file rotation.
	// r := rand.New(rand.NewSource(130))

	// r := rand.Intn(120)
	// sleepSeconds := time.Duration(int64(r))
	// log.Printf("Sleep for %d seconds!", sleepSeconds)
	// time.Sleep(sleepSeconds * time.Second)
	time.Sleep(130 * time.Second)
	logger.Printf(logFmt, "Finished task", rcv.ID, time.Now())
	logChan <- struct{}{}
}

// Write access to global variables within goroutines is thread unsafe.
// Create a channel to communicate in gorounites and control number
// of already running goroutines.
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

	// Channel to write lot log file
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
		go task.WriteLog(logger, logChan)
	}

loop:
	for {
		select {
		case <-tickerFlush.C:
			// flush the buffer, write to log file
			logW.Write(&buf)
			log.Println("Flush!")

		case <-tickerRotate.C:
			// flush the buffer, write to log file, close log file
			logW.Write(&buf)
			logW.Close()
			log.Println("Rotate!")

		case <-logChan:
			// Wait for task to complete
			log.Println("Task done!")
			break loop
		}
	}

	logW.Write(&buf)
}
