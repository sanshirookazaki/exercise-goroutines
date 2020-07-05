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

	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
