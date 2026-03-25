package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var emt1, emt2 Empty1
	var emt3, emt4 Empty2

	fmt.Println(emt1, emt2, emt3, emt4)
	println(&emt1, &emt2, &emt3, &emt4) // zerobase
	// 0x65fd6a060eb8 0x65fd6a060eb8 0x65fd6a060eb8 0x65fd6a060eb8--> This is a zerobase
	// ee7
	slice1 := []int{} // ptr:0x65fd6a060eb8 len:0 cap:0
	println(&slice1)

	uintptr1 := uintptr(unsafe.Pointer(&slice1))

	sliceheader := (*[3]int)(unsafe.Pointer(uintptr1))

	fmt.Printf("0x%x\n", sliceheader[0])
}

type Empty1 struct{} // Empty struct. What does it mean? it does not contain any memory
type Empty2 struct{}
