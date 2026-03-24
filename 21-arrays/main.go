package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {

	var arr1 [5]int // zero values [0 0 0 0 0]

	fmt.Println(arr1)

	for i := range arr1 {
		arr1[i] = rand.IntN(100)
	}
	fmt.Println(arr1)

	arr2 := [5]int{10, 11, 12, 13, 14} // short hand declaration
	println("len of arr2:", len(arr2))
	arr3 := [...]int{123, 2, 34, 67, 6, 4, 98, 6, 7, 8, 34, 56, 65} // the lengh is evaluated at compile time by the compiler
	println("len of arr3:", len(arr3))
	// arr3[14] = 123

	// for i := 0; i <= len(arr3); i++ {
	// 	println(arr3[i])
	// }

	fmt.Println("arr1 type:", reflect.TypeOf(arr1))
	fmt.Println("arr2 type:", reflect.TypeOf(arr2))
	fmt.Println("arr3 type:", reflect.TypeOf(arr3))

	sum1 := SumOf(arr1)
	sum2 := SumOf(arr2)
	println(sum1, sum2)

	//SumOf(([5]int)(arr3))

	sum3 := SumOf13(arr3)
	println(sum3)

	var arr4 [10]int

	mimlen := min(len(arr4), len(arr1)) // finding the min len, the smaller array

	for i := range mimlen {
		arr4[i] = arr1[i]
	}

	fmt.Println(arr4)

}

// collection of elements of same type
// the length of array is fixed, and known to the compiler(compile time)
// array cannot be extended, it is fixed length
// access the array with index
// the type of the array --> includes its length as well
// one array to another array cannot be type casted

// Type casting
// any number type can be to any number type
// float32 -- int in all combinations
// cannot type cast any nyumber type to bool
// cannot type cast from string to any number type
// can type cast from number type to string but it gives a char notation
// an array of one type cannot be typested to array of another type
// [5]int --cannot be to [6]int or 6[float32] or 5[float32]

func SumOf(arr [5]int) int { // practically very very rarely write functions with array as an input , rather use slice
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOf13(arr [13]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
