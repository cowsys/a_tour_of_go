package main

import (
	"fmt"
)

func main() {
	// 要素2のstring型配列を定義
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 確保された要素数以上を利用するとエラー
	// a[2] = "World"
	// ret->./06array.go:18: invalid array index 2 (out of bounds for 2-element array)

	fmt.Println(len(a))
}
