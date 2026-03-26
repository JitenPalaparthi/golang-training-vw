package main

import "fmt"

func main() {

	defer fmt.Println("End of main")           // 1
	defer fmt.Println("Another defer in main") // 2

	defer func() {
		sum := fib(5)
		println("sum of fib:", sum)
	}()

	defer func() { //func1
		defer println("end of func1") // 1
		fmt.Println("Start of func1")
	}()
	fmt.Println("Start of main")

	// another defer example

	str := "Hello World"

	for _, v := range str {
		defer print(string(v), " ")
	}
	println()

	a, b := 10, new(10)
	defer fmt.Println("defer a:", a, "defer b:", *b)
	defer myprint(a, b)
	a += 1
	*b += 1

	println("normal a", a, "normal b:", *b)

}

func myprint(a int, b *int) {
	println("a:", a, "b:", *b)
}

func fib(r int) int {
	a, b, sum := 0, 1, 0
	for i := 1; i <= r; i++ {
		println(a)
		sum += a
		a, b = b, a+b
	}
	return sum
}

// defer , defers the execution to the end of the stack frame
