package main

import "fmt"

func main() {
	arr01 := []int{1, 2, 3, 4, 5}
	// 切片的for
	for index, value := range arr01 {
		fmt.Println(index, arr01[index])
		fmt.Println(value)
	}
	// map for

	//m*n. 外层：控制行。内层：控制列
	for m := 1; m <= 9; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d*%d=%d\t", n, m, n*m)
		}
		fmt.Println()
	}

}
