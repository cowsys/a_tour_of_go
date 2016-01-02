package main

import (
	"fmt"
)

func main() {
	sum := 1
	// 以下のように書くとfmt/lintで潰される
	// for ; sum < 1000; {
	// 「初期化と後処理ステートメントの記述は任意」
	// だが、ない方がベターという思想
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
