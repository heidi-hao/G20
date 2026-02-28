package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int) //空map
	addEntry(m, "new", 42)
	fmt.Println(m)

}

func addEntry(m map[string]int, key string, value int) {
	m[key] = value
}
