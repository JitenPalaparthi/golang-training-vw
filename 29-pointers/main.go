package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [5]int{10, 11, 12, 13, 14}

	a, b, c, d, e := 10, 11, 12, 13, 14

	var arr2 [5]*int

	arr2[0], arr2[1], arr2[2], arr2[3], arr2[4] = &a, &b, &c, &d, &e

	var arr3 *[5]int // arr3 is just a pointer

	fmt.Println(arr1, arr2, arr3)
	fmt.Println("Size of arr3:", unsafe.Sizeof(arr3))
	arr3 = &arr1

	if arr3 != nil {
		println("iterating an array pointer")
		for i, v := range *arr3 {
			(*arr3)[i] = v * v
			// arr3[i]= v*v
			println(i, v)
		}
	}

	println("iterating array of pointers")
	for i, v := range arr2 {
		if v != nil {
			println(i, *v)
		}
	}

	fmt.Println(arr1, arr2, arr3)

}
