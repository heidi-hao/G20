package main

import "fmt"

// 定义接口
type Animal interface {
	Speak() string //方法名
	GetName() string
}

// 定义不同的类型
type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Cow struct {
	Name string
}

// 实现接口方法
func (d Dog) Speak() string {
	return "汪汪汪"
}
func (d Dog) GetName() string {
	return d.Name
}

// 使用接口(多态)
func MakeSound(a Animal) {
	fmt.Printf("%s动物叫声: %s\n", a.GetName(), a.Speak())
}

func main() {
	dog := Dog{Name: "旺财"}
	MakeSound(dog)

}
