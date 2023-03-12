package main

import (
	"log"
	"time"
)

// Os canais também podem agendar notificações:

func scheduleNotification(t time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(t)
		ch <- struct{}{}
	}()
	return ch
}

func main() {
	log.Println("send first")
	<-scheduleNotification(time.Second)
	log.Println("send second")
	<-scheduleNotification(time.Second)
	log.Println("lastly send")
}
