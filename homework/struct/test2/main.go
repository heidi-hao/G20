package main

import "fmt"

func main() {

	book := Book{
		Title:       "不回来的你",
		Author:      Author{Name: "heidi", Emain: "1498263520@qq.com"},
		Category:    Category{Genre: "2段", Level: "初级"},
		PublishYear: 2002,
		Price:       25,
	}

	fmt.Println(book.Author)
	fmt.Println(book.Category)
	fmt.Println(book.Name)

}

type Author struct {
	Name  string
	Emain string
}

type Category struct {
	Genre string
	Level string
}

type Book struct {
	Title string
	Author
	Category
	PublishYear int
	Price       float64
}
