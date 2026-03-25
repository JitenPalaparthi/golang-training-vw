package main

import "fmt"

func main() {

	var (
		num  int    = 10 // ToString,Sq and Cube on int by type casting
		num1 myint1 = 20
		num2 myint2 = 30
		num3 myint3 = 40
	)

	fmt.Println("num to string:", myint1(num).ToString())
	fmt.Println("num to Sq:", myint2(num).Sq())
	fmt.Println("num to Cube:", myint3(num).Cube())

	fmt.Println("num1 to string:", num1.ToString())
	fmt.Println("num1 to Sq:", myint2(num1).Sq())
	fmt.Println("num1 to Cube:", myint3(num1).Cube())

	fmt.Println("num2 to string:", myint1(num2).ToString())
	fmt.Println("num2 to Sq:", num2.Sq())
	fmt.Println("num2 to Cube:", myint3(num2).Cube())

	fmt.Println("num3 to string:", myint1(num3).ToString())
	fmt.Println("num3 to Sq:", myint2(num3).Sq())
	fmt.Println("num3 to Cube:", num3.Cube())

	var float1 = 3.23 // can call the functions but may lose data due to casting

	fmt.Println("float1 to string:", myint1(float1).ToString())
	fmt.Println("float1 to Sq:", myint2(float1).Sq())
	fmt.Println("float1 to Cube:", myint3(float1).Cube())

}

type integer = int // this is not a new type., it is an alias

type myint1 int // myint1 is a new type
// benefit of creating such types is can attach methods

func (m myint1) ToString() string {
	return fmt.Sprint(m)
}

type myint2 int

func (m myint2) Sq() int {
	return int(m * m)
}

type myint3 myint2

func (m myint3) Cube() int {
	return int(m) * myint2(m).Sq()
	//return int(m * m * m)
}

//type mymap map[string]any
