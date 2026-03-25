package main

import "fmt"

func main() {

	r1 := Rect{L: 12.3, B: 13.56}
	a1 := Area(r1)      // call or pass by value
	p1 := Perimeter(r1) // call or pass by value
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", a1, p1)

	// Suprise values
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", r1.A, r1.P)

	a2 := AreaP(&r1)      // call or pass by ref
	p2 := PerimeterP(&r1) // / call or pass by ref
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", a2, p2)

	//
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", r1.A, r1.P)
	println("\nMethods\n")

	r2 := Rect{L: 12.3, B: 13.56}
	a3 := r2.Area()      // . is calling a method or a field, here it is a method
	p3 := r2.Perimeter() // . is calling a method or a field, here it is a method
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", a3, p3)
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", r2.A, r2.P)

	// just because the reicever is pointer , no need to create a pointer
	a4 := r2.AreaP()      // . is calling a method or a field, here it is a method
	p4 := r2.PerimeterP() // . is calling a method or a field, here it is a method
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", a4, p4)
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", r2.A, r2.P)

	// becase the receiver is a pointer no need to use */& kind of pointer semantics. Go does internally

	a5, p5 := (&r2).Area(), (&r2).Perimeter() // no need to use (&r2) just use r2, the compiler can understd the method uses pointer or normal receiver
	fmt.Printf("Area:%.2f Perimeter:%.2f\n", a5, p5)

	// create a poiter Rect

	r3 := new(Rect)
	r3.L = 654.54
	r3.B = 787.3

	fmt.Printf("Area:%.2f Perimeter:%.2f\n", r3.Area(), r3.Perimeter())
}

func Area(r Rect) float64 {
	r.A = float64(r.L * r.B) // r.A is a local variable for this function and it stores on inside
	return r.A
}

func Perimeter(r Rect) float64 {
	r.P = 2 * float64(r.L+r.B)
	return r.P
}

func AreaP(r *Rect) float64 {
	r.A = float64(r.L * r.B) // r.A is a local variable for this function and it stores on inside
	return r.A
}

func PerimeterP(r *Rect) float64 {
	r.P = 2 * float64(r.L+r.B)
	return r.P
}

type Rect struct {
	L, B float32
	A, P float64
}

// methods are not created inside the type . Rather methods are attached to the type
// methods contain receiver
// each method can have only one receiver which is the type variable for which the method is to be attached

// two kinds of receivers.. pointer vs normal

func (r Rect) Area() float64 { // normal receiver which is call by value
	r.A = float64(r.L * r.B) // r.A is a local variable for this function and it stores on inside
	return r.A
}

func (r Rect) Perimeter() float64 { // normal receiver which is call by value // (r Rect) receiver
	r.P = 2 * float64(r.L+r.B)
	return r.P
}

// pointer receivers
// when to use it? 99% we use pointer receivers.. Particularly if the method is mutating the variable/structre then use pointer receiver

func (r *Rect) AreaP() float64 { // normal receiver which is call by value
	r.A = float64(r.L * r.B) // r.A is a local variable for this function and it stores on inside
	return r.A
}

func (r *Rect) PerimeterP() float64 { // normal receiver which is call by value // (r Rect) receiver
	r.P = 2 * float64(r.L+r.B)
	return r.P
}
