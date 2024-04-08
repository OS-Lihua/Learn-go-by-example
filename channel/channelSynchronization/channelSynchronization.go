package main

import (
	"fmt"
	"time"
)

func work(done chan bool) {
	fmt.Println("hello")
	time.Sleep(3000)
	fmt.Println("world")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go work(done)
	<-done
}
