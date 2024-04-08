// 通道选择器
package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	go func() {
		//time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		//time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// ||  ||			阻塞
		//  \  /
		//   \/
		//   ||
		select {
		// 阻塞
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		// 阻塞
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
