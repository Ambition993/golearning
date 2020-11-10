package main

import (
	"fmt"
	"time"
)

func main() {
	go test() // 开启了一个协程

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

}
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println("golang")
		time.Sleep(time.Second)
	}
}

//主线程和协程一起执行
