package main

import (
	"fmt"
)

func main() {

	//创建一个字符串切片，包含"Go", "Python", "Java"，并打印长度和容量
	slice := []string{"Go", "Python", "Java"}
	fmt.Println(slice, len(slice), cap(slice))

	//给定切片[]int{1,2,3,4,5}，创建子切片包含第2到第4个元素（不包括第4个），并打印结果
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3]
	fmt.Println(slice2)

	//从空切片开始，使用append添加5个整数，观察长度和容量的变化。
	slice3 := []int{}
	fmt.Println(slice3)
	slice4 := append(slice3, 1, 2, 3, 4, 5)
	fmt.Println(slice4, len(slice4), cap(slice4))

	//遍历一个整数切片，计算所有元素的和。
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice5)
	sum := 0
	for i := 0; i < len(slice5)-1; i++ {
		sum += slice5[i]
	}
	fmt.Println("切片:", slice5)
	fmt.Println("sum:", sum)

	//创建一个切片，拷贝到另一个切片，修改原切片，验证拷贝是否独立。//创建新切片加新底层数组
	slice6 := []int{1, 2, 3, 4, 5}
	slice7 := make([]int, len(slice6)) //make给新切片分配一个全新的数组
	copy(slice7, slice6)
	fmt.Println(slice6, len(slice6), cap(slice6))
	slice7[0] = 1000
	fmt.Println(slice7)

	//值传递
	slice8 := []int{5, 3, 6, 1, 2, 4}
	slice9 := MySort(slice8)
	fmt.Println(slice9)

	//引用传递
	MySortV2(&slice8)
	fmt.Println(slice8)

}

// 返回一个新的切片(排序结果)
func MySort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 冒泡排序：引用传递
// 直接修改传入的切片(在原切片上进行排序)
func MySortV2(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
