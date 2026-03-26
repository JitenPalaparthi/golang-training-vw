package main

import (
	"log/slog"
	"os"
)

func main() {
	filename := "data.txt" //fmt.Sprint("data-", rand.Text()) + ".txt"
	fn1 := func() {
		defer func() {
			if r := recover(); r != nil {
				println(r)
				println("recovering from panic")

				f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
				if err != nil {
					slog.Error(err.Error())
					//os.Exit(1)
				}
				defer f.Close()
				f.Write([]byte("Hello Wrold\n"))
			}

		}()

		f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			panic("custom panic:" + err.Error()) // user defined panic
		}
		defer f.Close()
		f.Write([]byte("Hello Wrold"))
	}

	func() {
		fn1()
	}()

	fn2 := func() { // func2, read from the file
		f, err := os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			slog.Error(err.Error())
			//os.Exit(1)
		}
		println("Reading from file")
		defer f.Close()
		buf := make([]byte, 100)
		_, err = f.Read(buf)
		if err != nil && err.Error() != "EOF" {
			println(err.Error())
		} else {
			println(string(buf))
		}
	}
	fn2()

	println("End of main")
}

// panic panics, but before going to be panic, it looks for deferred functions.
// all deferes before panic are executed and then only it panics

// when ever there is a panic, it looks for defer, and in the defer ,can write recover loginc and minimize the crash only to that func
