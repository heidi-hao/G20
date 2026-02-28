package main

import "fmt"

func main() {
	emp := Employee{
		Name:       "alice",
		Id:         1005,
		Salary:     10000,
		Department: "A",
	}

	fmt.Println(emp.GetInfo())
	emp.RaiseSalary(10)
	fmt.Println(emp.Salary)
	emp.ChangeDepartment("D")
	fmt.Println(emp.GetInfo())

}

type Employee struct {
	Name       string
	Id         int
	Salary     float64
	Department string
}

func (e *Employee) RaiseSalary(percent float64) {
	e.Salary = e.Salary * (1 + percent/100)
}

func (e Employee) GetInfo() string {
	return fmt.Sprintf("name:%s,id:%d,salary:%.2f,department:%s", e.Name, e.Id, e.Salary, e.Department)
}

func (e *Employee) ChangeDepartment(dept string) {
	e.Department = dept
}
