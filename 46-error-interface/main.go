package main

import (
	"fmt"
	"os"
)

func main() {
	fw := New("data.txt")
	_, err := fmt.Fprintln(fw, "Hello World") // writtn on file not on stdout
	if err != nil {
		//println(err.Error())
		fe, ok := err.(*FileError)
		if ok {
			fmt.Println(fe.Msg)
		} else {
			fmt.Println("not be able to asserted-->", err.Error())
		}

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
	f, err := os.OpenFile(fw.FileName, os.O_RDWR|os.O_APPEND, 0644) // 0777
	if err != nil {
		return 0, NewFileError(101, "File:"+fw.FileName+" is not available")
	}
	defer f.Close() // Will explain defer later

	// n, err = f.Write(p)
	// return n, err
	return f.Write(p)
}

type FileError struct {
	Code int
	Msg  string
}

func NewFileError(code int, msg string) *FileError {
	return &FileError{code, msg}
}

func (fe *FileError) Error() string {
	return fmt.Sprintf("Code:%d Message:%s", fe.Code, fe.Msg)
}

// type error interface {
//     Error() string
// }

// Dynamic Dispatch
// At runtime .. it checks for the required functions and execute

// virtual tables (in c++ VTables)
// in go itables --> For every concrete type , if a dynamic dispatch is implemented
// it maints a in memory table with concrete type , function pointer
