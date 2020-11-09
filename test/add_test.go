package main

import (
	"testing" // 单元测试的包
)

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	if res != 1 {
		t.Fatalf("期望%v 实际%v\n", 1, res)
	}
	t.Log("ok")
}
