package main

import (
	"fmt"
)

func main() {
	/*
	   数组拷贝
	*/
	//ipv4 4字节表示一个ipv4地址，简短声明127.0.0.1
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println(ip) //[127 0 0 1]

	//如何copy这个数组,
	ip2 := ip
	fmt.Println("ip2:", ip2) //ip2: [127 0 0 1]

	//
	ip2[0] = 10
	fmt.Println(ip2) // [10 0 0 1]

	/*
	   数组遍历
	*/
	// 对象的编号，通常叫索引，index(0,....)
	//1.for：定义变量(开始元素的索引) 条件(不满足这个条件，就结束循环) 变化(索引加1，找一下元素)

	//for i := 0; i < len(ip); i++ {
	//	fmt.Println(ip[i])

	//}

	//for u := range ip {
	//fmt.Println(ip[u])

	//}

	//专门针对数组，切片，字符串的遍历
	// range array
	//(index,value)
	for i, v := range ip {
		fmt.Println(i, v)
	}

	//排序
	arr1 := [5]int{3, 5, 1, 4, 2}
	sorted_arr1 := MySort(arr1)
	fmt.Println("source arr1:", arr1)
	fmt.Println("sorted_arr1", sorted_arr1)

	//数组指针
	MySortV2(&arr1)
	fmt.Println("source arr1:", arr1)

}

// Go的函数参数必须有名字
func MySort(arr [5]int) [5]int { // 参数：名字叫 arr，类型是 [5]int  返回值类型[5]int
	//冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				//值交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 这是 通过指针直接修改原数组 的排序函数
func MySortV2(arr *[5]int) { //*指针  [5]int 长度为5的数组 总的来说就是arr 是一个 指向 [5]int 数组的指针
	//冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
