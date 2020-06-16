package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	q := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		c <- 0
		time.Sleep(1 * time.Second)
		q <- 0
	}()

	stay(c, q)

	fmt.Println("end")
}

func stay(c, q chan int) {
	for {
		select {
		case <-c:
			fmt.Printf("%T \n", c)
		case <-q:
			fmt.Println("q")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("default")
		}
	}
}
