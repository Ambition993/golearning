package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := ""
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var count charCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("err")
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.CharCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.Others++
			}

		}
	}
	fmt.Println(count.CharCount, count.NumCount, count.SpaceCount, count.Others)
}

type charCount struct {
	CharCount  int
	NumCount   int
	SpaceCount int
	Others     int
}
