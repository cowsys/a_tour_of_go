package main

import (
	"fmt"
	"math/cmplx"
)

// 複数の変数の定義
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//// goの型一覧
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // uint8 の別名
//
// rune // int32 の別名
//      // unicode のコードポイントを表す
//
// float32 float64
//
// complex64 complex128

func main() {
	// constは変数と同じフォーマットで定義
	// %T:型
	// %v:値(型によらず柔軟にやってくれる)
	const f = "%T(%v)\n"

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
