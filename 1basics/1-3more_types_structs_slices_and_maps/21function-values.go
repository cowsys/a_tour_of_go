package main

import (
	"fmt"
	"math"
)

// 関数の引数として関数を渡すことも可能
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 関数定義も変数に格納・利用できる
	// そうじて"関数値(function value)という"
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
