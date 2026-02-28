package main

import "fmt"

func main() {
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println(x, y)

	arr := [3]int{1, 2, 3}
	doubleArray(&arr)
	fmt.Println(arr)

	cfg := &Config{}
	initConfig(cfg, "localhost", 8080)
	fmt.Println(cfg)

}

// 交换两个整数的值(通过指针)
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp

}

// 将所有的元素加倍
func doubleArray(arr *[3]int) {
	//循环遍历每一个数组的数字
	for i := 0; i < len(arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}

}

// 初始化结构体指针
type Config struct {
	Host string
	Port int
}

func initConfig(c *Config, host string, port int) {
	c.Host = host
	c.Port = port

}
