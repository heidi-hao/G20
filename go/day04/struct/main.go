package main

import (
	"fmt"
)

func main() {
	//创建了一个指针
	person := &Person{ //person的类型是：*Person
		Name: "heidi",
		age:  18, //age可以赋值是因为在同一个包里 小写是未导出字段(私有) 但是在同一个package main里，是可以访问的
	}

	//通过一个一个的函数 去处理数据的 编程方式 就是面向对象
	ChangePersonAge(person, 28)
	//修改体重 ChangeWeight 的函数
	fmt.Println(person.age)

	//面向对象有更好的封装 也贴近现实世界
	person.ChangePersonAge(20)
	fmt.Println("修改后的年纪:", person.age)
	fmt.Println(person.age)
	person.我的名字()

	//修改weight
	person.SetWeight(80)
	fmt.Println("修改后的体重:", person.Weight)

	fmt.Println("结构体与面向对象")
	//创建交通工具
	//能做的动作 -> 海豚跳(Jump)
	bike := &Bike{TranficTool{ //结构体复合字面量 匿名嵌入字段的初始化写法
		Name:  "bicycle",
		Speed: 20,
		Color: "black",
	}, PlayTool{}}

	//bike.TranficTool.Drive() 这个是先访问嵌入字段在调用它的方法 和下面的是等价的 这不是继承是组合
	bike.Drive()

}

// Person是一个表示人的结构体 包含姓名 年级 性别 体重
// 也是一个复合数据类型
type Person struct {
	Name string
	//不对外暴露的字段
	age           int
	Gender        Gender
	Weight        uint
	FavoriteColor []string
}

// Gender XXX表示枚举值
// Gender 是一种自定义类型 底层类型是int 这是Go很重要的类型安全机制
// Gender != int
type Gender int

// Sex == Gender
type Sex = Gender //sex和Gender是完全同一个类型

// 展开以后就是 GenderMale = 0 GenderFemale = 1 类型是Gender不是int
const (
	GenderMale Gender = iota
	GenderFemale
)

// 方法版本 只要方法里要修改对象，就用指针接收者 也就是为啥是*Person
func (p *Person) ChangePersonAge(newAge int) {
	p.age = newAge
}

// 修改的
// SetWeight(p *Person)
func (p *Person) SetWeight(newWeight uint) {
	p.Weight = newWeight
}

// 不涉及到对象修改的 把这个方法绑定给值也是可以的
// GetWeight(p Person)
// GetWeight 是值接收者
func (p Person) GetWeight() uint { //没有涉及任何的修改字段只是读取
	return p.Weight
}

func (p *Person) 我的名字() string {
	return p.Name
}

// 普通函数 传的是一个指针 所以这个函数真的修改了原来的person
func ChangePersonAge(p *Person, newAge int) {
	p.age = newAge
}

type TranficTool struct {
	// ``表示字符串 都是Struct Tag 结构体标签 lib 反射使用
	Name  string `json:"name"`
	Speed int    `json:"speed"`
	Color string `json:"color"`
	move  bool
}

// Drive 是交通工具的行驶方法 给类型TranficTool绑定了一个Drive方法 等价于TranficTool这个"对象"可以调用Drive
func (t *TranficTool) Drive() {
	fmt.Println("%s 正在以%d的速度前进\n", t.Name, t.Speed)
}

func (t *TranficTool) Stop() {
	fmt.Println("正在进行stop")
}

// PlayTool是一个表演工具 空结构体 这个类型本身没有字段 只是用来"挂行为"
type PlayTool struct{}

// 这是一个方法 不是普通函数 给PlayTool加了一个方法Play p是方法的接受者
func (p *PlayTool) Play() {
	fmt.Println("开始我的表演:后空翻")
}

func (b *PlayTool) Jump() {
	fmt.Println("海豚跳: 高级版")
}

// Bike 是自行车
type Bike struct { //嵌入字段 TranficTool PlayTool
	//嵌入了交通工具 继承了交通工具的属性和方法(可以上班)
	TranficTool
	//嵌套了PlayTool，继承了PlayTool的方法(可以表演)
	PlayTool
}
