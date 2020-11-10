package main

import "fmt"

func main() {
	intChan := make(chan int, 500)
	resChan := make(chan int64, 2000)
	exitChan := make(chan bool, 4)
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go getRes(intChan, resChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(resChan)
	}()

	for {
		val, ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("res is %v", val)
	}
}

func putNum(intChan chan int) {
	for i := 1; i < 15; i++ {
		intChan <- i
	}
	close(intChan)
}

func getRes(intChan chan int, resChan chan int64, exitChan chan bool) {

	for {
		var res int64 = 1
		val, ok := <-intChan
		if !ok {
			break
		}
		for i := 1; i <= val; i++ {
			res = res * int64(i)
		}
		resChan <- res
	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//这里我们还不能关闭 primeChan
	//向 exitChan 写入true
	exitChan <- true
}
