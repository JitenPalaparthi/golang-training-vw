package main

func main() {

	var ptr1 *int

	if ptr1 == nil {
		println("nil pointer ,which does not hold any address")
	}

	var num1 int = 9898943
	ptr1 = &num1 // not nil now, bcz it holds the address

	if ptr1 != nil {
		*ptr1 = 100
	}

	println(num1)

	var ptr2 **int = &ptr1

	**ptr2 = 200

	println(num1)

	//ptr1 += 8

	ptr3 := new(int)
	*ptr3 = 100
	println(*ptr3)

	ptr4 := new("Hello World") // 16 bytes

	println(*ptr4)

	num := 100
	Incr(num)
	println("call by value:", num)

	IncrP(&num)
	println("call by reference:", num)

	*&num = 300 // technically it is valid , even we dont need to do that .
	println(num, *&num)

}

// pointer is a variable that holds the address
// pointer can be nil, there is no null in Go..It is only nil
// go does not support pointer arthimetic directly

// General problems with pointers
// -------------------------------
// Null pointer dereference
// Double Free
// Use after free
// Dangling pointer
// Memory leak

func Incr(num int) { // call / pass by value
	num++
}

func IncrP(num *int) { // call / pass by ref
	if num != nil {
		*num++
	}
}
