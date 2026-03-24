package main

import "math/rand/v2"

func main() {
	num := rand.IntN(10000)
	switch { // empty switch
	case num%2 == 0:
		println(num, "is even")
	default:
		println(num, "is odd")
	}

	switch num := rand.IntN(10000); { // delcaration and expression not a condition
	case num > 0 && num <= 100:
		println(num, "is between 0-100")
	case num > 100 && num <= 200:
		println(num, "is between 100-200")
	default:
		println(num, "is greater than 200")
	}

	// a := 10
	// a++
	// if a := rand.IntN(100); a > 10 {

	// }
}

// No need to give break in go, which is auto break
// Scenarios when break is avoided in other prgramming Langs?
