package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"time"
)

func main() {

	ch := make(chan string)

	go Sender("sender-1", ch, time.Second*5, time.Millisecond*500)

	for c := range ch {
		println(c)
	}

}

func Sender(name string, ch chan<- string, d time.Duration, p time.Duration) {
	timeout := TimeOut(d)
	for {
		time.Sleep(p)
		select {
		case ch <- fmt.Sprint(name, "--->", rand.IntN(10000)):
		case <-timeout:
			close(ch)
			runtime.Goexit()
		}
	}
}

func TimeOut(d time.Duration) <-chan struct{} {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(d)
		timeout <- struct{}{}
		close(timeout)
	}()
	return timeout
}
