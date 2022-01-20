package main

import (
	"fmt"
	"io"
	"os"
)

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

func main() {
	var err error

	inputStream, err := os.Open("input.txt")
	check(err)

	// inputStream := os.Stdin

	outputStream := os.Stdout
	var b byte
	b, err = getByte(inputStream) // returns io.EOF at end of stream

	for err == nil {
		err = putByte(outputStream, b)
		check(err)
		b, err = getByte(inputStream)
	}
	if err == io.EOF {
		fmt.Println("SUCCESS")
	}
}
