package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's saturday?")

	today := time.Now().Weekday()
	fmt.Printf("(today:%s(integer expression: %d))\n", today, today)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
		// NOTE:多言語のように明示的にbreakしなくても
		//      該当するcaseの評価後case分を抜ける
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
