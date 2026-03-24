package main

func main() {
	i := 1

odd:
	println(i, "is odd number")
	i++
	goto loop

loop:
	println("in main lable-->", i)
	if i > 10 {
		goto exitit
	}
	if i%2 == 0 {
		goto even
	} else {
		goto odd
	}
even:
	println(i, "is even number")
	i++
	goto loop

exitit:
	println("Exiting the goto for even and odd")
}
