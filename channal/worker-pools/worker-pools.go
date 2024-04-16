// 协程池
package main

import (
	"fmt"
	"time"
)

func work(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println("Work Id ", id, "start job", i)
		time.Sleep(2 * time.Second)
		fmt.Println("Work Id ", id, "finish job", i)
		// ##
		// results <- i * 2
	}
	results <- id
}

func main() {
	const numJobs = 10
	const numChannals = 3
	jobs := make(chan int, 1)
	results := make(chan int, 1)

	// 开了 numChannals 个 工作协程
	for i := 1; i <= numChannals; i++ {
		go work(i, jobs, results)
	}

	// 发送完 job
	for i := 1; i <= numJobs; i++ {
		// 阻塞,等待接收端就绪
		// 因为chan 长度只有 1
		jobs <- i
		// fmt.Println(1, time.Now())
	}

	close(jobs)

	// 如果 ## 使用这种方法,阻塞协程
	// 当两个channel 缓冲区都用完了，非常容易造成 死锁
	// ##
	// for i := 1; i <= numJobs; i++ {
	//	 <-results
	// }

	for i := 1; i <= numChannals; i++ {
		<-results
	}
}
