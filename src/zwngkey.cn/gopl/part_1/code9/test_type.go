package main

import "fmt"

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	TFahren       Fahrenheit = 100
)

func main() {
	v := CToF(AbsoluteZeroC)
	fmt.Println(v)
	q := CToF(FreezingC)
	fmt.Println(q)
	a := FtoC(TFahren)
	fmt.Println(a)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
