package main

import "fmt"

func main() {
	go test()
	go hello()
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() occurrs an err")
		}
	}()
	var myMap map[string]int
	myMap["hello"] = 1 //there is mistake will panic
}
func hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
}
