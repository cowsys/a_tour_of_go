package main

import (
	"fmt"
)

func fibonacci() func(int) int {
	bn := [2]int{0, 1}

	return func(i int) int {
		if i < 2 {
			return i
		} else {
			ret := bn[0] + bn[1]

			bn[0] = bn[1]
			bn[1] = ret
			// fmt.Println(bn)
			return ret
		}
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 50; i++ {
		fmt.Println(f(i))
	}
}
