package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 要素の削除はdeleteを用いる
	delete(m, "Answer")
	// deleteした要素にアクセスするとゼロ値が値として帰る
	fmt.Println("The value:", m["Answer"])

	// deleteした要素にアクセスしても警告は発生しないが
	// 以下のように空のMapとなっている
	fmt.Println("The value:", m)

	// v: Mapの値
	// ok:指定したキーの値が存在するかどうか
	// を返す
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
