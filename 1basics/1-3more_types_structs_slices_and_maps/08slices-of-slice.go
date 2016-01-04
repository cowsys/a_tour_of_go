package main

import (
	"fmt"
	"strings"
)

func main() {
	// 多重構造のスライスの定義
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	game[0][0] = "X"
	game[0][2] = "X"

	game[1][0] = "O"

	game[2][0] = "X"
	game[2][2] = "O"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
		// NOTE:さらにJoinの第二引数には
		//      セパレータ文字列を付与できる
		//      phpのimplode相当のものが実行可能
		// fmt.Printf("%s\n", strings.Join(s[i], "@"))
	}
}
