package main

import (
	"fmt"
)

func main() {
	var i, j int = 1, 2
	// ***関数の中では***
	// var宣言の代わりに短い":="の代入文を使い
	// 暗黙的な型宣言が可能

	// ***関数の外では***
	// キーワードで始まる宣言(var, funcなど)が必要で
	// :=での暗黙的な宣言はできない
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
