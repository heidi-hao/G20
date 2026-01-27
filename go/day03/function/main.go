package main

import "fmt"

func main() {
	//1,2参数:实际传入的参数
	// x = 2 y = 3
	fmt.Println(sum(1, 2))
	//切片作为参数传递(通过切片来包装)
	fmt.Println(sum2([]int{2, 3, 4}))
	fmt.Println(sum3(2, 3, 4))

	//引用类型示例 底层扩容的影响
	//->底层数组A(未扩容)
	arg1 := []int{2, 3, 4}
	result1 := appendAndChange(arg1)
	fmt.Println("result1:", result1)

	arg2 := []int{2, 3, 4}
	appendAndChangeV2(&arg2)
	fmt.Println(arg2)

	counterFn := counter()
	fmt.Println(counterFn(), counterFn(), counterFn())

	// fib
	fmt.Println(fib(5))

}

// (x int,y int)
// var x int
// var y int
// x,y:形参 形式参数(函数定义时的参数，声明的参数)

func sum(x int, y int) int {
	return x + y
}

func sum2(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

// 能不让函数直接接收多个不确定参数
// ...int：不定长参数---->[]int
func sum3(args ...int) int {
	var total int
	for _, arg := range args {
		total += arg
	}
	return total
}

// [2,3,4]
func appendAndChange(s []int) []int {
	fmt.Println(s, len(s), cap(s))
	s[0] = 9
	s = append(s, 10, 11)
	fmt.Println(s, len(s), cap(s))
	s[0] = 8
	return s
}

// []int -> x
// *[]int -> box(x)
func appendAndChangeV2(s *[]int) {
	// 指针的值copy到了v
	//[]int
	// v := *s
	// 产生一个新的值复杂对象(赋值的是元数据)
	//box(x) -> x

	//这里的v就不再是s(v值)， v 是 s值的一个快照
	//不是操作这个地址本身 操作的是值
	fmt.Println(s, len(*s), cap(*s))
	(*s)[0] = 10 //影响外部底层数组(底层数组未扩容时(A))
	fmt.Println(s, len(*s), cap(*s))

	//解开box 把10，11放进去
	// s的扩容就是box的扩容了 对我们外部可见
	*s = append(*s, 10, 11) //可能扩容 之后的修改不在影响外部 cap从3->6
	(*s)[0] = 20
}

// 闭包
func counter() func() int {
	//依赖的变量
	i := 0

	//m 会被GC回收掉，外部没有引用记录，意思就是被用完了
	m := 10
	fmt.Println(i, m)
	//i还被返回的依赖着 除非这个函数不用了
	return func() int {
		i++
		return i
	}
}

// 斐波那契数列 fn(n) = fn(n-1) + fn(n-2)
// 0 1 1 2 3 5 8 13 21 34
// fib(2) = 1 : n = 2 fib(1) + fib(0)
// fib(0) = 0
// fib(1) = 1
// 函数作为参数传递 递归调用 直到函数退出
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 遍历目录里面的文件
