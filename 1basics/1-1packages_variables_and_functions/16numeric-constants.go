package main

import (
	"fmt"
)

// constもまとめて定義可能
const (
	// NOTE:数値の定数は
	//      ***高精度な値(values)***
	//      intで足りない時などに利用
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// NOTE:型のない定数はその状況によって
	//      必要な型をとる

	// intとしても利用できる
	fmt.Println(needInt(Small))

	// float64としても利用できる
	fmt.Println(needFloat(Small))

	// float64としても利用できる
	fmt.Println(needFloat(Big))

	// NOTE:以下はオーバフローしてエラー
	// ret->./16numeric-constants.go:35:
	// constant 1267650600228229401496703205376 overflows int
	// fmt.Println(needInt(Big))

	// NOTE:ただし柔軟性を持っているのは
	//      数値系のみの様子
	// const STR = "xxx"
	// compile errorとなる
	// fmt.Println(needInt(STR))

}
