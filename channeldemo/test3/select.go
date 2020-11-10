package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	strChan := make(chan string, 5)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	for i := 0; i < 5; i++ {
		strChan <- "hello"
	}
	for {
		select {
		case v := <-intChan:
			fmt.Println(v)
		case v := <-strChan:
			fmt.Println(v)
		default:
			fmt.Println("finish")
			break
		}

	}
}
