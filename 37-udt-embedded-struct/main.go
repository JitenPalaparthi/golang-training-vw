package main

import "fmt"

func main() {

	var emp1 Employee
	emp1.ID = 101
	emp1.Name = "Jiten"
	emp1.Email = "JitenP@Outlook.Com"
	emp1.Mobile = "9618558500"
	emp1.Status = "active"
	emp1.Address.City = "Bangalore"
	emp1.Address.PinCode = "560086"

	fmt.Println(emp1)

	emp2 := Employee{ID: 102, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Address: struct {
		City    string
		PinCode string
	}{City: "Trivandrum", PinCode: "695401"}}

	fmt.Println(emp2)
}

type Employee struct {
	ID      int
	Name    string
	Email   string
	Mobile  string
	Status  string
	Address struct { // embedded struct
		City, PinCode string
	}
}
