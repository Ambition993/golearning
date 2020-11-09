package main

import (
	"encoding/json"

	"fmt"
)

func main() {
	testStruct()
	testMap()
	testSlice()
	unMarshal()
}

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Brithday string  `json:"brithday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func unMarshal() {
	str := testMap()
	var m map[string]interface{}
	//反序列化的时候不需要make操作
	//已经封装了这个操作了

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(m)
}
func testStruct() {
	monster := Monster{
		"haha",
		599,
		"1999-11-11",
		11.23,
		"hah",
	}
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("wrong")
	}
	fmt.Println(string(data))
}
func testMap() string {
	m := make(map[string]int)
	m["hello"] = 1
	m["nihao"] = 2
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(string(data))
	return string(data)
}
func testSlice() {
	slice := make([]int, 3)
	slice = append(slice, 1, 2, 34)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(string(data))
}
