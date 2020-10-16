package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"log"
	"time"
)

/*
使用标准库time的时间运算需要先定义一个time.Duration对象，time库预定义的只有纳秒到小时的精度：
const (
  Nanosecond  Duration = 1
  Microsecond = 1000 * Nanosecond  --纳秒
  Millisecond = 1000 * Microsecond
  Second      = 1000 * Millisecond  -- time.Second
  Minute      = 60 * Second
  Hour        = 60 * Minute    --小时
)

其它的时长就需要自己使用time.ParseDuration构造了，而且time.ParseDuration不能构造其它精度的时间。
如果想要增加/减少年月日，就需要使用time.Time的AddDate方法：
需要注意的是，时间操作都是返回一个新的对象，原对象不会修改。carbon库也是如此
*/
func main() {
	now := time.Now()

	fmt.Println("now is:", now)

	fmt.Println("one second later is:", now.Add(time.Second))
	fmt.Println("one minute later is:", now.Add(time.Minute))
	fmt.Println("one hour later is:", now.Add(time.Hour))

	d, err := time.ParseDuration("3m20s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("3 minutes and 20 seconds later is:", now.Add(d))

	d, err = time.ParseDuration("2h30m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2 hours and 30 minutes later is:", now.Add(d))
	//还可以传递负数代表减少
	fmt.Println("3 days and 2 hours later is:", now.AddDate(0, 0, -3).Add(time.Hour*2))

	fmt.Println("--------------------carbon-----------------------")
	now2 := carbon.Now()

	fmt.Println("now is:", now2)

	fmt.Println("one second later is:", now2.AddSecond())
	fmt.Println("one minute later is:", now2.AddMinute())
	fmt.Println("one hour later is:", now2.AddHour())
	//其实给上面方法传入负数就表示减少
	fmt.Println("3 minutes and 20 seconds later is:", now2.AddMinutes(-3).AddSeconds(20))
	fmt.Println("2 hours and 30 minutes later is:", now2.AddHours(2).AddMinutes(30))
	fmt.Println("3 days and 2 hours later is:", now2.AddDays(3).AddHours(2))

	/*
		carbon还提供了：

		增加季度的方法：AddQuarters/AddQuarter，复数形式介绍一个表示倍数的参数，单数形式倍数为1；
		增加世纪的方法：AddCenturies/AddCentury；
		增加工作日的方法：AddWeekdays/AddWeekday，这个方法会跳过非工作日；
		增加周的方法：AddWeeks/AddWeek。
		其实给上面方法传入负数就表示减少，另外carbon也提供了对应的Sub*方法。
	*/
}
