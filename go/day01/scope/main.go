package main

import "fmt"

// 全局变量
var (
	//main 包里里面的变量
	age = 18
)

func main() {

	var a = "parent"
	fmt.Println("outer:", a) //parent {}创建新的代码块作用域 声明一个全新的局部变量
	{
		a := "child"             //子作用域声明了同名新变量 遮蔽外层的a
		fmt.Println("inner:", a) //child
	}
	//回到外层
	fmt.Println("outer again:", a) //parent

	//局部变量
	var b = "parent"
	fmt.Println("outer:", b) // parent
	{
		// var b = "child"
		// 新的局部变量b
		b := "child"             //子作用域声明了同名新变量 遮蔽外层a
		fmt.Println("inner:", b) //child
	}
	fmt.Println("outer again:", b) //parent

	fmt.Println(age) // 18 全局变量是不会被遮蔽的

}
