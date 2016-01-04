package main

import (
	"fmt"
)

// 戻り値としてクロージャの関数値(func(int) int)を返す関数の定義
func adder() func(int) int {
	// クロージャと下記のsum変数がバインド(bind)される
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// クロージャを関数値として設定
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
