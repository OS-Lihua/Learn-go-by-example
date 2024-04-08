package main

import "fmt"

func main() {
	pings := make(chan string, 2)
	pings <- "hello"
	pings <- "world"

	close(pings)

	for elem := range pings {
		fmt.Println(elem)
	}
}
