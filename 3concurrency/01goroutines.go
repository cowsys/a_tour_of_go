package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	// goroutineは同じアドレス空間で実行されるため
	// 共有メモリへのアクセスは必ず同期する必要がある
	s := "hogehoge"
	go say(s)
	s = "fugafuga"
	// 以下の結果はfugafugaが表示される
	// 上記のgo says(s)は引数を評価したうえで
	// goroutineとしてsayを実行しているため
	// 変数の変更の影響を受けない
	say(s)
}
