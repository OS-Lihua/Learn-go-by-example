package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// 每 200 ms 向 通道发送一个 msg,打点器
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//request 1 2024-04-11 11:58:31.1208194 +0800 CST m=+0.203278401
	//request 2 2024-04-11 11:58:31.3258516 +0800 CST m=+0.408310601
	//request 3 2024-04-11 11:58:31.5312663 +0800 CST m=+0.613725301
	//request 4 2024-04-11 11:58:31.7325585 +0800 CST m=+0.815017501
	//request 5 2024-04-11 11:58:31.9339185 +0800 CST m=+1.016377501	每个都相差 200 ms 左右

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		// 先发送 3 个,连续触发
		burstyLimiter <- time.Now()
	}

	go func() {
		/// 每 200 ms 触发一次
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
	//request 1 2024-04-11 11:58:31.9339185 +0800 CST m=+1.016377501
	//request 2 2024-04-11 11:58:31.9339185 +0800 CST m=+1.016377501
	//request 3 2024-04-11 11:58:31.9344445 +0800 CST m=+1.016903501		// 连续触发
	//request 4 2024-04-11 11:58:32.1367082 +0800 CST m=+1.219167001
	//request 5 2024-04-11 11:58:32.3361064 +0800 CST m=+1.418565401		// 和上一个差 200 ms 左右
}
