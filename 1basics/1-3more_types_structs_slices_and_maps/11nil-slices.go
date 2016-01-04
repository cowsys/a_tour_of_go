package main

import (
	"fmt"
)

func main() {
	var z []int
	// スライスのゼロ値はnilであり
	// nilのスライスのlen, capはともに0
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nill!")
	}
}
