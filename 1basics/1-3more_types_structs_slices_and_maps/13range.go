package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// スライスはrange関数を使ってループさせる
	// iはループの順番が、
	// vはスライスの値が入る
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
