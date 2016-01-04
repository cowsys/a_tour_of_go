package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// key:string, value:VertexのMap変数を定義
var m map[string]Vertex

func main() {
	// Mapはmake関数で作成する
	m = make(map[string]Vertex)
	fmt.Println(m)

	// NOTE:上述のvar m map...を定義せずに以下でもOK
	// m := make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}
