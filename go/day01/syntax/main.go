package main

import "fmt"

//根据值自动推导变量类型，必须放到函数内部(在包的顶部，这个是最常用的)
// 在程序运行时，是在编译时完成类型推导，运行时分配内存

var (
	socre = 99.5
)

// const 定义变量
// 通常常量大写
// 什么时候需要变量，当一个值不会被修改的时候，使用常量
// 变量时，做完程序数据存在的
// 变量，值在编译时就已经确定下来了，被编译成为程序的一部分

const (
	PI           = 3.14
	SKILL_GOLANG = "Golang"
)

func main() {
	//有一个插槽来存储用户的数据，或者程序的输入，这种插槽叫变量
	//变量的声明：var 变量名 变量类型
	//变量的赋值：变量名 = 值

	//名称变量(最麻烦的写法)
	var name string
	name = "bob"

	//变量的声明和赋值可以合并成一句话
	var age int = 18

	//自动推导变量类型，只能放在函数内部，(在函数内部，这个最常用)
	// := 左侧是变量名，右侧是值
	country := "china" //var country string = "china"
	_ = country        // 避免未使用报错 我存在可以不用

	fmt.Println("hello," + name + "!")
	fmt.Println("age =", age)
	fmt.Println("socre =", socre)
	fmt.Println("country =", country)

	// hello, bob!
	// age = 18
	// score = 99.5
	// country = china

}
