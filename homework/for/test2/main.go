package main

import "fmt"

func main() {

	height := 5

	for i := 1; i <= height; i++ {

		// ① 打印前面的空格
		for s := 0; s < height-i; s++ {
			fmt.Print(" ")
		}

		// ② 打印递增数字 1 ~ i
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
