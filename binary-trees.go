package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func main() {
	c := make(chan int)
	t := tree.New(1)
	go Walk(t, c)
	for i := range c {
		fmt.Println(i)
	}
}
