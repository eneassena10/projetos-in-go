package main

import (
	"fmt"
	"math/rand"
	"time"
)

func run(ch chan int32) <-chan int32 {
	defer close(ch)
	time.Sleep(time.Second * 5)
	ch <- rand.Int31n(300)
	return ch
}

func longTimedOperation() <-chan int32 {
	ch := make(chan int32)
	go run(ch)
	return ch
}

func main() {
	ch := longTimedOperation()
	fmt.Println(ch)
}
