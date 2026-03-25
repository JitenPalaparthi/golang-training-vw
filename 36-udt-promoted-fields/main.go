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
	emp1.Address.Status = "active"

	fmt.Println(emp1)

	// pointer of Employee variable

	emp4 := new(Employee)

	(*emp4).ID = 101
	(*emp4).Name = "Jiten" // since emp4 is a pointer can given (*emp4) but without that also works like in arrauys
	emp4.Email = "JitenP@Outlook.Com"
	emp4.Mobile = "9618558500"
	emp4.Status = "active"
	emp4.Address.City = "Bangalore"
	emp4.Address.PinCode = "560086"
	emp4.Address.Status = "active"

	fmt.Println(*emp4) // or emp4
	fmt.Println("City:", emp4.Address.City)

	var std1 Student
	std1.ID = 101
	std1.Name = "Jiten"
	std1.Email = "JitenP@Outlook.Com"
	std1.Mobile = "9618558500"
	std1.Status = "active"
	std1.City = "Bangalore"        // can directly give or access fields or methods of a promoted field
	std1.PinCode = "560086"        // can directly give or access fields or methods of a promoted field
	std1.Address.Status = "active" // it is ambiguity, it cannot understand what status is asked so gives direlyt student status

	fmt.Println(std1)
	fmt.Println("City:", std1.City)
	fmt.Println("Address status:", std1.Address.Status)

}

type Employee struct {
	ID      int
	Name    string
	Email   string
	Mobile  string
	Status  string
	Address Address // in Go field name can be the type name
}

type Address struct { // another type called address
	City, PinCode, Status string // Status -> active, inactive
}

// Promoted field

type Student struct {
	ID      int
	Name    string
	Email   string
	Mobile  string
	Status  string
	Address // Only give the type, promoted field
}
