package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func init() {

	xlocationLosangeles, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println("location is ",xlocationLosangeles)
	fmt.Println("time in losangels is :",time.Now().In( xlocationLosangeles ))
	fmt.Println("time,NOW() is :",time.Now())


	c := cron.NewWithLocation( xlocationLosangeles )


	err := c.AddFunc("*/1 */1 16 * * *", func() { fmt.Println("Los_Angels:Every second ,every minite", time.Now().In(xlocationLosangeles)) })

	if err != nil {
		panic(err)
	}
	cTimenow := cron.New()
	err = cTimenow.AddFunc("*/1 */1 7 * * *", func() { fmt.Println("time_Now():Every second ,every minite", time.Now()) })


	if err != nil {
		panic(err)
	}

	c.Start()
	cTimenow.Start()

}

func main() {

	time.Sleep(10000 * time.Second)

}
