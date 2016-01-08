package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// *Vertexに対するメソッドレシーバを設定
// *Vertexはこの関数を利用することができる
func (v *Vertex) Abs() float64 {
	// 自分自身はメソッドレシーバ(v *Vertex)に定義した
	// "v"で参照することができる
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}

	// 引数はないが、関数内で自身を自己参照している
	fmt.Println(v.Abs())
}
