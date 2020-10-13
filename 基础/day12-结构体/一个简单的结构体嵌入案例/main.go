package main

import "fmt"

func main() {
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}

	//得意于匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径
	//又可以愉快的玩耍了
	/*		var w Wheel
			w.X = 8  // w.Circle.Point.X = 8
			w.Y = 8  // w.Circle.Point.Y = 8
			w.Radius =5 // w.Circle.Radius	= 5
			w.Spokes =20*/
	/*	var w Wheel
		w = Wheel{8, 8, 5, 20}
		w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}*/
	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	/*	w = Wheel{
		Circle:Circle{
			Point: Point{X:8, Y:8} ,
			Radius: 5,
		},
		Spokes: 20,
	}*/

	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)

}
