package main

import (
	// "fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)

	sa := strings.Fields(s)
	for i := 0; i < len(sa); i++ {
		// fmt.Println(sa[i])
		ret[sa[i]]++
	}
	return ret
}
func main() {
	wc.Test(WordCount)
}
