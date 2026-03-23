package main

func main() {

	a, b := 10, 20

	a, b = b, a

	println(a, b)

	a1, b1, c1 := 10, 20, 30
	a1, b1, c1 = b1, c1, a1
	println(a1, b1, c1)

	c1 = a + b + (a+b)/2 - 10 + (a * b) + (a/b)/2 + 10*2  // arth expression
	c2 := a+b+(a+b)/2-10+(a*b)+(a/b)/2+10*2 > 100         // comparision -> true|false
	c3 := a+b+(a+b)/2-10+(a*b)+(a/b)/2+10*2 > 100 && true // logical expression
	c4 := c1 << 2                                         // bitwise expression

	println(c1, c2, c3, c4)

	c1, _ = 100, 200 // the second value is not stored any where bcz the expression is into _ blank identifer

	_ = a + b + (a+b)/2 - 10 + (a * b) + (a/b)/2 + 10*2

}
