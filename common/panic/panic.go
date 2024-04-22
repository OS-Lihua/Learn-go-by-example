package main

import "os"

func main() {
	panic("a problem")
	// 抛出错误,程序退出,后续不可执行
	_, err := os.Create("./tmp/file")
	if nil != err {
		panic(err)
	}
}
