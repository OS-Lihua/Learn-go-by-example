// 通道同步
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
	// 阻塞 直到 “work” 协程 中
	// done <-
	<-done
}
