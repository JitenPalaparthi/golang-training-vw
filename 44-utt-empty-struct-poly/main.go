package main

import (
	"fmt"
	"math"
)

func main() {

	var emt1 Empty1
	var emt2 Empty2

	emt1.WhoAmI()
	emt2.WhoAmI()

	a1 := A1{}.Area(23.4)
	a2 := A2{}.Area(12.3, 45.23)

	fmt.Println(math.Ceil(a1), math.Round(a2))

}

type Empty1 struct{} // Empty struct. What does it mean? it does not contain any memory
type Empty2 struct{}

func (Empty1) WhoAmI() {
	println("I am Empty1 struct ")
}

func (Empty2) WhoAmI() {
	println("I am Empty12struct ")
}

type A1 struct{}
type A2 struct{}

func (A1) Area(s float32) float64 {
	return float64(s * s)
}

func (A2) Area(l, b float32) float64 {
	return float64(l * b)
}
