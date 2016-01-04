package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// structのフィールドはstructのポインタを通じてアクセスする
	// ポインタを通じたindirectionは透過的
	p.X = 1e9

	// 以下のようにもアクセスできる
	// NOTE:違いは何？
	v.Y = 100

	fmt.Println(v)
}
