package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := FibGen(10, "fib-gen-1")
	ch2 := FibGen(10, "fib-gen-2")
	ch3 := FibGen(10, "fib-gen-3")
	ch4 := FibGen(10, "fib-gen-4")
	ch5 := FibGen(10, "fib-gen-5")
	out, sig := FanIn(ch1, ch2, ch3, ch4, ch5)

	go func() {
		for v := range out {
			println(v)
		}
	}()
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

func FanIn(chs ...chan string) (out chan string, sig chan struct{}) {
	out = make(chan string, 10)
	sig = make(chan struct{})
	go func() {
		wg := new(sync.WaitGroup)
		wg.Add(len(chs))
		for _, ch := range chs {
			go func() {
				for c := range ch {
					out <- fmt.Sprint("Out----->", c) // process
				}
				wg.Done()
			}()
		}

		go func() {
			wg.Wait()
			close(out)
			sig <- struct{}{}
		}()

	}()
	return out, sig
}
