package main

import "math/rand/v2"

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	i := 1
	a, b := 0, 1

	for i <= 10 { // for loop but act like while loop
		print(a, " ")
		a, b = b, a+b
		i++
	}
	println()

	i = 1

	for ; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	println()
	for i := 1; i <= 10; {
		if i%2 == 0 {
			i++
			continue
		}
		print(i, " ")
		i++
	}
	println()

	max := 0
	for { // unconditional loop
		i = rand.IntN(50)
		print(i, " ")
		if i > max {
			max = i
		}
		if i%5 == 0 {
			break // break the loop
		}
	}
	println("\nmax:", max)

}
