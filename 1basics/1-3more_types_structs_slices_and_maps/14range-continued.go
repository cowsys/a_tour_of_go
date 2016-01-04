package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)

	for i := range pow {
		// "<<"はシフト演算子。左にシフト
		pow[i] = 1 << uint(i)
	}

	// 利用しない値は_で受けることができる
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
