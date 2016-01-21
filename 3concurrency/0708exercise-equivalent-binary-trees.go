package main

import "golang.org/x/tour/tree"
import "container/list"

// type Tree struct {
// 	Left *Tree
// 	Value int
// 	Right *Tree
// }

func Walk(t *tree.Tree, ch chan int) {
	target := t
	stacked := list.New()

	for {
		if target.Left != nil {
			stacked.PushFront(target)
			target = target.Left
		} else {
			ch <- target.Value
			target = stacked.Remove(stacked.Front())

			ch <- target.Value
			target = tmp.Right
		}
	}
}

// func Same(t1, t2 *tree.Tree) bool

func main() {
	var v int
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for {
		select {
		case v <- ch:
			fmt.Println(v)
		default:
			return
		}
	}
}
