package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//// NOTE:Time.Hour()
	// // Hour returns the hour within the day specified by t, in the range [0, 23].
	// func (t Time) Hour() int {
	// 	return int(t.abs()%secondsPerDay) / secondsPerHour
	// }

	// 条件のないswitchは分岐が多岐に渡るif文をすっきり書ける
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
