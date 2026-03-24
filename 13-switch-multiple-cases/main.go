package main

import "math/rand/v2"

func main() {

	char := uint8(rand.IntN(256))

	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is vovel")
	default:
		println(string(char), "is either consonent or a special char")
	}
}

// No need to give break in go, which is auto break
// Scenarios when break is avoided in other prgramming Langs?
