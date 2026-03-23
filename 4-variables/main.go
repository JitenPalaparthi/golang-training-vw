package main

import (
	"fmt"
	"reflect"
)

type char = int32 // alias

func main() {
	// There is no type called char in go
	// rune  --> not a separate variable , it is an alias for the exiting int32

	// char in go is just a number.

	var char1 int32 = 'A'
	var char2 rune = 'B' // rune is used only for convenction purpose but there is not real type called rune
	var char3 uint8 = 'C'
	var char4 rune = '象'
	var char5 uint64 = '形'
	var char6 = '😍'
	var char8 byte = 'D' // just an alias for uint8

	fmt.Println(char1, char2, char3, char4, char5, char6, char8)

	fmt.Println("type of char2:", reflect.TypeOf(char2))

	var char7 char = '😍'

	fmt.Printf("Type of char7:%T\n", char7)

	str1 := "Hello World" // collection of bytes
	str2 := "你好 😍"        // collection of bytes
	fmt.Println(str1, str2, len(str1), len(str2))
	// What is a string --> String is collection of bytes --> note: not chars in Go

	sum := char1 + char2 + char4
	println(sum)

}

// char is a utf-8 sequence -> 4 bytes
