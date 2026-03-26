package main

import "os"

func main() {
	filename := "data.txt"
	func() {
		func() {
			f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0644)
			if err != nil {
				panic("custom panic:" + err.Error()) // user defined panic
			}
			defer f.Close()
			f.Write([]byte("Hello Wrold"))
		}()
	}()
	println("End of main")
}
