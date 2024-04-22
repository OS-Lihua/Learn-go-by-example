// recover
// 1. 可以从 panic 中 恢复recover
// 2. recover 可以阻止 panic 中止程序,并让它继续执行
package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	/// recover 用于捕获 panic
	defer func() {
		// 必须在 defer 函数中调用 recover
		if r := recover(); r != nil {
			fmt.Println("Recovered, Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("main ending")
}
