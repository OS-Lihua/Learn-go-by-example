// 遍历带缓冲区的通道
package main

import "fmt"

func main() {
	pings := make(chan string, 3)
	pings <- "hello"
	pings <- "world"
	pings <- "!"
	close(pings)
	// 关闭缓冲区后，不能发送
	// pings <- "good"	//	ERROR

	// 关闭缓冲区后，依然可以接收
	fmt.Println(<-pings)

	//  buffer size : 3
	// 	<-	|————-————-————|	<-
	// 	<-	|————-————-————|	<-
	// 1. 关闭通道是指 关闭通道的发送端
	// 2. 但是通道的接收端 依然保持开启

	for elem := range pings {
		fmt.Println(elem)
	}
	// range 通道 这个过程
	// 就是 逐个 接收 (<- chan) 的 过程

	// 这里不再阻塞
	<-pings
	// 3. 当 通道缓冲区 的 数据 全部被接收
	// 4. 通道的接收端 也被关闭
}

// 这里加深了对 close(channel) 的理解
// 比较好的比喻就是:
// 先关水龙头,把水放完,再关出水口
