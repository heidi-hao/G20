package main

import (
	"fmt"
)

func main() {
	// 1.声明切片 切片不是一个值，是一个boxed的结构体，array unsafe Pointer //数组指针
	//底层数组的长度：容量10 当前有几个元素3
	slice1 := make([]int, 3, 5) //
	fmt.Println(slice1, len(slice1), cap(slice1))
	//底层数组的长度：容量10，当前有几个元素
	slice1 = append(slice1, 4, 5)
	fmt.Println(slice1, len(slice1), cap(slice1))
	//扩容
	slice1 = append(slice1, 6)
	fmt.Println(slice1, len(slice1), cap(slice1))

	//2.声明并赋值(不常用)
	var slice2 []int
	slice2 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice2, len(slice2), cap(slice2))

	//3.声明赋值放在一起(常用)
	var slice3 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice3, len(slice3), cap(slice3))

	//4.快捷声明赋值（常用）
	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice4, len(slice4), cap(slice4))

	/*
	 切片的访问
	 1.通过下标(元素的索引)访问切片元素
	 2.元素值修改，直接通过下标赋值即可
	*/
	fmt.Println("slice4[0]=", slice4[0])
	slice4[0] = 100
	fmt.Println("slice4[0]=", slice4[0])

	/*
		     切片中添加元素
			 append(切片变量,元素1,元素2,.....)
			 注意append后时参数一个新的切片变量 所以要接收返回值
			 工作: 申请了新的底层数组，老数据copy过去，然后添加新的元素
			 append的元素如果超过容器，会出发地址扩容，新的数组
	*/

	slice4_v2 := append(slice4, 6, 7, 8) //超过了 需要扩容
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
	fmt.Println("slice4_v2:", slice4_v2, len(slice4_v2), cap(slice4_v2))

	/*
		1.nil 切片: 没有分配内存(array指针是nil),长度和容量都是0
		2.空切片
		切片时一个应用类型(盒子)
	*/
	var slice5 []int //nil切片。nil 本身是“无类型的空” 所以不可以使用推断式
	fmt.Println(slice5, len(slice5), cap(slice5))
	slice5 = append(slice5, 10, 11, 12)
	fmt.Println(slice5, len(slice5), cap(slice5))

	slice6 := []int{}
	fmt.Println(slice6, len(slice6), cap(slice6))
	slice6 = append(slice6, 10)
	fmt.Println(slice6, len(slice6), cap(slice6))

	/*
		     切片的遍历
			 1.for 循环遍历 实际遍历的也是底层的数组
			 2.for range遍历
	*/
	fmt.Println("切片的遍历: for ==")
	for i := 0; i < len(slice4_v2); i++ {
		fmt.Println("slice4_v2[%d]=%d\n", i, slice4_v2[i])
	}

	fmt.Println("切片的遍历: for range == ")
	for index, value := range slice4_v2 {
		fmt.Println("slice4_v2[%d]=%d\n", index, value)
	}

	/*
		    切片是一种应用类型
			注意: 数组时引用类型吗？不是
			1.如果是数组:是2块隔离的内存地址空间(线性分配) 数据改变不影响原来的
			2.如果是切片:底层数组是同一块内存地址空间 数据改变都会改变
	*/
	fmt.Println("==切片是一种引用类型==")
	slice7 := []int{1, 2, 3} //切片类型
	slice8 := slice7
	fmt.Println("slice7:", slice7)
	fmt.Println("slice8:", slice8)
	slice8[0] = 1000
	fmt.Println("slice7:", slice7)
	fmt.Println("slice8:", slice8)

	/*
		切片的拷贝
		如果希望把底层的arry里的数据copy到另一个切片中
		使用内置函数：copy(目标切片，源切片)
		拷贝: 底层数据已copy一份，申请一片新的内存来存储这个值，只copy一层，浅拷贝
		深copy是递归copy，需要自己实现，依赖三方库，非常耗性能，一般不推荐使用，最常用的方式是json序列化和反序列化来实现深copy
	*/
	fmt.Println("==切片的拷贝==")
	//创建新切片加新底层数组
	slice9 := make([]int, len(slice7)) //slice9 := make([]int, len(slice7)) 创建了一个“长度 = slice7 长度”的新切片，并且分配了一个全新的底层数组
	copy(slice9, slice7)               //拷贝数量 = min(len(dst), len(src)
	fmt.Println("slice9:", slice9)
	slice9[0] = 2000
	fmt.Println("slice7:", slice7)
	fmt.Println("slice9:", slice9)

	fmt.Println("==切片 切割(共享底层数组)==")
	slice10 := []int{1, 2, 3}
	fmt.Println("slice10:", slice10) // [1,2,3]
	slice11 := slice10[1:3]          //[2,3] 切片切割不会创建新数组
	fmt.Println("slice11:", slice11)
	slice11[0] = 200                 //[200,3]
	fmt.Println("slice10:", slice10) //[1,200,3]
	fmt.Println("slice11:", slice11)

	/*
		切片的删除
		切片的删除是通过切片操作实现的，即使用切片语法来一处指定范围的元素
	*/
	fmt.Println("==切片的删除==")
	slice12 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice12:", slice12)

	//删除索引是2的元素 删除 i：拼接 i 左边 和 i 右边 slice = append(slice[:i], slice[i+1:]...)
	slice12 = append(slice12[:2], slice12[3:]...)
	fmt.Println("slice12:", slice12)

	/*
		函数作为参数
	*/
	fmt.Println("==函数作为参数==")
	slice13 := []int{5, 3, 4, 2, 1}
	fmt.Println("排序前:", slice13)

	//值传递
	slice13_sorted := MySort(slice13)
	fmt.Println("排序后(值传递):", slice13_sorted)
	fmt.Println("原切片slice13:", slice13)

	//引用传递
	//MySortV2(&slice13)
	//fmt.Println("排序后的(引用传递):",slice13)
}

// 冒泡排序:值传递
// 返回一个新的切片
func MySort(arr []int) []int {
	//冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
