//大写开头就必须加上注释  这和谁说理去
package main

import "fmt"

//Celsius is float64
type Celsius float64

// Fahrenheit is float64
type Fahrenheit float64

const (
	//AbsoluteZeroC is const
	AbsoluteZeroC Celsius = -273.15
	//FreezingC is const
	FreezingC Celsius = 0
	//BoilingC is consts
	BoilingC Celsius = 100
)

//CToF is fuc
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//FToC is fuc
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Println(CToF(220.0))
}
