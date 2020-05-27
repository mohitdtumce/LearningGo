package main

import (
	"bytes"
	"fmt"
)

// ==================================================
// Example 1

// Writer Interface
// Interfaces doesn't describe data rather behaviour so there will not be any data fields present. Instead there will be method signatures.
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter will implement Writer
// Implicit implementation of interface can be done using structs
type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// ==================================================
// Example 2

// Incrementer Interface
type Incrementer interface {
	Increment() int
}

// IntCounter of type int
type IntCounter int

// Increment method implemented
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// ==================================================
// Example 3: Composing Interfaces

// Writr Interface
type Writr interface {
	Write([]byte) (int, error)
}

// Closr Interface
type Closr interface {
	Close() error
}

// WritrClosr Interface
type WritrClosr interface {
	// Embedding interfaces into one another
	Writr
	Closr
}

// BufferedWritrClosr Struct
type BufferedWritrClosr struct {
	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	// A Buffer needs no initialization
	buffer *bytes.Buffer
}

// Write Method Implementation
func (bwc *BufferedWritrClosr) Write(data []byte) (int, error) {
	/*
		func (b *Buffer) Write(p []byte) (n int, err error)

		Write appends the contents of p to the buffer, growing the buffer as needed. The return value n is the length of p; err is always nil.
		If the buffer becomes too large, Write will panic with ErrTooLarge.
	*/
	n, err := (*bwc).buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	/*
		func (b *Buffer) Len() int

		Len returns the number of bytes of the unread portion of the buffer;
		b.Len() == len(b.Bytes()).
	*/
	for (*bwc).buffer.Len() > 8 {
		/*
			func (b *Buffer) Read(p []byte) (n int, err error)

			Read: reads the next len(p) bytes from the buffer or until the buffer is drained. The return value n is the number of bytes read.
			If the buffer has no data to return, err is io.EOF (unless len(p) is zero); otherwise it is nil.
		*/
		_, err := (*bwc).buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// Close Method Implementation
func (bwc *BufferedWritrClosr) Close() error {
	fmt.Println("Length of buffer", (*bwc).buffer.Len())
	for (*bwc).buffer.Len() > 0 {
		/*
			func (b *Buffer) Next(n int) []byte

			Next returns a slice containing the next n bytes from the buffer, advancing the buffer as if the bytes had been returned by Read.
			If there are fewer than n bytes in the buffer, Next returns the entire buffer. The slice is only valid until the next call to a read or write method.
		*/
		data := (*bwc).buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// NewBufferedWritrClosr Method 
func NewBufferedWritrClosr() *BufferedWritrClosr {
	return &BufferedWritrClosr{
		buffer: bytes.NewBuffer([]byte{}),
	}
	
}

func main() {
	// ==================================================
	// 12.0 Interface -
	/*
		1. Basics,
		2. Composing Interfaces,
		3. Type Conversion,
			3.1 The Empty Interface
			3.2 Type Switches
		4. Implementing with value vs interface
		5. Best Practices
	*/

	// Example 1: Inplementation of interface using struct
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))

	// Example 2:
	var myInt IntCounter = 0
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
	fmt.Println(inc)

	// Example 3: Composing Interfaces
	var wc WritrClosr = NewBufferedWritrClosr()
	wc.Write([]byte("I Learnt About Buffer Package. Nice Hurray !!"))
	wc.Close()

	// Type Conversion interface.(type)
	bwc := wc.(*BufferedWritrClosr)
	fmt.Println(bwc)

}
