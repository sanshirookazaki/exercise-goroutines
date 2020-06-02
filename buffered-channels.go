package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go ten(c)
	for v := range c {
		fmt.Println(v)
	}
}

func ten(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
