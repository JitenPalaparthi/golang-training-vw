package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch1, ch2, out := make(chan string), make(chan string), make(chan string)
	sig := make(chan struct{})
	go Fib(25, "sender-1", ch1)
	go Fib(25, "sender-2", ch2)
	go Rec(ch1, ch2, out, sig)

	for v := range out {
		println(v)
	}

	<-sig // This would wait until the value is received
	// what are you doing with value that is received from sig
}

func Fib(r int, name string, ch chan<- string) {
	a, b := 0, 1
	for i := 1; i <= r; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- fmt.Sprint(name, "-->", a)
		a, b = b, a+b
	}
	close(ch) // close is just a lock, not to make chan as nil
}

func Rec(ch1, ch2 <-chan string, out chan<- string, sig chan<- struct{}) {
	count := 0
	for {
		if count == 2 {
			close(out)
			sig <- struct{}{}
			runtime.Goexit()
		}
		select {
		case v, ok := <-ch1:
			if ok {
				println(v)
				out <- fmt.Sprint("out-->", v)
			} else {
				count++
			}
		case v, ok := <-ch2:
			if ok {
				println(v)
				out <- fmt.Sprint("out-->", v)
			} else {
				count++
			}
		}
	}
}
