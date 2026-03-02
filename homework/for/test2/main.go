package main

import "fmt"

func main() {

	height := 5

	for i := 1; i <= height; i++ {

		// ① 打印前面的空格 第一行有四个空格 5-1
		for s := 0; s < height-i; s++ {
			fmt.Print(" ")
		}

		// ② 打印递增数字 1 ~ i 四个空格让然后第一行的数字1 第二行的数字2
		for n := 1; n <= i; n++ {
			fmt.Print(n)
		}

		// ③ 打印递减数字 i-1 ~ 1
		for n := i - 1; n >= 1; n-- {
			fmt.Print(n)
		}

		// 换行
		fmt.Println()
	}
}
