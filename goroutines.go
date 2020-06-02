package main

import (
	"fmt"
	"time"
)

func main() {
	go say("green")
	say("red")
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

func say(s string) {
	fmt.Println(s)
}
