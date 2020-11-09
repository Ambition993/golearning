package main

import (
	"encoding/json"

	"fmt"
)

func main() {
	testStruct()
	testMap()
	testSlice()
}

type Monster struct {
	Name     string
	Age      int
	Brithday string
	Sal      float64
	Skill    string
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
func testMap() {
	m := make(map[string]int)
	m["hello"] = 1
	m["nihao"] = 2
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(string(data))
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
