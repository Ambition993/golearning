package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 1
	reflectTest(&num)
	//这里传入的是一个指针的类型所以
	fmt.Println(num)
}

func reflectTest(i interface{}) {
	rval := reflect.ValueOf(i)
	// 这里直接设置一个值会出错的
	//所以先找到这个指针指向的值然后再设置值
	//rval.SetInt(33)
	rval.Elem().SetInt(111)
}
