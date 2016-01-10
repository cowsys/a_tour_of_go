package main

import (
	"fmt"
	"math"
	"math/rand"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", fmt.Sprint(float64(e)))
}

func Sqrt(x float64) (float64, error) {
	const ACCEPTABLE_DIFFERENCE = 0.01
	var beforeZ float64

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	// 負の値の場合はSqrtの戻り値として
	// errorインタフェースを満たしたオブジェクトが返る
	// 結果、PrintlnでError()メソッドがコールされ評価される
	fmt.Println(Sqrt(-2))
}
