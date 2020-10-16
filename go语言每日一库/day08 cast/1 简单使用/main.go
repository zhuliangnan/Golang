package main

/*
cast是一个小巧、实用的类型转换库
*/
import (
	"fmt"
	"time"

	"github.com/spf13/cast"
)

func main() {
	fmt.Println("-------ToString---------")
	// ToString
	fmt.Println(cast.ToString("leedarjun"))        // leedarjun
	fmt.Println(cast.ToString(8))                  // 8
	fmt.Println(cast.ToString(8.31))               // 8.31
	fmt.Println(cast.ToString([]byte("one time"))) // one time
	fmt.Println(cast.ToString(nil))                // ""

	var foo interface{} = "one more time"
	fmt.Println(cast.ToString(foo)) // one more time

	fmt.Println("-------ToInt---------")
	// ToInt
	fmt.Println(cast.ToInt(8))     // 8
	fmt.Println(cast.ToInt(8.31))  // 8
	fmt.Println(cast.ToInt("8"))   // 8
	fmt.Println(cast.ToInt(true))  // 1
	fmt.Println(cast.ToInt(false)) // 0

	var eight interface{} = 8
	fmt.Println(cast.ToInt(eight)) // 8
	fmt.Println(cast.ToInt(nil))   // 0

	fmt.Println("-------ToTime---------")
	//ToTime
	now := time.Now()
	timestamp := 1579615973
	timeStr := "2020-01-21 22:13:48"

	fmt.Println(cast.ToTime(now))       // 2020-01-22 06:31:50.5068465 +0800 CST m=+0.000997701
	fmt.Println(cast.ToTime(timestamp)) // 2020-01-21 22:12:53 +0800 CST
	fmt.Println(cast.ToTime(timeStr))   // 2020-01-21 22:13:48 +0000 UTC

	d, _ := time.ParseDuration("1m30s")
	ns := 30000
	strWithUnit := "130s"
	strWithoutUnit := "130"

	fmt.Println(cast.ToDuration(d))              // 1m30s
	fmt.Println(cast.ToDuration(ns))             // 30µs
	fmt.Println(cast.ToDuration(strWithUnit))    // 2m10s
	fmt.Println(cast.ToDuration(strWithoutUnit)) // 130ns

	fmt.Println("-------ToSlice---------")
	sliceOfInt := []int{1, 3, 7}
	arrayOfInt := [3]int{8, 12}
	// ToIntSlice
	fmt.Println(cast.ToIntSlice(sliceOfInt)) // [1 3 7]
	fmt.Println(cast.ToIntSlice(arrayOfInt)) // [8 12 0]

	sliceOfInterface := []interface{}{1, 2.0, "darjun"}
	sliceOfString := []string{"abc", "dj", "pipi"}
	stringFields := " abc  def hij   "
	any := interface{}(37)
	// ToStringSliceE
	fmt.Println(cast.ToStringSlice(sliceOfInterface)) // [1 2 darjun]
	fmt.Println(cast.ToStringSlice(sliceOfString))    // [abc dj pipi]
	fmt.Println(cast.ToStringSlice(stringFields))     // [abc def hij]
	fmt.Println(cast.ToStringSlice(any))              // [37]

	fmt.Println("-------ToMap---------")
	m1 := map[string]string{
		"name": "darjun",
		"job":  "developer",
	}

	m2 := map[string]interface{}{
		"name": "jingwen",
		"age":  18,
	}

	m3 := map[interface{}]string{
		"name": "pipi",
		"job":  "designer",
	}

	m4 := map[interface{}]interface{}{
		"name": "did",
		"age":  29,
	}

	jsonStr := `{"name":"bibi", "job":"manager"}`

	fmt.Println(cast.ToStringMapString(m1))      // map[job:developer name:darjun]
	fmt.Println(cast.ToStringMapString(m2))      // map[age:18 name:jingwen]
	fmt.Println(cast.ToStringMapString(m3))      // map[job:designer name:pipi]
	fmt.Println(cast.ToStringMapString(m4))      // map[job:designer name:pipi]
	fmt.Println(cast.ToStringMapString(jsonStr)) // map[job:manager name:bibi]
}
