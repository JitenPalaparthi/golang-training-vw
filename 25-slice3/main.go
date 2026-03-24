package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 5, 10)
	FillSlice(slice1)
	slice2 := slice1
	PrintHeader("slice1", slice1)
	PrintHeader("slice2", slice2)
	slice1 = append(slice1, 99)

	// What is the header of slice1 and slice2?
	// Header is changed but not the allocation
	PrintHeader("slice1", slice1)
	PrintHeader("slice2", slice2)
	slice1[0] = 9999
	slice1[5] = 5555
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2", slice2)

	arr1 := [5]int{11, 234, 34, 47, 5}
	slice3 := arr1[:] // it converts from array to slice, both are to the same pointer at this point of time
	// slice3 has a header that is created based on arr1,
	// ptr: &arr[0] len:len(arr),cap: len(arr)
	// an array can be convertd to a slice
	PrintHeader("slice3", slice3)
	fmt.Println("Address arr1:", &arr1[0])
	max, min, _ := GetMaxAndMin(arr1[:]) // write funcs/methods using slice as params in majority of the cases
	fmt.Println("max:", max, "min:", min)

	slice4 := make([]int, 3)
	copy(slice4, slice3) // only values are copied , kind of deep copy .. headers are diffenet
	fmt.Println("slice4:", slice4)
}

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(100)
	}
}
func PrintHeader(name string, slice []int) {
	println("-----------------------")
	println(name)
	fmt.Println(slice)
	if slice != nil && len(slice) > 0 {
		fmt.Println("Ptr:", &slice[0])
	}
	fmt.Println("Len:", len(slice))
	fmt.Println("Cap:", cap(slice))
	println("-----------------------")
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
