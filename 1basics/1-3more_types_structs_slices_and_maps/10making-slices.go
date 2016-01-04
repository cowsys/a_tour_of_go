package main

import (
	"fmt"
)

func main() {
	// スライスはmakeで生成
	// ゼロ値に初期化した配列をメモリに割り当て
	// その配列への参照を持つスライスを返す
	a := make([]int, 5)
	printSlice("a", a)

	// makeの第三引数には
	// スライスの容量(capacity)を指定
	b := make([]int, 0, 5)
	printSlice("b", b)

	// NOTE:空のスライスの再スライスは
	//      ゼロ値が設定されたうえで
	//      再スライスされる

	//      また以下の例からも
	//      capは保持されて再スライスされる
	c1 := b[:2]
	c2 := b[:1]
	printSlice("c1", c1)
	printSlice("c2", c2)

	// 設定されているlen以上の要素を再スライス
	// したときもゼロ値が設定される

	// NOTE:ただしこのときはcapが3となり、
	//      capを保持してのコピーとならない
	d := c1[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf(
		"%s len=%d cap=%d %v\n",
		//         cap(x)でスライスの容量を返す
		s, len(x), cap(x), x,
	)

}
