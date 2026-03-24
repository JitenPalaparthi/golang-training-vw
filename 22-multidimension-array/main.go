package main

import "fmt"

func main() {

	//var arr2d [2][2]int

	arr2d := [2][2]int{{10, 11}, {12, 13}}

	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	// if arr2d == nil { // cant check nil on array
	// }

	// iterate using normal loop

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			print(arr2d[i][j], " ")
		}
		println()
	}

	println()
	for _, arr1 := range arr3d {
		for _, arr2 := range arr1 {
			for _, v := range arr2 {
				print(v, " ")
			}
			println()
		}
	}

	arr1 := [6]any{10, 13.4, true, float32(94.334), "Hello Wrold", int8(100)}

	sum := 0.0
	for _, a := range arr1 {
		switch v := a.(type) {
		case int:
			sum += float64(v)
		case uint:
			sum += float64(v)
		case int8:
			sum += float64(v)
		case uint8:
			sum += float64(v)
		case int16:
			sum += float64(a.(int64)) // can also assert and convert again to float64
		case uint16:
			sum += float64(v)
		case int32:
			sum += float64(v)
		case uint32:
			sum += float64(v)
		case int64:
			sum += float64(v)
		case uint64:
			sum += float64(v)
		case float32:
			sum += float64(v)
		case float64:
			sum += v
		}
	}

	fmt.Printf("Sum:%.3f\n", sum)

}
