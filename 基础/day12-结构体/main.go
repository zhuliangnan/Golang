package main

import (
	"fmt"
	"time"
)

func main() {

	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time //若未初始化则表示UTC时间，公元1年，1月，1日，0时，0分，0秒
		Position  string
		Salary    int
		ManagerID int
	}
	var dilbert Employee
	dilbert.Name = "dilbert"
	/*
		dilbert.Salary = 5000
		fmt.Println(dilbert)  // {0 dilbert  0001-01-01 00:00:00 +0000 UTC  5000 0}*/

	name := &dilbert.Name
	*name = "new" + *name
	fmt.Println(*name)

	type Point struct {
		X, Y int
	}
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // false
	fmt.Println(p == q)                   //false

}
