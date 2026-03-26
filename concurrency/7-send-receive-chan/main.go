package main

import "time"

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	go Fib(10, ch)
	go Rec(10, ch, sig)
	<-sig // This would wait until the value is received
	// what are you doing with value that is received from sig
}

func Fib(r int, ch chan<- int) {
	a, b := 0, 1
	for i := 1; i <= r; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- a
		a, b = b, a+b
	}
}

func Rec(r int, ch <-chan int, sig chan<- struct{}) {
	for i := 1; i <= r; i++ {
		println(<-ch)
	}
	sig <- struct{}{}
}
