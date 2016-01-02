package main

import (
	"fmt"
)

// 変数の宣言のみ行う場合は
// varを付与する
// 値はセットされていないが、デフォルト値として
// 型に応じたデフォルト値がセットされている
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
