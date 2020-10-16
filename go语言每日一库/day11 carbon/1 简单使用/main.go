/*
一线开发人员每天都要使用日期和时间相关的功能，各种定时器，活动时间处理等。
标准库time使用起来不太灵活，特别是日期时间的创建和运算。
carbon库是一个时间扩展库，基于 PHP 的carbon库编写。提供易于使用的接口
*/

package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"time"
)

func main() {

	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("Japan")
	fmt.Printf("Right now in Japan is %s\n", today)

	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())

	nextOlympics, _ := carbon.CreateFromDate(2016, time.August, 5, "Europe/London")
	nextOlympics = nextOlympics.AddYears(4)
	fmt.Printf("Next olympics are in %d\n", nextOlympics.Year())

	if carbon.Now().IsWeekend() {
		fmt.Printf("Happy time!")
	}
}
