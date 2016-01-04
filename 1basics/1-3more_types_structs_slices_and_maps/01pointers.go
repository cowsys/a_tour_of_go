package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	// NOTE:ポインタのゼロ値はnil
	// NOTE:***ポインタの演算はない***
	// &{varname}で{varname}のポインタを取得
	p := &i
	// *{pointer}でポインタの指す要素を参照
	// 用語としてdereferencing/indirectingという
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
