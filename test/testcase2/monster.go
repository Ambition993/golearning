package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	//先序列化 再保存
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("error")
		return false
	}
	filePath := "D:/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("save err")
		return false
	}
	return true
}
func (this *Monster) Restore() bool {
	filepath := "D:/monster.ser"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("readfile err", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err")
		return false
	}
	return true
}
