package main

import "fmt"

func main() {

	func() {
		println("hello World")
	}() // executor

	fn1 := func(a, b int) int {
		return a + b
	} // no exector.. a function is stored in fn1

	sum := fn1(10, 20)
	fmt.Println(sum)

	var fn2 func(int, int) int // function can be created as a variable

	// var st1 struct{ id int } struct can also be created as a variable

	fn2 = func(a, b int) int {
		return a - b
	}

	sub := fn2(20, 15)
	fmt.Println(sub)

	fn2 = mul

	m := fn2(10, 5)
	println(m)

	num := 10

	r := func() int {
		return num * num
	}()

	println(r)

	a := calc(10, 8, func(i1, i2 int) int {
		return i1 + i2
	})

	println(a)

	b := calc(10, 8, fn1) // this can accept as fn1 bcz fn1 satisfies this signature func(int,int)int

	println(b)

	c := calc(10, 3, mul)
	println(c)
}

func mul(a, b int) int {
	return a * b
}

func calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// a function without name is called as anonymous function
//
