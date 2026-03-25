package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice := make(intSlice, 10)
	slice.FillWithRand()

	fmt.Println(slice)

	max := slice.Max()
	min := slice.Min()

	fmt.Println("max:", max, "min:", min)

	slice1 := []int{10, 4, 65, 3, 46, 789, 4, 5, 76, 34, 33, 34, 3, 34, 54, 67}
	max = intSlice(slice1).Max()
	min = intSlice(slice1).Min()
	fmt.Println("max:", max, "min:", min)

	arr1 := [5]int{10, 124, 324, 34, 46}

	max = intSlice(arr1[:]).Max()
	min = intSlice(arr1[:]).Min()
	fmt.Println("max:", max, "min:", min)

	var empty struct{} = struct{}{}

	fmt.Println(empty)

}

type intSlice []int //it is a new type

func (slice *intSlice) FillWithRand() {
	if slice != nil {
		for i := range *slice {
			(*slice)[i] = rand.IntN(100)
		}
	}
}

func (slice intSlice) Min() int {
	min := 0
	if slice != nil && len(slice) > 0 {
		min = (slice)[0]
		for _, v := range slice {

			if v < min {
				min = v
			}

		}
	}
	return min
}

func (slice intSlice) Max() int {
	max := 0
	if slice != nil && len(slice) > 0 {
		max = (slice)[0]
		for _, v := range slice {

			if v > max {
				max = v
			}

		}
	}
	return max
}
