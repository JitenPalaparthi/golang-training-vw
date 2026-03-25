package main

import (
	"errors"
	"fmt"
)

func main() {

	slice1 := []int{10, 11, 12, 13, 14}

	sum := GetSumByAddElems(slice1, 15, 16, 17, 18, 19, 20)

	fmt.Println("sum:", sum, "slice1:", slice1)

	sum, err := GetSumByAddElemsP(&slice1, 15, 16, 17, 18, 19, 20)
	if err != nil {
		println(err.Error())
	}

	fmt.Println("Using pointer way", "sum:", sum, "slice1:", slice1)

	sum, slice1 = GetSumByAddElemsR(slice1, 15, 16, 17, 18, 19, 20) // technically we return the slice which has been appended in side the function
	// to the outside variable by copying back to the outside slice. So the header is copied

	fmt.Println("sum:", sum, "slice1:", slice1)

}

func GetSumByAddElems(slice []int, elems ...int) (sum int) {
	slice = append(slice, elems...)
	for _, v := range slice {
		sum += v
	}
	return sum
}

func GetSumByAddElemsR(slice []int, elems ...int) (sum int, sls []int) {
	slice = append(slice, elems...)
	for _, v := range slice {
		sum += v
	}
	return sum, slice
}

// passing slice as a refernece
func GetSumByAddElemsP(slice *[]int, elems ...int) (sum int, err error) {
	if slice == nil {
		return 0, errors.New("nil slice")
	}

	*slice = append(*slice, elems...)
	for _, v := range *slice {
		sum += v
	}
	return sum, nil
}
