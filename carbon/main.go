package main

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
)

func main() {
	// lang := carbon.NewLanguage()
	// lang.SetLocale("zh-CN")
	// c := carbon.SetLanguage(lang)
	// fmt.Println(carbon.Now().ToDateTimeString())
	// fmt.Println(c.Tomorrow().ToMonthString())
	// fmt.Println(c.Now().Season())
	// fmt.Println("Hello, World!")

	fmt.Println(carbon.Now().Layout(carbon.RFC3339Layout))
	fmt.Println(carbon.Now().Format(carbon.RFC3339Format))

	// not recommend
	fmt.Println(time.Now().Format(carbon.RFC3339MilliLayout))
	fmt.Println(time.Now().Format(carbon.RFC3339Layout))

}
