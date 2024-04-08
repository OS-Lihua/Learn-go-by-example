// 关闭通道
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			// jobs int
			// more bool
			job, more := <-jobs
			if more {
				fmt.Println("receive: ", job)
			} else {
				fmt.Println("receive all msg")
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("send: ", i)
	}
	close(jobs)
	fmt.Println("send all msg")
	<-done
}

// 	结果 与(https://gobyexample-cn.github.io/closing-channels)网站不同
//	send:  0
//	send:  1
//	send:  2
//	send:  3
//	send all msg
//	receive:  0
//	receive:  1
//	receive:  2
//	receive:  3
//	receive all msg
