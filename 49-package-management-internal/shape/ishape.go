package shape

import "fmt"

type IShape interface {
	IArea
	Perimeter() float64
	IWhat
}

type IArea interface {
	Area() float64
}

type IWhat interface {
	What() string
}

func Shape(ishape IShape) { // kind of dependency injection
	fmt.Printf("Area of %s:%.2f\n", ishape.What(), ishape.Area())
	fmt.Printf("Perimeter of %s:%.2f\n", ishape.What(), ishape.Perimeter())
}

// The below can be accessed
type Circle struct { // This is is not exported, restricted
	R float32 // restricted
}

// The below cannot be accesses due to starts with lowecase
type circle struct { // This is is not exported, restricted
	r float32 // restricted
}

func greet() {
	println("Hello from shape pacakge")
}

func Greet() {
	println("Hello from shape pacakge")
}
