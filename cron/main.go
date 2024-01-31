package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		fmt.Println(time.Now(), "run every 5 second")
	})
	c.AddFunc("4 * * * ?", func() {
		fmt.Println(time.Now(), "run minute 4 every hour")
	})
	c.Start()
	select {}
}
