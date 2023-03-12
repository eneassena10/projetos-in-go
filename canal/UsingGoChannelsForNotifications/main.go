package main

import (
	"log"
	"time"
)

type T = struct{}

func main() {
	completed := make(chan T)
	go func() {
		log.Println(("ping"))
		time.Sleep(time.Second * 5)
		<-completed
	}()
	completed <- struct{}{}
	log.Println("pong")
}
