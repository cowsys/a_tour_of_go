package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	//func Intn(n int) int
	//	Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	//	from the default Source. It panics if n <= 0.
}
