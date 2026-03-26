package main

func main() {
	func() { // func1
		func() { //func1.1
			//num := 0
			//println(100 / num) // integer divide by zero
			var ptr *int
			*ptr = 100 // dereferenceing a nil pointer , which leads a panic
		}()
		println("end of func1")
	}()
	println("end of main")
}

// exit status 2
