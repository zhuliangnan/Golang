package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"log"
	"time"
)

func main() {

	/*
		Go 标准库time创建某个时区的时间，需要先加载时区
	*/
	loc, err := time.LoadLocation("Japan")
	if err != nil {
		log.Fatal("failed to load location: ", err)
	}

	d := time.Date(2020, time.July, 24, 20, 0, 0, 0, loc)
	fmt.Printf("The opening ceremony of next olympics will start at %s in Japan\n", d)
	/*
		Go 标准库time创建某个时区的时间，需要先加载时区
	*/

	c, err := carbon.Create(2020, time.July, 24, 20, 0, 0, 0, "Japan")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The opening ceremony of next olympics will start at %s in Japan\n", c)
}
