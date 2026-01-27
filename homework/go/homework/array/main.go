package main

import "fmt"

func main() {
	//整数数组
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	//字符串数组
	arr1 := [5]string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(i, arr1[i])
	}

	//创建一个数组 拷贝到另一个数组，修改原数组的第一个元素，观察拷贝后的数组是否改变
	addr := [...]int{10, 20, 40, 50, 30}
	addr1 := addr
	fmt.Println("addr1:", addr1)
	addr1[0] = 100
	fmt.Println("addr1:", addr1)

	//数组作为函数参数

	nums := [5]int{3, 5, 1, 4, 2}
	nums1 := SumNums(nums)
	fmt.Println("sum:", nums1)

}
func SumNums(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
