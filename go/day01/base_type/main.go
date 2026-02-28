package main

import "fmt"

// 自定义类型
type Status string

const (
	// 自定义类型的值，常量不允许修改, 充当枚举(可选项)
	STATUS_ONLINE  Status = "online"
	STATUS_OFFLINE Status = "offline"
	STATUS_UNKNOW  Status = "unknow"
)

func main() {

	u := User{Name: "heidi", age: 18}
	fmt.Println("%+v \n", u)

	a := []byte{} //长度为 0 的切片
	// %+v 打印结构体字段名
	// []
	fmt.Printf("%+v \n", a)

	// 字符采用unicode编码, 这里的 值得这个字符串编码是 97
	b := 'a'
	fmt.Printf("%+v\n\\n", b)

	name := "bob \"引号\""
	// 用于格式化 string类型的值
	fmt.Printf("hello, %+v\n", name)

	multiLine := `这是一个多行字符串
可以包含换行符号
	也可以包含 "双引号"
	还可以包含 '单引号'`
	fmt.Printf("\n%+v", multiLine)

	// 判断变量是否online
	status := STATUS_UNKNOW
	if status == STATUS_ONLINE {

	}
}

type User struct {
	Name string
	age  int
}
