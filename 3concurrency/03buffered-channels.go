package main

import (
	"fmt"
)

func main() {
	// buffered channelの生成
	// 第二引数にバッファの長さを与える
	ch := make(chan int, 2)

	// チャネルへの送信
	ch <- 1
	ch <- 2

	// チャネルからの受信
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// NOTE:以下の例は
	// バッファが空の状態で受信しようとするため
	// fatalとなる
	// fmt.Println(<-ch)

	// NOTE:以下の例は
	// バッファが詰まった状態でさらに送信しようとするため
	// fatalとなる
	// ch <- 1
	// ch <- 2
	// ch <- 3
}
