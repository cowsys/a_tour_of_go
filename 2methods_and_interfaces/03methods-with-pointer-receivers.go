package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// メソッドは名前付きの型(named type)
// もしくは
// 名前付きの型へのポインタ
// に割り当てることが可能

// この関数ではVertexの***ポインタ型(poointer type)***
// に割り当てている(ポインタのメソッドレシーバ)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 値のメソッドレシーバの定義...しようと思ったが、
// method redeclaredとなりできなかった
// どうやらポインタ・値のメソッドレシーバは共存できない様子
// func (v Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)
	fmt.Printf("After scaling: %+v,  Abs: %v\n", v, v.Abs())
}
