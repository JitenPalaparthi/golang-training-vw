package main

import "math/rand/v2"

func main() {

	day := uint8(rand.IntN(10))

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("noday")
	}

}

// No need to give break in go, which is auto break
// Scenarios when break is avoided in other prgramming Langs?
