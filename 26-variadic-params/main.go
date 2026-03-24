package main

import "fmt"

func main() {
	fmt.Println("hi", true, "how ", 1231, "you") // can give any number of values,
	slice1 := []int{}
	slice1 = append(slice1, 10, 20, 30, 40, 50) // can append any number of values
	// any number of values are variadic parmeters
	arr1 := [5]int{43, 5, 34, 5, 7}
	s := SumOf()
	println("sum:", s)
	s = SumOf(10, 20)
	println("sum:", s)
	s = SumOf(10, 20, 4, 56, 45, 5, 8, 6, 4, 5, 57, 67, 4, 45, 57, 68, 234, 54, 46, 56, 656, 5634, 45)
	println("sum:", s)
	// can pass a slice

	s = SumOf(slice1...)
	println("sum:", s)

	// can pass array
	s = SumOf(arr1[:]...)
	println("sum:", s)

	//var v ...int // this is not possible, only be used in funcs/methods
}

// a variadic param means can pass any number of arguments (any number means subjective to the space)
// variadic can only be used in functions or methods
// cannot be used a variable
// a variadic must be the only param or the last parameter in a func/method
// a slice can be passed as a variadic arg
// an array can also be converted as slice and be passed as a variadic argument
// nums ...int variadic parameter
// can use range loop on variadic param
// variadic arguments are 0-n

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
