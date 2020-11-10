package main

import (
	"fmt"
	"time"

	"sync"
)

func main() {
	//计算1-200的各个数字的阶乘 并且放在map里面
	//最后显示出来 使用goroutine

	//思路 ： 编写一个函数 来计算
	// 启动协程来统计结果
	//map 是全局的
	for i := 1; i < 30; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	lock.Lock()
	for k, v := range MyMap {
		fmt.Println(k, v)
	}
	lock.Unlock()
}

var (
	MyMap = make(map[int]int)
	lock  sync.Mutex
)

func test(num int) {
	res := 1
	for i := 1; i < num; i++ {
		res *= i
		lock.Lock()
		MyMap[i] = res
		lock.Unlock()
	}
}
