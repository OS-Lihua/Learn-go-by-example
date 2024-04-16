package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
	const numChannals = 5
	var wg sync.WaitGroup

	for i := 1; i <= numChannals; i++ {
		wg.Add(1)

		i := i
		go func() {
			// 压栈
			// Done -> Add(-1)
			defer wg.Done()
			worker(i)
		}()
	}
	// 等待计数器为 0
	wg.Wait()
}
