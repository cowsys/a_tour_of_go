package main

import (
	"fmt"
)

//            2つ以上の引数が同じ型の時は最後にだけ型を記述
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
