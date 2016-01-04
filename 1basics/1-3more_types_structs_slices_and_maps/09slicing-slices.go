package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	// 以下の例は1 - (4-1):1 - 3の要素を返す
	fmt.Println("s[1:4] ==", s[1:4])

	// 以下の例は0 - (3-1):0 - 2の要素を返す
	fmt.Println("s[:3] ==", s[:3])

	// 以下の例は4 - lastの要素を返す
	fmt.Println("s[4:] ==", s[4:])
}
