package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Sqrt(x float64) float64 {
	const ACCEPTABLE_DIFFERENCE = 0.01
	var beforeZ float64

	z := float64(rand.Intn(10))
	for {
		beforeZ = z
		z = z - ((math.Pow(z, 2) - x) / 2 * z)
		// fmt.Println(z)

		// NOTE:math.Absで絶対値を取得
		if math.Abs(beforeZ-z) < ACCEPTABLE_DIFFERENCE {
			break
		}
	}

	return z
}

func main() {
	fmt.Println("final output:", Sqrt(2))
	fmt.Println("final output(origin):", math.Sqrt(2))
}
