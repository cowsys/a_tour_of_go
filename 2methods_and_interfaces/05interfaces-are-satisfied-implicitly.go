package main

import (
	"fmt"
	"os"
)

// type Reader interface {
// 	Read(b []byte) (n int, err error)
// }

type Writer interface {
	Write(b []byte) (n int, err error)
}

// type ReadWriter interface {
// 	// インタフェースの内容として
// 	// インタフェースを定義することも可能
// 	Reader
// 	Writer
// }

func main() {
	var w Writer

	w = os.Stdout
	// // Stdin, Stdout, and Stderr are open Files pointing to the standard input,
	// // standard output, and standard error file descriptors.
	// var (
	// 	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	// 	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	// 	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	// )

	fmt.Fprintf(w, "hello writer\n")
	//// NOTE:Fprintfの定義は以下の様なかたち
	// // Fprintf formats according to a format specifier and writes to w.
	// // It returns the number of bytes written and any write error encountered.
	// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {

	//// NOTE:io.Writerはこの例で定義したWriterと同じインタフェースを持っている
}
