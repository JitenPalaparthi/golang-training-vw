package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 100, 100)
	FillSlice(&slice1)
	// Find the repeated elements in the slice

	fmt.Println(slice1)

	map1 := make(map[int]int, 100)

	for _, v := range slice1 {
		c, ok := map1[v]
		if ok {
			map1[v] = c + 1
		} else {
			map1[v] = 1
		}
		// can do the below since the default value is 0 if no key existed
		// c := map1[v]
		// map1[v] = c + 1
	}

	fmt.Println(map1)

	keys, values := GetMapKeysNValues(map1)

	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)
	k, v := GetMaxRepeated(map1)
	fmt.Println("max repeated Key:", k, "Max Repeated Number:", v)

	keys, values = GetMaxRepeatedPairs(map1, v)
	fmt.Println("max repeated keys and values:", keys, values)

}

func FillSlice(slice *[]int) {
	if slice != nil {
		for i := range *slice {
			(*slice)[i] = rand.IntN(100)
		}
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

func GetMapKeysNValues(m map[int]int) ([]int, []int) {
	keys, values := make([]int, len(m)), make([]int, len(m))
	index := 0
	for k, v := range m {
		keys[index] = k
		values[index] = v
		index++
	}
	return keys, values
}

func GetMaxRepeated(m map[int]int) (key, value int) {
	for k, v := range m {
		if v > value {
			value = v
			key = k
		}
	}
	return key, value
}

func GetMaxRepeatedPairs(m map[int]int, value int) ([]int, []int) {
	var keys, values []int
	for k, v := range m {
		if v == value {
			keys = append(keys, k)
			values = append(values, v)
		}
	}
	return keys, values
}
