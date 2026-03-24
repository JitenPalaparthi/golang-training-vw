package main

func main() {

	for i := range 10 { // range from 0 to 9
		println(i)
	}
	// range loop can only be used in strings, array, slice, map, channel

	//str1 := "Hello World"
	str1 := "Hello, 世界" // collection of bytes

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]))
	}

	println()

	println(len(str1))

	for i, c := range str1 {
		println(i, "-->", string(c))
	}

	for _, c := range str1 {
		println("-->", string(c))
	}

	for i := range str1 {
		println(i)
	}

	for i, v := range getString("World") {
		println(i, " ", string(v))
	}

}

// works the best with collections like arrays, strings, slices, maps, channels
// It iterates will thte end of the collection. Means len, size, close of a channel

func getString(s string) string {
	return "Hello " + s
}
