package main

import (
	"fmt"
	"math"
)

// ビルトインの型も型として新たに定義し
// メソッドレシーバ付き関数を定義することで
// 型を拡張することができる
// NOTE:が、有効なのは***パッケージ内***の型のみ
//      他のパッケージのものにはできない
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	f2 := MyFloat(math.Sqrt2)
	fmt.Println(f2.Abs())
	// Sqrt2   =
	//   1.41421356237309504880168872420969807856967187537694807317667974
	// http://oeis.org/A002193
}
