package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := "Hello World"
	str2 := "Hello World! How are you doing"

	// A string is a data structure, it is stored in stack(in most cases), but data is stored in differnet memories

	str3 := "A string is a data structure, it is stored in stack(in most cases), but data is stored in differnet memories"

	fmt.Println("Str1 Size:", unsafe.Sizeof(str1))
	fmt.Println("Str2 Size:", unsafe.Sizeof(str2))
	fmt.Println("Str3 Size:", unsafe.Sizeof(str3))

	a, b, ok, char := 10, 20, true, 'A'

	str4 := fmt.Sprint(" ", a, " ", b, " ", ok, " ", char) // Kind of a formatter
	fmt.Println("str4 Size:", unsafe.Sizeof(str4))

	str1 = "hello World, how are you doing, I am busy in learning Golang" // Youa re able to mutate but how to say string is immutable

	// str4 would not be stored in data segment

	// One thing to be understood in Go is nil
	// There is no null but nil
	// nil is a curious case in Go.. The nature of nil is purely depends on different constructurs
	// Where ever there is a pointer , directly or indirecly in the defined struct, can check nil except string
	// string is always check with ""

	// can arth ops ?, it is not arth , but it is concat

	str5 := str1 + str2 + str3 + str4 // Concatination
	fmt.Println(str5)

	if Greet == nil { // it be checked whether nil or not
		println("yes nil")
	}

}

func Greet() {
	println("Hello World")
}

// string , no need to check nil or not.. just check "" or not
// string is immutable
// string can store data : if it is string literal it stores in data segment

/*
String Header --> A structure
-------------
Len: int --> 8 bytes
Ptr: Address --> 8 bytes --> The address to the original data
*/
