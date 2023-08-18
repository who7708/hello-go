package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	sonic "github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
)

// 测试字节的golang json库使用
func main() {

	type Data struct {
		Name string `json:"name"`
		Age  int8   `json:"age"`
	}

	data := []Data{{
		Name: "zhangsan",
		Age:  11,
	}, {
		Name: "lisi",
		Age:  33,
	}}

	// 使用标准 json 库
	// r1, _ := json.Marshal(&data)
	r1, _ := json.MarshalIndent(&data, "", "  ")
	fmt.Printf("reflect.TypeOf(r1): %v\n", reflect.TypeOf(r1))
	fmt.Printf("r1: %v\n", r1)
	fmt.Printf("string(r1): %v\n", string(r1))

	// json-iterator/go 库
	JSON := jsoniter.ConfigCompatibleWithStandardLibrary
	// r2, _ := JSON.Marshal(&data)
	// r2, _ := JSON.MarshalToString(&data)
	r2, _ := JSON.MarshalIndent(&data, "", "  ")
	fmt.Printf("reflect.TypeOf(r2): %v\n", reflect.TypeOf(r2))
	fmt.Printf("r2: %v\n", r2)
	fmt.Printf("string(r2): %v\n", string(r2))

	// sonic 字节的json库
	// r3, _ := sonic.Marshal(&data)
	// r3, _ := sonic.MarshalString(&data)
	r3, _ := sonic.ConfigDefault.MarshalIndent(&data, "", "  ")
	fmt.Printf("reflect.TypeOf(r3): %v\n", reflect.TypeOf(r3))
	fmt.Printf("r3: %v\n", r3)
	fmt.Printf("string(r3): %v\n", string(r3))

}
