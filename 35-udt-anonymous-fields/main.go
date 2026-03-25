package main

import "fmt"

func main() {
	cc1 := ColourCode{1122, "RED", "pure red"}
	cc2 := ColourCode{int: 3232, string: "GREEN", subColour: "light green"} // types become fields
	fmt.Println(cc1, cc2)

}

type subColour = string // This is alias to the string
type ColourCode struct {
	int       // anonymous fields , no name to the filed just the type
	string    // anonymous fileds
	subColour // can use another string by creating an alias
}
