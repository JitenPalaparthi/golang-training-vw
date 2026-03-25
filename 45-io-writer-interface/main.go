package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello World") // formatted println
	// io.Writer
	// io.Reader
	fw := New("data.txt")
	_, err := fmt.Fprintln(fw, "Hello World") // writtn on file not on stdout
	if err != nil {
		println(err.Error())
	}

}

// interfaces are auto implemented

type FileWriteOps struct {
	FileName string
}

func New(fn string) *FileWriteOps {
	return &FileWriteOps{fn}
}

// Write(p []byte) (n int, err error)

func (fw *FileWriteOps) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644) // 0777
	if err != nil {
		return 0, err
	}
	defer f.Close() // Will explain defer later

	// n, err = f.Write(p)
	// return n, err
	return f.Write(p)
}

// type error interface {
//     Error() string
// }
