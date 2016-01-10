package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// type Reader struct {
	// 	// contains filtered or unexported fields
	// }
	// 	A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
	// 	io.ByteScanner, and io.RuneScanner interfaces by reading from a string.
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		//// NOTE:io.Reader
		// type Reader interface {
		// 	Read(p []byte) (n int, err error)
		// }
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}

	}
}
