package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// NOTE:ifステートメントでは条件の前に評価するための簡単なステートメントを書くことが可能
	// ここで宣言された変数はifのスコープ内だけで有効
	// Pow returns x**y, the base-x exponential of y.
	if v := math.Pow(x, n); v < lim {
		return v
	}

	// NOTE:if内でのみ有効なvを利用すると以下のエラーが発生
	// ./06if-with-a-short-statement.go:16: undefined: v
	// return v
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
