package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	c := sum(a, b)
	fmt.Println("a+b=", c)

	//使用指针作为函数参数
	var result int

	//&result （类型安全的指针）东西叫指针 有类型约束
	fmt.Println("result 地址=", &result)
	sumWithPointer(a, b, &result)
	fmt.Println("a+b=", result)

	//外部改变内部的值
	x := 1
	change(&x)
	fmt.Println(x) //100

}

func change(p *int) { //p是一个指向int的指针
	*p = 100
}

func sum(a int, b int) int {
	return a + b
}

// 先写一个函数 然后在去写后续的逻辑
// 这个函数没有return
// *int 就是指针类型 传递是 int这种类型的变量的地址
func sumWithPointer(a int, b int, c *int) {
	*c = a + b
}
