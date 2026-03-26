package main

func main() {
	// r1 := NewRect(10.3, 14.5)
	// r2 := NewRect(10.4, 9.45)
	// s1 := NewSquare(12.3)
	// s2 := NewSquare(34.3)

	shapeSlice := make([]IShape, 0)
	shapeSlice = append(shapeSlice, NewRect(10.3, 14.5), NewRect(10.4, 9.45), NewSquare(12.3), NewSquare(34.3))
	for _, sh := range shapeSlice {
		Shape(sh)
		println()
	}

}

// Rect, Square
