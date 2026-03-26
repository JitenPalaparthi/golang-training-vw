package main

import "os"

func main() {
	filename := "data.txt"
	defer func() {
		println("Some function...main")
	}()

	func() {
		defer func() {
			println("Some function...func1")
		}()
		func() {
			defer func() {
				println("Some function...func1.func1")
			}()
			f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0644)
			if err != nil {
				panic("custom panic:" + err.Error()) // user defined panic
			}
			defer func() {
				println("Some function...func1.func1--after panic")
			}()
			defer f.Close()
			f.Write([]byte("Hello Wrold"))
		}()
	}()
	println("End of main")
}

// panic panics, but before going to be panic, it looks for deferred functions.
// all deferes before panic are executed and then only it panics
