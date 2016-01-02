package main

import (
	"fmt"
)

//                   関数の戻り値として型だけでなく変数名も指定している
//                   named return valueという
func split(sum int) (x, y int) {
	// named return valueとなっているため、:=としなくてよい
	x = sum * 4 / 9
	y = sum - x

	// 関数の戻り値の定義で変数名を指定していた場合
	// return に具体的な変数名を指定しなくてもよい
	// naked returnという

	// naked returnは短い関数で利用すべき
	// 頭のなかで覚えておける範疇で利用する
	return
}

func main() {
	fmt.Println(split(17))
}
