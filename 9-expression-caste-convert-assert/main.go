package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		num1 int     = 36764545
		num2 uint8   = 39
		num3 float32 = 65.56
		num4 float64 = 7645645.4354353
		num5 uint64  = 4435345353
		num6 int32   = 234324
		ok1          = true // if true add 1
		str1         = "42323423"
		str2         = "42323423.34334"
		any1         = any(544364.456)
		any2 any     = float32(422534443.435)
		any3 any     = uint32(324234)
		sum  float64
	)

	sum = float64(num1) + float64(num2) + float64(num3) + num4 + float64(num5) + float64(num6)
	if ok1 {
		sum += 1 // no need to add 1.0 .. even sum is float64 go can automatically make it as float64 addition
	}

	v1, _ := strconv.Atoi(str1)
	sum += float64(v1)

	v2, _ := strconv.ParseFloat(str2, 64)
	sum += v2

	sum = sum + any1.(float64)
	sum = sum + float64(any2.(float32)) + float64(any3.(uint32))
	fmt.Printf("Sum:%.3f\n", sum)
}
