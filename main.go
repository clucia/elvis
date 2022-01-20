package main

import (
	"fmt"
	"io"
	"os"
)

// These functions are to remain unchanged
func getByte(r io.Reader) (b byte, err error) {
	buf := make([]byte, 1)
	_, err = r.Read(buf)
	b = buf[0]
	return b, err
}

func putByte(w io.Writer, b byte) (err error) {
	buf := []byte{b}
	_, err = w.Write(buf)
	return err
}

func check(err error) {
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(-1)
	}
}

// This program copies from inputStream to outputStream by reading and writing
// bytes through the provided interface.
//
// A new requirement has come up:
//    Any place that "elvis" appears in the stream, it must be replaced with "Elvis"
//    Everything else must be passed through unchanged.
//    The input stream may be arbitrarily long and free of delimiters, so the solution
//    must operate on streams.
//
// Please do not use any library functions unless excepted below:
//   bytes.Compare:
//     func Compare(a, b []byte) int
//
// The assigment will involve modifying/replacing the loop from line 61 to 67

func main() {
	var err error

	inputStream, err := os.Open("input.txt")
	check(err)

	// inputStream := os.Stdin

	searchBytes := []byte{'e', 'l', 'v', 'i', 's'}
	replaceBytes := []byte{'E', 'l', 'v', 'i', 's'}
	_ = searchBytes
	_ = replaceBytes

	outputStream := os.Stdout

	var b byte
	for err == nil {
		b, err = getByte(inputStream) // returns io.EOF at end of stream
		if err != nil {
			break
		}
		err = putByte(outputStream, b)
	}
	if err == io.EOF {
		fmt.Println("SUCCESS")
	}
}
