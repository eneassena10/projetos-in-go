package main

import (
	"log"
	"time"
)

func writeToChannel(c chan int, x int) {
	log.Println(x)
	c <- x
	close(c)
	log.Println(x)
}

func main() {
	c := make(chan int)
	x := 10
	go writeToChannel(c, x)
	time.Sleep(time.Second * 1)
}
