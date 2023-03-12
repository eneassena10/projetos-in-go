package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "one = 1"
	c <- "two = 2"
	close(c)
	for el := range c {
		fmt.Println(el)
	}
}
