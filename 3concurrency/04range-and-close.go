package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// channelの明示的なクローズ
	close(c)
	// 以下のようにすることでokでチャネルのクローズ状態を受け取れる
	// v, ok :=  <-ch

	// closeしたチャネルに送信するとpanic
	// c <- 5

	// NOTE:上記例ではチャネルをcloseしているが、
	// 通常はcloseする必要はない
	// 送信が終了したことを明示的に示したい場合のみ
	// closeする
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// channelがcloseするまで受信した値を表示するループ
	for i := range c {
		fmt.Println(i)
	}
}
