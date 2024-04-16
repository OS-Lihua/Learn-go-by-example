// 通道缓冲区
package main

import "fmt"

func main() {
	// 初始化通道缓冲区 为 2
	// 默认为 0
	// 注意:
	// 1. 实际可以缓冲的长度 = 发送端可以发送的长度 = 缓冲区长度 + 1
	// 2. 缓冲区长度用完了,仍然是可以发送信息的(zui相当于无缓冲区的channel)
	// 3. 当channel 实际缓冲区的长度用完后, 发送端开始阻塞
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
