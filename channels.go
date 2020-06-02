package main

import "fmt"

func main() {
	s := []int{1, 3, 5, 7, 9}

	c := make(chan int)
	go sum(s, c)
	x := <-c
	fmt.Println(x)
}

func sum(s []int, c chan<- int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
