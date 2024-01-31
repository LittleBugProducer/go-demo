package main

import (
	"fmt"
	"unsafe"
)

func main() {
	t := struct{}{}
	fmt.Println(unsafe.Sizeof(t))
	fmt.Printf("t pos: %p\n", &t)

	str := "123"
	fmt.Printf("str pos: %p\n", &str)

	t1 := struct{}{}
	fmt.Println(unsafe.Sizeof(t1))
	fmt.Printf("t1 pos: %p", &t1)

	stuMap := make(map[string]struct{})
	stuMap["1"] = struct{}{}
}
