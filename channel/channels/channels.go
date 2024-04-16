// 通道
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	go func() {
		// 由于 此通道是没有缓冲区的
		// 如果 通道没有接收方，则发送不了
		messages <- "hello,world"
	}()
	// 阻塞, 等 通道send 信息
	msg := <-messages
	fmt.Println(msg)
}
