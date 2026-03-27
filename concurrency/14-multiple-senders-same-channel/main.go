package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"
	"time"
)

func main() {

	ch := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go Sender("sender-1", wg, ch, time.Millisecond*100, time.Millisecond*10)
	go Sender("sender-2", wg, ch, time.Millisecond*140, time.Millisecond*10)
	go Sender("sender-3", wg, ch, time.Millisecond*120, time.Millisecond*10)

	go func() {
		wg.Wait() /// wait here ..until the wait group state is zero
		close(ch)
	}()

	for c := range ch {
		println(c)
	}

}

func Sender(name string, wg *sync.WaitGroup, ch chan<- string, d time.Duration, p time.Duration) {
	timeout := TimeOut(d)
	for {
		time.Sleep(p)
		select {
		case ch <- fmt.Sprint(name, "--->", rand.IntN(10000)):
		case <-timeout:
			//close(ch)
			wg.Done()
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

// The sender would never know that the channel is closed.
// if more than one sender sending the data on the same channel,
// one of the goroutine closes the channel, the other senders do not know about it.
// so they try to send data on a closed channel.
