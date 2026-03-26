package main

import (
	"shapes/rect"
	"shapes/shape"
	"shapes/square"
)

func main() {
	shapeSlice := make([]shape.IShape, 0) // IShape is from a pacakge called shape
	shapeSlice = append(shapeSlice, rect.NewRect(10.3, 14.5), rect.NewRect(10.4, 9.45), square.NewSquare(12.3), square.NewSquare(34.3))
	for _, sh := range shapeSlice {
		shape.Shape(sh)
		println()
	}

	// c1 := shape.Circle{R: 100}
	shape.Greet() // can access as it is unrestricted , starts with Uppercase
	//shape.greet() // cant access it due to restricted , starts with lowercase
}

// to create a package ,it should have a directory
// There are no special access specifier/modifiers in go --> No public,private,protected, kind of...
// inside a package , any type/function/method/fields can be called but to be accessed from outside pacakge
// including main , they have to be exported/unrestricted

// Func,Method,Type,Field,Var (Anything) inside a pacakge --> If it starts with UpperCase then
// they are unrestricted that is similar to public
// if not they are restricted that is similar to private
