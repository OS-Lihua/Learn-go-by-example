package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("test.txt")
	defer closeFile(f)
	writeFile(f, "hello,world!")
}

func createFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, str string) {
	fmt.Println("writing:")
	fmt.Fprintf(f, str)
}

func closeFile(f *os.File) {
	err := f.Close()
	if nil != err {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
