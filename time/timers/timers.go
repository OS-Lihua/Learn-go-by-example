// 定时器
package main

import (
	"fmt"
	"time"
)

func main() {

	// 声明定时器, 时间通道
	//   	|————|
	//   	|————|
	// 2 seconds 后 会在发送端发送信息
	timer1 := time.NewTimer(2 * time.Second)
	// 接收端 阻塞
	// 等待 发送端发送信息
	// 发送端发送信息,触发定时器
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 声明 1 second 的 定时器
	timer2 := time.NewTimer(time.Second)
	go func() {
		// 阻塞
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// 关闭 定时器
	// 关闭后, <-timer2.C 将一直阻塞
	//  timer2 将无法触发
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 结束
	time.Sleep(3 * time.Second)
}
