package main

import (
	"fmt"
)

func main() {
	sum := 1
	// セミコロンは省略できる
	// whileはgoにはない
	// forを活用し記述する
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
