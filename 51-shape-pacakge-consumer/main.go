package main

import (
	"math/rand/v2"

	"github.com/JitenPalaparthi/shapes-pacakge-demo/shape"
	"github.com/JitenPalaparthi/shapes-pacakge-demo/shape/rect"
	"github.com/JitenPalaparthi/shapes-pacakge-demo/shape/square"
)

// init can be a package level constructor/ first call
func init() { // it is a special function, generally developer has no control over the sequence of execution of init func
	println("first init from main")
}

func init() { // it is a special function, generally developer has no control over the sequence of execution of init func
	println("second init from main")
}

func init() { // it is a special function, generally developer has no control over the sequence of execution of init func
	println("third init from main")
}

var (
	Global = rand.IntN(1000) // evaluated even before main
)

func init() {
	println(Global) // called even before main,
}

// Expressions and exdecutions can be done even before main is called .. main.main
// main is a package and main is a func --> main.main
// before calling main.main few things can be called .
func main() {
	shapeSlice := make([]shape.IShape, 0) // IShape is from a pacakge called shape
	shapeSlice = append(shapeSlice, rect.NewRect(10.3, 14.5), rect.NewRect(10.4, 9.45), square.NewSquare(12.3), square.NewSquare(34.3))
	for _, sh := range shapeSlice {
		shape.Shape(sh)
		println()
	}
	shape.SayHi()
}
