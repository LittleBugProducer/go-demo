package main

import "fmt"

var a *Person

func init() {

	a = &Person{"213"}
}

type Person struct {
	Name string
}

func main() {

	fmt.Println(5)

}
