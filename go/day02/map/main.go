package main

import "fmt"

func main() {
	personList := []Person{}
	personList = append(personList, Person{
		Name: "heidi",
		Age:  18,
	})
	personList = append(personList, Person{
		Name: "abby",
		Age:  20,
	})

	for _, person := range personList {
		if person.Name == "heidi" {
			fmt.Println("find heidi:", person)
			break
		}
	}

	//如果我这个切片非常大 上百万 上千万
	//通过切片找一个元素 效率非常低
	//这时可以使用map来优化查找效率

	//使用map来存储人员信息 key 是姓名 value 是Person对象
	personMap := map[string]Person{}
	personMap["heidi"] = Person{
		Name: "heidi",
		Age:  22,
	}
	personMap["abby"] = Person{
		Name: "abby",
		Age:  21,
	}
	fmt.Println("通过map直接找到heidi:", personMap["heidi"])

	//创建和初始化
	//原理：使用make创建map 指定初始容量可以提高性能。也可以使用字面量初始化
	//使用make创建
	m1 := make(map[string]int)     //空map
	m2 := make(map[string]int, 10) //指定初始容量

	//字面量初始化
	m3 := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	//访问元素
	//原理：通过键访问值，如果键不存在，返回零值。可以使用第二个返回值检查键是否存在
	m4 := map[string]int{"a": 1, "b": 2}

	value := m4["a"]
	fmt.Println(value)
	if key, value := m4["c"]; value {
		fmt.Println("exist:", key)
	} else {
		fmt.Println("not exist")
	}

}

type Person struct {
	Name string
	Age  int
}
