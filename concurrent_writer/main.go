package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// MyWriter struct
type MyWriter struct {
	filePath string
	fd       *os.File
}

// Create and write to file hander.
// func (myW *MyWriter) Write(w io.Writer) {
func (myW *MyWriter) Write() {
	ctime := time.Now()
	hour := ctime.Hour()
	min := ctime.Minute()

	// File is closed or rotating
	if myW.fd == nil {
		myW.filePath = fmt.Sprintf("tmp_%02d%02d.txt", hour, min)

		f, err := os.OpenFile(myW.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		// Store new file handler
		myW.fd = f
	}

	// In real life we read something here and write to file.
	// Now just write currrent date and time.
	b := []byte(fmt.Sprintf("Now: %s\n", ctime.Format(time.RFC3339)))
	myW.fd.Write(b)
}

// Close filehandler
func (myW *MyWriter) Close() {
	// File is opened
	if myW.fd != nil {
		if err := myW.fd.Close(); err != nil {
			log.Println(err)
		}
	}
}

// Rotate filehandler
func (myW *MyWriter) Rotate() {
	myW.Close()
	time.Sleep(5 * time.Second)
}

func main() {
	myW := MyWriter{filePath: "", fd: nil}
	defer myW.Close()

	go func() {
		for {
			myW.Rotate()
		}
	}()

	myW.Write()
	time.Sleep(30 * time.Second)
	myW.Write()
	time.Sleep(30 * time.Second)
	myW.Write()
	time.Sleep(30 * time.Second)
}
