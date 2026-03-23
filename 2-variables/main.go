package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"reflect"

	v1 "math/rand"
)

func main() {

	var num1 int = 1231231

	num1 = 34334

	var num2 int8 // zero value

	var float1 float32 = 123.34

	var ( // group of variables
		float2, float3 float64 = 12.123, 3234324324234.234532423434324
		str3                   = "Hello World" // Type inference
	)

	var ok1 = true // Type inference

	var str1 string = "hello World"

	var num3 = 40 // type inference , int, it takes 8 bytes

	ok2 := true // shorthand delcaration

	num4, num5, ok3 := 23, 989.34, false // shorthand declaration of multiple variables

	num6 := 32312 // shorthand declaration

	str2 := "Hello Go learners!" // shorthand declaration

	fmt.Println(num1, num2, num3, num4, num5, num6, str1, str2, str3, float1, float2, float3, ok1, ok2, ok3)

	fmt.Printf("%.2f\n", float3)

	num1 = 123 // not a new variable, just a mutation

	fmt.Printf("num1 in decimal %d num1 in binary %b num1 in hexa %x\n", num1, num1, num1)

	num7 := 0b1111011
	num8 := 0x7b

	fmt.Printf("num7 in decimal %d num7 in binary %b num7 in hexa %x\n", num7, num7, num7)
	fmt.Printf("num8 in decimal %d num8 in binary %b num8 in hexa %x\n", num8, num8, num8)

	fmt.Printf("num8 type: %T num8 in decimal %d num8 in binary %b num8 in hexa %x\n", num8, num8, num8, num8)

	fmt.Println("Type of num7:", reflect.TypeOf(num7))

	num9 := rand.IntN(99999999) // more efficient , performent

	num10 := v1.Intn(99999999)

	fmt.Println(num9, num10)

	println("sin:", math.Sin(23123))

}

/* Numbers
int,uint, int8,uint8,int16,uint16,int32,uint32,int64,uint64,float32,float64,complex64,complex128,uintptr
Alias types: rune, byte
*/

// Boolean --> bool
// String --> string
// Special --> any/interface{}
