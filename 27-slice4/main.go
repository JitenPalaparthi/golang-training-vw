package main

import "fmt"

func main() {

	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20} // len:11 cap:11
	slice2 := slice1
	slice3 := slice1[:5]
	slice4 := slice1[3:8]
	slice5 := slice1[5:]

	PrintHeader("slice1", slice1)
	PrintHeader("slice2", slice2)
	PrintHeader("slice3", slice3)
	PrintHeader("slice4", slice4)
	PrintHeader("slice5", slice5)

	slice6 := slice1[:]
	slice6 = append(slice6, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)
	PrintHeader("slice6", slice6)
	// [10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
	// to delete 20
	slice7 := append(slice6[:10], slice6[11:]...)
	fmt.Println(slice7)

	sum := GetSumByAddElems(slice6, 31, 32, 33, 34, 35)
	fmt.Println("Sum:", sum)

	fmt.Println(slice6)

	sum, slice6 = GetSumByAddElemsR(slice6, 31, 32, 33, 34, 35)
	fmt.Println("Sum:", sum)

	fmt.Println(slice6)
	clear(slice6) // make the slice back to zerovalue but not nil, the ptr, len, cap all remains but all values become zerovalues
	fmt.Println(slice6)
}

// 0x54ecd4fec000
// 0x54ecd4ff0000

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

// When ever append is used inside a method or a function, be extra causious
// bxz appnd changes the headr of the slice which would be propagated out side slice
// since slice is only a header copy
