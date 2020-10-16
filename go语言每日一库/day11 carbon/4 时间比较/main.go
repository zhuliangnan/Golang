package main

/*
标准库time可以使用time.Time对象的Before/After/Equal判断是否在另一个时间对象前，后，或相等。
carbon库也可以使用上面的方法比较时间。除此之外，它还提供了多组方法，每个方法提供一个简短名，一个详细名：

Eq/EqualTo：是否相等；
Ne/NotEqualTo：是否不等；
Gt/GreaterThan：是否在之后；
Lt/LessThan：是否在之前；
Lte/LessThanOrEqual：是否相同或在之前；
Between：是否在两个时间之间。
另外carbon提供了：

判断当前时间是周几的方法：IsMonday/IsTuesday/.../IsSunday；
是否是工作日，周末，闰年，过去时间还是未来时间：IsWeekday/IsWeekend/IsLeapYear/IsPast/IsFuture。
*/
import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func main() {
	t1, _ := carbon.CreateFromDate(2010, 10, 1, "Asia/Shanghai")
	t2, _ := carbon.CreateFromDate(2011, 10, 20, "Asia/Shanghai")

	fmt.Printf("t1 equal to t2: %t\n", t1.Eq(t2))
	fmt.Printf("t1 not equal to t2: %t\n", t1.Ne(t2))

	//Gt/GreaterThan：是否在之后；
	fmt.Printf("t1 greater than t2: %t\n", t1.Gt(t2))
	//Lt/LessThan：是否在之前；
	fmt.Printf("t1 less than t2: %t\n", t1.Lt(t2))

	t3, _ := carbon.CreateFromDate(2011, 1, 20, "Asia/Shanghai")
	//Between：是否在两个时间之间。
	fmt.Printf("t3 between t1 and t2: %t\n", t3.Between(t1, t2, true))

	now := carbon.Now()
	//是否是工作日
	fmt.Printf("Weekday? %t\n", now.IsWeekday())
	//是否是周末
	fmt.Printf("Weekend? %t\n", now.IsWeekend())
	//闰年
	fmt.Printf("LeapYear? %t\n", now.IsLeapYear())
	//现在时间
	fmt.Printf("Past? %t\n", now.IsPast())
	//未来时间
	fmt.Printf("Future? %t\n", now.IsFuture())

	fmt.Println("---------使用carbon计算两个日期之间相差多少秒、分、小时、天---------------")
	vancouver, _ := carbon.Today("Asia/Shanghai")
	london, _ := carbon.Today("Asia/Hong_Kong")
	fmt.Println(vancouver.DiffInSeconds(london, true)) // 0

	ottawa, _ := carbon.CreateFromDate(2000, 1, 1, "America/Toronto")
	vancouver, _ = carbon.CreateFromDate(2000, 1, 1, "America/Vancouver")
	fmt.Println(ottawa.DiffInHours(vancouver, true)) // 3

	fmt.Println(ottawa.DiffInHours(vancouver, false)) // 3
	fmt.Println(vancouver.DiffInHours(ottawa, false)) // -3

	t, _ := carbon.CreateFromDate(2012, 1, 31, "UTC")
	fmt.Println(t.DiffInDays(t.AddMonth(), true))  // 31
	fmt.Println(t.DiffInDays(t.SubMonth(), false)) // -31

	t, _ = carbon.CreateFromDate(2012, 4, 30, "UTC")
	fmt.Println(t.DiffInDays(t.AddMonth(), true)) // 30
	fmt.Println(t.DiffInDays(t.AddWeek(), true))  // 7

	t, _ = carbon.CreateFromTime(10, 1, 1, 0, "UTC")
	fmt.Println(t.DiffInMinutes(t.AddSeconds(59), true))  // 0
	fmt.Println(t.DiffInMinutes(t.AddSeconds(60), true))  // 1
	fmt.Println(t.DiffInMinutes(t.AddSeconds(119), true)) // 1
	fmt.Println(t.DiffInMinutes(t.AddSeconds(120), true)) // 2
}
