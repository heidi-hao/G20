
## 切片 slice
Go中的slice依赖于数组，他的底层就是数组，所以数组有的优点，slice都有
且slice可以通过append向slice中追加元素，长度不够的时候可以动态扩容，通过再次slice切片，可以得到更小的slice结构，可以迭代，遍历等

基于数组的复合数据结构
type slice struce {
    array unsafe.Pointer //数组指针
    cap int // 容量
    len int // 长度
}

每个slice结构都由三部分组成
1.容量(capacity):即底层数组的长度，表示这个slice目前最多能扩展到这么长
2.长度(length):表示slice当前的长度，即当前容纳的元素个数
3.数组指针(array):指向底层数组的指针

比如创建一个长度为3 容量为5 int类型的切片
s := make([]int,3,4)
fmt.Println(a, len(s), cap(s)) // [0 0 0] 3 5

切片和数组 为了灵活(不需要制定长度，自动扩容)
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

	//1.声明并赋值
    
	slice12 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice12:", slice12)
}

创建和初始化
原理：切片是动态数组，底层基于固定大小的数组，可以根据动态扩容。创建时可以指定长度和容量。


