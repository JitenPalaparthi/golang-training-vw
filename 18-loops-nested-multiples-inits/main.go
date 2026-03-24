package main

import "math/rand/v2"

func main() {

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			println("i:", i, "j:", j)
		}
	}

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			println("i:", i, "j:", j)
			if j == 5 {
				break // It breaks only the inner loop, not the outer loop
			}
		}
	}

	println("break both the loops")

	done := false
	for i := 1; i <= 10; i++ {
		if done {
			break
		}
		for j := 1; j <= 10; j++ {
			println("i:", i, "j:", j)
			if j == 5 {
				done = true
				break // It breaks only the inner loop, not the outer loop
			}
		}
	}

	println("break both the loops")

outer:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			println("i:", i, "j:", j)
			if j == 5 {
				break outer // It breaks only the inner loop, not the outer loop
			}
		}
	}

exit:
	for {
		num := rand.IntN(1000)
		rf := false
		switch {
		case num%3 == 0:
			println(num, "is divisible by 3")
			rf = true
			fallthrough
		case num%2 == 0:
			if rf {
				if num%2 == 0 {
					println(num, "is divisible by 2")

				} else {
					println(num, "is divisible by 2")
				}
				if num%17 == 0 {
					println(num, "is divisible by 17 so exiting....")
					break exit
				}
			}
		}
	}

	println("Multiple init in for loop")

	for i, j := 1, 10; j >= 5 && i <= 5; i, j = i+1, j-1 {
		println("i-- ", i, "--j ", j)
	}

}
