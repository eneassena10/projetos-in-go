package main

import (
	"log"
	"time"
)

func writeToChannel(c chan int, x int) {
	log.Println("valor de x 1: ", x)
	c <- x
	close(c)
	log.Println("valor de x 2:", x)
}

func main() {
	c := make(chan int)
	x := 10
	go writeToChannel(c, x)
	time.Sleep(time.Second * 1)
	log.Println("Read:", <-c)
	time.Sleep(1 * time.Second)
	_, ok := <-c
	if ok {
		log.Println("Channel is open")
	} else {
		log.Println("Channel is closed")
	}
}
