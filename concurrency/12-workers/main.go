package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := FibGen(20, "fib-gen-1")
	sig := make(chan struct{})
	workers := uint(5)
	go Workers(workers, ch, sig)
	<-sig
}

func FibGen(r int, name string) chan string {
	ch := make(chan string)
	go func() {
		a, b := 0, 1
		for i := 1; i <= r; i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- fmt.Sprint(name, "-->", a)
			a, b = b, a+b
		}
		close(ch) // close is just a lock, not to make chan as nil
	}()
	return ch
}

func Workers(w uint, ch <-chan string, sig chan struct{}) {
	go func() {
		wg := new(sync.WaitGroup)
		wg.Add(int(w))
		for i := 1; i <= int(w); i++ {
			go func(wk int) {
				for v := range ch {
					println("received by --", wk, "-->", v)
				}
				wg.Done()
			}(i)
		}
		go func() {
			wg.Wait()
			sig <- struct{}{}
		}()
	}()
}
