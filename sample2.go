package main

import "log"

func main() {
	resc := make(chan int)

	go func(resc chan int) {
		defer close(resc)
		for i := 0; i < 2; i++ {
			log.Println("sub:", i)
			resc <- i
		}
	}(resc)

	for i := range resc {
		log.Println("main:", i)
	}
}
