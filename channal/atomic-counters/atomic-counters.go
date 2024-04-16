package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				// 原子操作
				// 准确输出 50000
				atomic.AddUint64(&ops, 1)
				// 反例:
				// ops++
				// 输出一个未知数
			}
			wg.Done() // wg.Add(-1)
		}()
	}

	wg.Wait()
	fmt.Println("ops :", ops)
}
