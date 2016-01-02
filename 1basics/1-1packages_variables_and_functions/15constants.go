package main

import (
	"fmt"
)

// 定数の定義
const Pi = 3.14

//// NOTE:定数は以下のみで利用可能
// 文字(character)
// 文字列(string)
// boolean
// 数値(numeric)

//// NOTE:定数には:= を使えない

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
