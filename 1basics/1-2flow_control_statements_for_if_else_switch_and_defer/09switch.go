package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	//// NOTE:runtime.GOOS
	// const GOOS string = theGoos
	//     GOOS is the running program's operating system target: one of darwin,
	//     freebsd, linux, and so on.

	// NOTE:switch文でも条件の前に評価するための簡単なステートメントを書くことが可能
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,...
		fmt.Printf("%s", os)
	}

	// 例によってosはswithchスコープでのみ有効
	// fmt.Println(os)
}
