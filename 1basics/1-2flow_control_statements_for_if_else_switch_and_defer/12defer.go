package main

import (
	"fmt"
)

func main() {
	// deferに渡した関数の実行は
	// 呼び出し元の関数の終わり(return)まで遅延する

	// NOTE:deferに渡した関数の引数は
	//      ***すぐに評価***される
	defer fmt.Println("world")

	fmt.Println("hello")
}
