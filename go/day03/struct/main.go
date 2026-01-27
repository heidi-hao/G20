package main

import "fmt"

func main() {

	//slice 类型的初始化
	arr := []int{1, 2, 3}
	fmt.Println(arr)

	//map 类型的初始化
	m1 := map[string]int{
		"heidi": 18,
		"abby":  20,
	}
	fmt.Println(m1)

	//person 结构体类型的初始化
	p1 := Person{
		Name:   "heidi",
		Age:    18,
		Gender: GenderMale,
		Weight: 150,
		FavoriteColor: []string{
			"red",
			"yellow",
		},
	}
	fmt.Println(p1)

	p2 := Person{
		Name:   "abby",
		Age:    18,
		Gender: GenderFemale,
		Weight: 120,
		FavoriteColor: []string{
			"bule",
			"black",
		},
	}

	fmt.Println(p2)

	//访问结构体字段
	fmt.Println(p2.Age)

	p2.Age = 22
	fmt.Println(p2.Age)

	ChangePersonAge(&p2, 30)
	fmt.Println("ChangePersonAge", p2.Age)

}

type Person struct {
	Name          string
	Age           int
	Gender        Gender   //自定义枚举类型
	Weight        uint     //体重（非负）
	FavoriteColor []string //切片
}

// GenderXXX 表示枚举值
// Gender 是一种自定义类型 底层类型是int
// Gender != int

type Gender int //Gender 是一种 新类型 底层类型是int Gender ≠ int（类型安全）

// Sex == Gender
type Sex = Gender //sex是Gender的别名

// iota + const = 枚举 作用是：定义一组“性别枚举常量”，用有意义的名字代替 0、1，并且保证类型安全。
const (
	GenderMale Gender = iota //iota是常量计数器 从0开始递增
	GenderFemale
)

// ChangePersonAge 修改Person的年纪
func ChangePersonAge(p *Person, newAge int) {
	// 自动解开引用(*p)
	// * -> 解引用(Unbox)
	//(*p).age = newAge
	// p 指针类型，自动解引用
	// . 访问字段
	// 自动解引用
	// p.age == (*p).age
	p.Age = newAge
}

