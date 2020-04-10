package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// This is the custom struct from our package
type Person struct {
	Id   int
	Name string
	Age  int
}

// The method takes just a io.Writer as input
func (p *Person) Write(w io.Writer) {
	b, _ := json.Marshal(*p)
	// Inside our function we just write into the io.Writer
	// We don't care about which writer we use
	w.Write(b)
}

func main() {
	me := Person{
		Id:   1,
		Name: "Me",
		Age:  64}

	// For testing we use a bytes.Buffer
	// That type has an implementaion of the io.Writer
	var b bytes.Buffer

	// We call the write Method with our io.Writer
	// For testing we don't write our data into a
	// file. Just into a buffer.
	me.Write(&b)

	// Let's look what the Write method wrote into our
	// buffer
	fmt.Printf("%s\n", b.String())
}
