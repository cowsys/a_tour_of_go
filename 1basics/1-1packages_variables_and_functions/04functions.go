package main

import "fmt"

//         変数名の後ろに型を書く
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
