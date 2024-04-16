// 协程
// go func()
package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, s)
	}
}

func main() {
	go f("here")
	go f("yaco")
	f("main")
	time.Sleep(3)
	fmt.Println("ending ...")
}
