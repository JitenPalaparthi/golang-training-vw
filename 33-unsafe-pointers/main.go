package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [5]int{10, 11, 12, 13, 14}

	ptr1 := &arr1[0]
	// cannot perform arth operation on pointers but can perform on numbers
	// So what if a poiter is converted as a number perform arthm operation and convert it back to the pointer

	unsafeptr1 := unsafe.Pointer(ptr1) // 1
	uintptr1 := uintptr(unsafeptr1)    // 4
	uintptr1 += 8                      // perform arth ops

	ptr2 := (*int)(unsafe.Pointer(uintptr1)) //3 and 2

	println(*ptr2)

	println("using pointer arth on arrays")
	uptr := uintptr(unsafe.Pointer(&arr1[0]))
	for i := 0; i < 50; i++ {
		ptr := *(*int)(unsafe.Pointer(uptr))
		print(ptr, "  ")
		uptr = uptr + unsafe.Sizeof(arr1[0]) // instead of adding 8 .. just get the value from the application
	}

	str1 := "Hello World"
	// ptr , Len
	uptr1 := uintptr(unsafe.Pointer(&str1))

	strheader := (*[2]int)(unsafe.Pointer(uptr1))

	(*strheader)[1] = 20 //dereference to 20

	// ---ptr---- ---len----

	fmt.Println(str1)
	fmt.Println(len(str1))
	ptr := strheader[0]
	fmt.Printf("0x%x\n", ptr)

	// uintptr2 := unsafe.Pointer(&strheader[0])
	// fmt.Println(*(*byte)(unsafe.Pointer(uintptr2)))

}

// pointer arth is not allowed directly in go
/*
1. A pointer value of any type can be converted to a unsafe.Pointer.
2. A unsafe.Pointer can be converted to a pointer value of any type.
3. A uintptr can be converted to a unsafe.Pointer.
4. A unsafe.Pointer can be converted to a uintptr.
*/
