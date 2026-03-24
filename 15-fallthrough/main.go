package main

func main() {
	num := 48

	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough // dont check the case but just execute the body
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("false negative of misuing fallthrough")
	num = 2

	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough // dont check the case but just execute the body
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}
}

// What are the reasons, you avoid break?
