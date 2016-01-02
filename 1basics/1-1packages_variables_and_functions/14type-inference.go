package main

import (
	"fmt"
)

func main() {
	// 型推論によって初期値がセット
	v := 42.399 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	// ret->v is of type complex128

	v = 5
	fmt.Printf("v is of type %T\n", v)
	// ret->v is of type complex128
	// NOTE:一度設定した変数の型はそのまま

	w := 42.399
	fmt.Printf("v is of type %T\n", w)
	// ret->v is of type float64

}
