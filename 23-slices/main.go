package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"unsafe"
)

func main() {
	var slice1 []int // this is a declaration of slice and no size to be given like array
	// the above is only a declarion but not instantiation
	// ptr:nil lan:0 cap:0
	fmt.Println("size of slice1:", unsafe.Sizeof(slice1))
	if slice1 == nil {
		println("slice1 is nil")
		slice1 = make([]int, 10) // to allocate memory for 10 elements
		// ptr: 0x140011 len:10 cap:10
	}

	fmt.Println(slice1)
	slice2 := make([]int, 10, 20)
	// // ptr: 0x140011 len:10 cap:20
	fmt.Println(slice2)

	FillSlice(slice1) // what happen to this function, when you pass the slice1
	fmt.Println(slice1)
	FillSlice(slice2)
	fmt.Println(slice2)
	max, min, _ := GetMaxAndMin(slice1)
	println("max:", max, "min:", min)
	max, min, _ = GetMaxAndMin(slice2)
	println("max:", max, "min:", min)

	var slice3 []int
	max, min, err := GetMaxAndMin(slice3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		println("max:", max, "min:", min)
	}

	slice4 := []int{10, 8, 45, 65, 89, 4, 34, 76, 2, 87, 57, 98, 34, 90, 98, 23} // This is a slice
	max, min, _ = GetMaxAndMin(slice4)
	println("max:", max, "min:", min)

	slice5 := []int{} // The most trickest part of slice
	// ptr: zerobase ptr(dummy ptr), len:0 cap:0
	// zerobase pointer is a pointer that is created by compier and use it as a dummy pointer in some cases
	// this slice is not nil, that means there is some pointer
	if slice5 == nil {
		println("slice5 is nil")
	} else {
		fmt.Println("not nil slice:", slice5)
	}

	max, min, err = GetMaxAndMin(slice5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		println("max:", max, "min:", min)
	}
}

// slice is growable array , but the concept is not exactly array little different
// slice can be nil
// if slice is instantiated but not assigned any values then there are zero values for the slice
// a special function called make is used to instantiate a slice
// slice has a header
/*
Slice Header
------------
Ptr: Address to the original data 8 bytes
Len: int 8 bytes
Cap; int 8 bytes
// 8 bytes on 64 bit machines .. 4 on 32 bit machines
*/

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(100)
	}
}

func GetMaxAndMin(slice []int) (max, min int, err error) {
	if slice == nil {
		return 0, 0, errors.New("nil slice")
	}
	if len(slice) == 0 {
		return 0, 0, errors.New("slice has no elements")
	}

	max, min = slice[0], slice[0]
	for _, v := range slice {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return max, min, nil
}
