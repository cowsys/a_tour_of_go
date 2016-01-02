package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// NOTE:deferで登録した関数の実行は
		// stack(LIFO)として再実行される
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
