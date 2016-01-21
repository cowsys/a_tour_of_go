package main

import (
	"fmt"
)

// channelをつかって処理結果をやり取りするため
// 戻り値がない
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	// intを受け付けるchannelをmakeで生成
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// channelから値を受け取ったら次の行の処理が完了
	// NOTE:どちらの値が先に割り当てられるかは
	// goランタイム上で稼働するgoroutine次第
	// この例のように単純にやる場合は制御出来ない様子
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// NOTE:
	// 通常送信と受信の片方が準備できるまで
	// 送受信はブロックされる
	// これによって明確なロックや条件変数がなくても
	// goroutineの同期を可能にする
}
