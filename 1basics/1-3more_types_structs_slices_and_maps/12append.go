package main

import (
	"fmt"
)

func main() {
	var a []int
	printSlice("a", a)

	// スライスに新しい要素を追加するには
	// appendを利用する
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	// ここでcapが6になる
	// 5とならないのはおそらくデータ構造上の理由
	printSlice("a", a)

	a = append(a, 5)
	// ここではcapは6
	printSlice("a", a)

	a = append(a, 6)
	// ここではcapは12になる
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x,
	)
}
