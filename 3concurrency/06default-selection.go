package main

import (
	"fmt"
	"time"
)

func main() {
	// Tick is a convenience wrapper for NewTicker providing access to the ticking
	// channel only. While Tick is useful for clients that have no need to shut down
	// the Ticker, be aware that without a way to shut it down the underlying
	// Ticker cannot be recovered by the garbage collector; it "leaks".
	tick := time.Tick(100 * time.Millisecond)

	// After waits for the duration to elapse and then sends the current time
	// on the returned channel.
	// It is equivalent to NewTimer(d).C.
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!.")
			return
		default:
			// 処理中はブロッキングしてしまう点に注意
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
