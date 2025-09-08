package main

import (
	"fmt"
)

type Celsius float64

func (c Celsius) ToF() float64 { return float64(c)*9/5 + 32 }

func main() {

	var temperatura Celsius = 34

	fmt.Println(Celsius.ToF(temperatura))

}
