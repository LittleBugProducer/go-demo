package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	defer ants.Release()

	nonblockPrintArgs := func(i interface{}) {
		time.Sleep(time.Second * 1)
		fmt.Println("print args in ants with args:", i)
	}
	noblockPool, err := ants.NewPoolWithFunc(5, nonblockPrintArgs, ants.WithNonblocking(true))
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		err := noblockPool.Invoke(i)
		if err != nil {
			if errors.Is(err, ants.ErrPoolOverload) {
				fmt.Println("pool is full")
				break
			}
			fmt.Println("error:", err)
		}
	}

	time.Sleep(time.Second * 5)

	// printDemo := func() {
	// 	fmt.Println("print demo in ants")
	// }

	// for i := 0; i < 10; i++ {
	// 	ants.Submit(printDemo)
	// }

	// time.Sleep(time.Second * 5)

	// printArgs := func(i interface{}) {
	// 	fmt.Println("print args in ants with args:", i)
	// }

	// p, err := ants.NewPoolWithFunc(10, printArgs)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for i := 0; i < 15; i++ {
	// 	if err = p.Invoke(i); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	// time.Sleep(time.Second * 5)
}
