package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 10)
	FillSlice(slice1)
	fmt.Println(slice1)
	PrintHeader("slice1", slice1)

	slice2 := slice1 // this is header copy
	PrintHeader("slice2", slice2)

	slice2[0] = 9999 // change of slice2 changes the slice1 as well
	fmt.Println("slice1:", slice1)
	slice1 = append(slice1, 99) // extending a slice
	// append changes the header of the slice. Either only the length if value to be appended is
	// with in the cap, or all three of them like ptr, len and cap
	// what is the length beofre append? 10
	// what is the capasity before append? 10
	// when append --> a new element is added to the slice1
	// does it have the capasity --> No
	// It would double the capasity --> 10. to 20
	// reallocating the meory and copying all elemenbe from old location to the new location

	PrintHeader("slice1", slice1)

	slice2[1] = 8888
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

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
