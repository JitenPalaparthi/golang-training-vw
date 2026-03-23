package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {

	var num1 = rand.IntN(4343)

	var any1 any = num1
	fmt.Println("Type:", reflect.TypeOf(any1), "Value:", any1)

	var num2 int = any1.(int) // unboxing kind of --> type assertion
	println(num2)

	any1 = uint8(199) // uint

	//var num3 int = any1.(int)
	//println(num3)

	//var num3 int

	num3, ok := any1.(int) // ok is bool -> true is succefully asserted , false means failed to assert
	if ok {
		println(num3)
	} else {
		num3, ok := any1.(uint8)
		if ok {
			println(num3)
		} else {
			println("cound not be asserted to int")
		}
	}

}
