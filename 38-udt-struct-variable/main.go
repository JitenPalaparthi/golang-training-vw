package main

import "fmt"

func main() {

	var emp struct {
		Id          int
		Name, Email string
	} // not a struct it is just a variable

	//var fn1 func(string) string

	emp = struct {
		Id          int
		Name, Email string
	}{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com"}

	fmt.Println(emp)

}
