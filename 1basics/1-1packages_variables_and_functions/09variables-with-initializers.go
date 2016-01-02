package main

import (
	"fmt"
)

// 変数の宣言とともに初期値をセット
// "初期化子(initializer)を与える"という
var i, j int = 1, 2

func main() {
	// 型を指定していないが、
	// 初期化子の内容を評価し適切な型が利用される
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
