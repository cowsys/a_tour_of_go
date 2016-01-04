package main

import (
	"fmt"
)

// 構造体の定義
type Vertex struct {
	X int
	Y int
}

func main() {
	// 構造体をPrintlnに食わせると
	// 構造体の内容をいい感じに表示
	fmt.Println(Vertex{1, 2})
}
