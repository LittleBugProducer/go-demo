package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string `json:"name"`
}

var p *Person
var on sync.Once

func getSinglePerson() *Person {
	on.Do(func() {
		p = &Person{Name: "xiaohong"}
	})
	return p
}

func main() {
	xiaohong := getSinglePerson()
	fmt.Println("name:", xiaohong.Name)
}
