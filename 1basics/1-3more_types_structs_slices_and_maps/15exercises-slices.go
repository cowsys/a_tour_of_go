package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// サイズを指定した初期化スライスを定義したいときはmake
	// リテラルを設定するときはmakeを使わないイメージ
	s := make([][]uint8, dy)

	for i := 0; i < len(s); i++ {
		s[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			// s[i][j] = uint8((i + j) / 2)
			// s[i][j] = uint8(i * j)
			// s[i][j] = uint8(i^j)
			// s[i][j] = uint8(((i + j) / 2) * (i * j) * (i^j))
		}
	}

	return s
}

func main() {
	// なぜかPicの引数を定義しなくても動く
	// picパッケージ内でなにかやっている模様
	pic.Show(Pic)
}
