package main

import (
	"fmt"
	"math"
)

// インタフェース型はメソッドの集まりで定義
type Abser interface {
	Abs() float64
}

func main() {
	// インタフェース型の変数Abserを定義
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// NOTE:Abserインタフェースを実装
	// AbserにMyFloatのAbs()をバインドするイメージ
	a = f
	fmt.Println(a.Abs())

	// NOTE:Abserインタフェースを実装
	// AbserにVertexのAbs()をバインドするイメージ
	a = &v
	fmt.Println(a.Abs())

	// NOTE:これはエラーになる
	//      Abs()メソッドのメソッドレシーバは*Vertexであり
	//      VertexがAbserインタフェースを満たしていないため
	// a = v
	// fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
