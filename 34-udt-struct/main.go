package main

import "fmt"

func main() {

	var emp1 Employee
	emp1.ID = 101
	emp1.Name = "Jiten"
	emp1.Email = "JitenP@Outlook.Com"
	emp1.Mobile = "9618558500"
	emp1.Status = "active"

	fmt.Println(emp1)

	var emp2 = Employee{102, "Jiten", "JitenP@Outlook.com", "9618558500", "Active"}
	fmt.Println(emp2)
	emp3 := Employee{ID: 103, Name: "Jiten", Email: "JitenP@Outlook.Com"}
	fmt.Println(emp3)

	// pointer of Employee variable

	emp4 := new(Employee)

	(*emp4).ID = 101
	(*emp4).Name = "Jiten" // since emp4 is a pointer can given (*emp4) but without that also works like in arrauys
	emp4.Email = "JitenP@Outlook.Com"
	emp4.Mobile = "9618558500"
	emp4.Status = "active"
	fmt.Println(*emp4) // or emp4

	emp5 := &Employee{ID: 103, Name: "Jiten", Email: "JitenP@Outlook.Com"}
	fmt.Println(*emp5)

}

type Employee struct {
	ID     int
	Name   string
	Email  string
	Mobile string
	Status string
}
