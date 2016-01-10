package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

//         Error()メソッドを実装したMyErrorは
//         error型インタフェースの戻り値として扱うことができる
func run() error {
	return &MyError{
		time.Now(),
		"it did'nt work",
	}
}

func main() {
	if err := run(); err != nil {
		// NOTE:fmt.stringerと同様に、fmtパッケージは変数を文字列で出力する際に
		//      errorインタフェースを確認する
		fmt.Println(err)
	}
}
