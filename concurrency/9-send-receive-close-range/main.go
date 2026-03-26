package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	go Fib(25, ch)
	go Rec(ch, sig)
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
	close(ch) // close is just a lock, not to make chan as nil
}

func Rec(ch <-chan int, sig chan<- struct{}) {
	// for {
	// 	v, ok := <-ch // 1 value two .. the second one is bool .. if false chan is closed
	// 	if !ok {      // the ok is false, that means the channel is closed
	// 		sig <- struct{}{}
	// 		close(sig)
	// 		runtime.Goexit()
	// 	} else {
	// 		println(v)
	// 	}
	// }

	for v := range ch { // range on chan gives only 1 value , the received value, range iterates until the chanl is closed
		println(v)
	}
	sig <- struct{}{}
}
