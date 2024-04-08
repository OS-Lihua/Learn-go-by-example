// 通道方向
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	// 定义带缓冲区的通道
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "hello,world")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
