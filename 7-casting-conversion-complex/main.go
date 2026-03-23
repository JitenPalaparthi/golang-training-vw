package main

import (
	"fmt"
	_ "math"
	"math/rand/v2"
	"reflect"
	"strconv"
)

func main() {

	num1 := rand.IntN(256)

	var num2 uint8 = uint8(num1)

	//(uint8)num1

	println(num1, num2)

	var num3 uint = 98989 // 110000010 10101101

	var num4 uint8 = uint8(num3) // 8 bytes cannot be fit into 1 byte

	fmt.Printf("%b\n", num3)

	fmt.Printf("%d\n", 0b10101101)
	fmt.Println(num4)

	var float1 = 87879.3434
	//fmt.Printf("%b\n", float1)

	var num5 uint = uint(float1)
	fmt.Println(num5)

	ok1 := true // 1 byte

	//var num6 uint8 = uint8(ok1)

	var num6 = BoolToByte(ok1)
	println(num6)

	char1 := 'A' // char but technically it is a number

	println(string(char1), char1)

	var char2 uint64 = '形'

	println(string(char2), char2)

	var num7 int = 24418

	println(string(num7))
	// "24418"

	str1 := fmt.Sprint(num7)
	str2 := strconv.Itoa(num7)
	fmt.Println("Type:", reflect.TypeOf(str1), str1)
	fmt.Println("Type:", reflect.TypeOf(str2), str2)

	//str3 := "65753453"
	//var num8 = int(str3)
	//strconv.Atoi(str3)

	a, b, c, d := Calc(20, 10)
	println(a, b, c, d)

	_, b, _, d = Calc(20, 10) // _ blank identifier

	println(b, d)

	_, _, _, d = Calc(95, 40)

	str3 := "657a53453"
	//var num8 = int(str3)
	num8, err := strconv.Atoi(str3)
	if err != nil { // there is an error , and it failed to convert
		println(err.Error())
	} else {
		println(num8)
	}

	str3 = fmt.Sprint(rand.IntN(9989899)) //"65753453"
	num9, _ := strconv.Atoi(str3)         // not a good practice . must check the error
	println(num9)

	// complex numbers

	c1 := 100 + 10.5i // The whole thing is a complex number

	var r1, im1 float32 = 100.4, 1.5
	c2 := complex(r1, im1) // complex64

	c3 := complex(100.4, 1.5) // complex128

	c4 := complex64(c1) + c2 + complex64(c3)
	println(c4)

	r, im := real(c4), imag(c4)
	println(r, im)
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func Calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// can casting hapen between all types --No
// All number types can be casted among themself
// a bool cannot be casted to other type, but converted to and from string
