package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("A")
	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("B")
	}()

	wait.Wait()
}
