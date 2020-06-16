package main

import (
	"fmt"
)

func main() {
	done := make(chan int)
	defer close(done)

	go task(done)
	done <- 3
	done <- 4
	fmt.Println("end")
}

func task(done <-chan int) {
	fmt.Println(<-done)
	<-done
}
