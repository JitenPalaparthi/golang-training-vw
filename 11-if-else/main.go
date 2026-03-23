package main

import (
	"math/rand/v2"
	"strconv"
)

func main() {
	a, b := 10, 5
	if a+b+(a+b)/2-10+(a*b)+(a/b)/2+10*2 > 100 {
		println("the whole expression is more than 100")
	} else {
		println("the whole expression is less than 100")
	}

	if a+b+(a+b)/2-10+(a*b)+(a/b)/2+10*2 > 100 && true || false {
		println(true)
	} else {
		println(false)
	}

	// --

	age := uint8(rand.IntN(40))
	gender := GiveMagicGender(age)

	if age >= 18 && (gender == 'F' || gender == 'f') {
		println("she is eligible for marraige according to India", age)
	} else if age >= 21 && (gender == 'M' || gender == 72) {
		println("he is eligible for marriage according to India", age)
	} else {
		println("not eligible", age)
	}

	// What is different

	str1 := "23434w"

	if val, err := strconv.Atoi(str1); err != nil {
		println(err.Error())
	} else {
		println(val)
	}

	a, b = rand.IntN(10), rand.IntN(15)
	if v := a + b + (a+b)/2 - 10 + (a * b) + (a/b)/2 + 10*2; v > 150 {
		println("the whole expression is more than 100", v)
	} else {
		println("the whole expression is less than 100", v)
	}

}

func GiveMagicGender(age uint8) rune {
	if age%3 == 0 {
		return 'M'
	}
	return 'f'
}
