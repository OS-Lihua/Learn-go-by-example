// 通道缓冲区
package main

import "fmt"

func main() {
	// 初始化通道缓冲区 为 2
	// 默认为 0
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
