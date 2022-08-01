package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println(123)
	// 序列化
	v1 := []interface{}{
		123,
		true,
		3, 2,
		map[string]interface{}{
			"name": "sdf",
		},
		Person{"jojo", 12},
	}
	res, _ := json.Marshal(v1)
	data := string(res)
	println(data)
	// 反序列化
	var value []interface{}
	content := `[123,true,3,2,{"name":"sdf"},{"name":"jojo","age":12}]`
	json.Unmarshal([]byte(content), &value)
	fmt.Println(value)
}
