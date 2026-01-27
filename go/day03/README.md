
## 函数和结构体

基础数据类型，复合数据类型(数组 切片 Map) 前面章节已经开始使用函数了 函数的具体使用方法 

我们前面介绍的数组 只能保存同一种数据的类型 当我们需要记录多种不同类型的数据的时候 我们需要怎么办呢
结构体就是用于解决这个问题的 结构体是由一系列有相同类型或者不同类型的数据构成的数据集合 方便容量我们任何类型的数据


数据结构

1.基础数据结构： int float bool string byte(int8) rune
2.复合数据结构 array slice map

单独变量的方式：
name(string) 
age(int)
gender(string)
weight(uint)
favoriteColor([]string)

根据业务需要自己进行数据定义 这就是结构体
type Person struch {
    Name string
    Age  int
    Gender string
    Weight uint
    FavoriteColor []string
}

func handlePerson(*Person) error

//person 结构体类型的初始化
//把struct上这些变量 赋值的过程 就叫实例化
p1 := Person{
    Name:   "heidi",
    Age:   18,
    Gender:   GenderMale,
    Weight: 150,
    FavoriteColor: []string{
        "red",
        "yellow",
    },
}