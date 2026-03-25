package main

import "fmt"

func main() {

	var mp map[string]any // This is a declaration

	if mp == nil {
		println("nil map")
		mp = make(map[string]any)

	}

	// var mp1 map[[2]int]any // array can be a key
	// var mp2 map[[]int]any  // a slice cannot be a key

	mp["560086"] = "Bangalore-1"
	mp["560096"] = "Bangalore-2"
	mp["560034"] = "Bangalore-3"
	mp["560062"] = "Bangalore-4"
	mp["695401"] = "Trivandrum-1"
	mp["xyz"] = 3121
	mp["hasdata"] = true
	fmt.Println("Len of map", len(mp))

	// to iterate

	for k, v := range mp { // on map range loop gives key and value
		fmt.Println("key:", k, "value:", v)
	}

	v := mp["560086"] // get a value
	fmt.Println("Value:", v)

	v, ok := mp["5435423"] // actually the key does not exist. If ok is true the key exists and get a value
	// if false, the key does not exist and it gives a zero value
	if ok {
		fmt.Println("Value:", v)
	} else {
		println("key does not exist")
	}

	delete(mp, "xyz") // delete is a built in function
	// what if key does not exist , map does nothing , does not return any confirmation or ack
	// if key exists , delete , deletes the pair. If not , it does nothing

	v, ok = mp["xyz"] // actually the key does not exist. If ok is true the key exists and get a value
	// if false, the key does not exist and it gives a zero value
	if ok {
		fmt.Println("Value:", v)
	} else {
		println("key does not exist")
	}

	if err := Delete(mp, "xyz"); err != nil {
		println(err.Error())
	} else {
		println("successfully deleted")
	}

}

func Delete(m map[string]any, key string) error {
	if m == nil {
		return fmt.Errorf("nil map")
	}
	_, ok := m[key]
	if !ok {
		return fmt.Errorf("key:%v does not exist", key)
	}
	delete(m, key)
	return nil
}

// a map can be nil, if declared but not instantiated
// What key be the key --> Any datatype that satisfies == operation can be a key
// map is very efficient bcz on an avz O(1) lookup
// map is not ordered
// map is not thread safe

// arr1 := [2]int{10, 12}
// arr2 := [2]int{10, 12,14}

// if arr1 == arr2 { // Array can be a key since == operation can be performed on array
// }

// slice1 := []int{10, 12}
// slice2 := []int{10, 12}

// if slice1 == slice2 {

// } // == is an invalid operation on slice so , a slice cannot be a key

// take the key
// "Bangalore" --> Hash
// 0XAABB112 23DFF45FF

/*
Bucket --> each bucket contains array of 8 elements
	| -> Top Hash
	| -> Keys
	| -> Values
	| -> Overflow bucket ptrs

The number of buckets are based on LoadFactor -> 6.5
LoadFactor = Total-Elements/Number-of-Buckets
When the load factor is more than defined limit, the number of buckets are doubled.
The currect maps values are rearranged/rebalanced

How does it know which bucket to store --> HashFunction
It does the modula operation and gives a value for a key. That value is the bucket number
*/
