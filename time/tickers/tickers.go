// 打点器
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义打点器
	// 	<-	|————|  <-
	// 	<-	|————|  <-
	// 每 500 millisecond 向发送端发送一次信息(触发一次)
	tikers := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-tikers.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	tikers.Stop()
	done <- true
	fmt.Println("Ticker Stopped")
}
