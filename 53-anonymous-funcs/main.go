package main

func main() {

	map1 := make(map[string]any)

	map1["add"] = func(a, b int) int {
		return a + b
	}

	var fn1 func(int, int) int = func(a, b int) int {
		return a - b
	}

	map1["sub"] = fn1 // create a variable and pass that

	map1["mul"] = mul // named function

	map1["isCalc"] = true                                  // bool since map can be given any balue
	map1["name"] = "map that is using anonymous functions" // string since map can be given any value

	var fn2 Fn = func(i1, i2 int) int {
		return i1 / i2
	}

	map1["div"] = fn2

	map1["greet"] = func() {
		println("Hello World")
	}

	a, b := 20, 5

	for k, v := range map1 {
		//fmt.Println("key:", k, "value:", v, "type:", reflect.TypeOf(v))

		switch f := v.(type) {
		case bool:
			println("Key:", k, "Value:", f)
		case string:
			println("Key:", k, "Value:", f)
		case func():
			println("Key:", k)
			f()
		case Fn:
			println("Key:", f)
			r := f(a, b)
			println("result of Fn:", r)
			f.What() // a function can have another function since the func is a type
		case func(int, int) int:
			println("Key:", f)
			r := f(a, b)
			println("result of Fn:", r)
		default:
			println("not mentioned type")
		}
	}

}

type Fn func(int, int) int

func (f Fn) What() string {
	return "This is a type called Fn func(int, int) int"
}

func mul(a, b int) int {
	return a * b
}
