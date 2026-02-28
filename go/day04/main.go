package main

import "fmt"

func main() {
	var a = 10
	incrementPtr(&a)
	fmt.Println(a)

}

func incrementPtr(p *int) {
	*p++
}
