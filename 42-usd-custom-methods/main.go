package main

import "fmt"

func main() {

	var mm mymap // custom map declaration

	mm = make(mymap) // make works even for the types which are custom created from base types

	mm["560096"] = "Bangalore-2"
	mm["560034"] = "Bangalore-3"
	mm["560062"] = "Bangalore-4"
	mm["695401"] = "Trivandrum-1"
	mm["xyz"] = 3121
	mm["hasdata"] = true
	fmt.Println("Len of map", len(mm))

	if err := mm.Delete("xyz"); err != nil {
		println(err.Error())
	} else {
		println("key succsssfully deletes")
	}

	keys, values := mm.GetMapKeysNValues()
	fmt.Println("keys:", keys)
	fmt.Println("values:", values)

	normalmap := make(map[string]any) // base datatype, normal map not a user defined map

	normalmap["560096"] = "Bangalore-2"
	normalmap["560034"] = "Bangalore-3"
	normalmap["560062"] = "Bangalore-4"
	normalmap["695401"] = "Trivandrum-1"
	normalmap["xyz"] = 3121
	normalmap["hasdata"] = true

	if err := mymap(normalmap).Delete("xyz"); err != nil {
		println(err.Error())
	} else {
		println("key succsssfully deletes")
	}

	keys, values = mymap(normalmap).GetMapKeysNValues()
	fmt.Println("keys:", keys)
	fmt.Println("values:", values)

}

type mymap map[string]any

func (mm mymap) Delete(key string) error {
	if mm == nil {
		return fmt.Errorf("nil map")
	}
	_, ok := mm[key]
	if !ok {
		return fmt.Errorf("key:%v does not exist", key)
	}
	delete(mm, key)
	return nil
}

func (mm mymap) GetMapKeysNValues() ([]string, []any) {
	keys, values := make([]string, len(mm)), make([]any, len(mm))
	index := 0
	for k, v := range mm {
		keys[index] = k
		values[index] = v
		index++
	}
	return keys, values
}
