package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	//structの初期値割当
	// 一部分だけを行うときは"Name:"構文を使う
	// 以下の例だとフィールドXのみを設定している
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
