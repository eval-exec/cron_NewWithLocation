package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func init() {

	xlocationLosangeles, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(xlocationLosangeles)
	fmt.Println(time.Now().In( xlocationLosangeles ))

	c := cron.NewWithLocation( xlocationLosangeles )

	err := c.AddFunc("*/1 */1 6 * * *", func() { fmt.Println("Every second ,every minite", time.Now().In(xlocationLosangeles)) })

	if err != nil {
		panic(err)
	}

	c.Start()

}

func main() {

	time.Sleep(10000 * time.Second)

}
