package main

import (
	"sync"
)

var (
	wg *sync.WaitGroup
)

func init() {
	wg = new(sync.WaitGroup) // new wg pointer
}

func main() {
	wg.Add(2)
	go func() {
		sum := Fib(10)
		println("Sum:", sum)
		wg.Done()
	}()

	//wg.Add(1)
	go func() {
		sum := Fib(20)
		println("Sum:", sum)
		wg.Done()
	}()
	println("End Of main")
	wg.Wait() // wait until? the state becomes 0
}

func Fib(r int) int {
	a, b, sum := 0, 1, 0
	for i := 1; i <= r; i++ {
		sum += a
		print("a: ", a)
		a, b = b, a+b
	}
	println()
	return sum
}
