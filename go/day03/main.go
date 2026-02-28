package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "alice", Age: 20}

	//序列化为JSON字节
	data, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
