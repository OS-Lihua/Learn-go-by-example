// 非阻塞 通道操作
//
//	select{
//		case :
//		default :
//	}
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	// case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	// message 是 缓冲区为 0 的 通道
	// 没有接收方就绪(阻塞)
	// 发送失败
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
