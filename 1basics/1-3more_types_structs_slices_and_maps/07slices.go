package main

import (
	"fmt"
)

func main() {
	// 構文の形は似ているがこれは配列ではなくスライス
	// スライスの初期値を設定している
	// スライスの場合は長さを保持している&拡張性があるため
	// 要素数は空でもOK
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	// NOTE:len()の定義は:GoDoc builtinで確認可能
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}
