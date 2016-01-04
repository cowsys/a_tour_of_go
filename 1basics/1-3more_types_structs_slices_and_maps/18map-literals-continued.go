package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// mapリテラルは以下のように設定可能
var m = map[string]Vertex{
	//           明示的に値の型(Vertex)を定義しなくても
	//           リテラルの内容から型推論で適切な型名を設定できる
	//           NOTE:***ただしmapに渡すトップレベルの方が単純な型名である場合のみ***
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
