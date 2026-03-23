package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var any1 any
	var any2 interface{} // both any and interface{} are same

	if any1 == nil {
		println("any1 is nil")
	}
	fmt.Printf("Data in any1 :%v Type of any1:%T\n", any1, any1)

	fmt.Println("Size of Any:", unsafe.Sizeof(any1))

	any1 = 100
	fmt.Println("Size of Any:", unsafe.Sizeof(any1))

	any1 = true
	fmt.Println("Size of Any:", unsafe.Sizeof(any1))

	any1 = "Hello World"
	fmt.Println("Size of Any:", unsafe.Sizeof(any1))

	str1 := "Hello World, how are you doing"
	fmt.Println("Size of Any:", unsafe.Sizeof(any1))

	any1 = str1

	fmt.Println("Size of Any:", unsafe.Sizeof(any1))

	// any1 = Person{101, "Jiten"}

	any1 = Employee{101, "Jiten"}

	fmt.Println("Size of Any:", unsafe.Sizeof(any1))

	any1 = any2
	fmt.Println(any1, "Size of Any:", unsafe.Sizeof(any1))
	var any3 = any1

	fmt.Println(any3, "Size of Any:", unsafe.Sizeof(any3))

}

type Employee struct {
	ID   int
	Name string
}

// any is alias for interface{}
// can store any kind of data in any
// any can be nil
// how can any be type safe
/* Any Header
Data Pointer: Address of the original data --> 8 bytes
Type Pointer: Address to the type --> 8 bytes
*/
