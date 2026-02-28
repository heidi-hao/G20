package main

import (
	"fmt"
	"reflect"
)

func main() {
	//bike对象
	bike := TranficTool{Name: "bicycle", Speed: 20, Color: "black"}
	JsonEncode(&bike)
	fmt.Println(bike)
}

// obj -> []byte 类型是any
func JsonEncode(obj any) []byte {
	//反射获取类型信息
	ot := reflect.TypeOf(obj)

	if ot.Kind() != reflect.Pointer && ot.Kind() != reflect.Struct {
		panic("必须是结构体指针类型")
	}

	ot = ot.Elem() //获取指针指向的类型

	fmt.Println("类型名称:", ot.Name())
	fmt.Println("类型种类:", ot.Kind())
	//类型名称： TranficTool
	//类型种类： struct

	//反射获取值信息
	ov := reflect.ValueOf(obj).Elem()
	fmt.Println("值:", ov)
	//反射获取结构体字段信息(TranficTool)
	for i := 0; i < ot.NumField(); i++ {
		field := ot.Field(i)
		fieldValue := ov.Field(i)
		if field.Name == "Name" {
			fmt.Println()
		}
	}

}

type TranficTool struct {
	//`` 表示字符串 都是Struct Tag 结构体标签 lib 反射作用
	Name  string `json:"name"`
	Speed int    `json:"speed"`
	Color string `json:"color"`
}
