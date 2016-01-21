package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// selectは複数あるcaseのいずれかが準備できるようになるまでブロックし
		// 準備ができたcaseを実行する

		// NOTE:
		// *** もし複数のcaseが準備出来ている場合は、caseはランダムに選択される ***
		select {
		//   A:送信
		case c <- x:
			x, y = y, x+y
		//   A:受信
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//        B:受信
			fmt.Println(<-c)
		}
		//   B:送信
		quit <- 0
	}()

	fibonacci(c, quit)
}
