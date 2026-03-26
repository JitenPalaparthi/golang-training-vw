package main

import "sync"

var Counter = 0

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go Incr(wg, 1000)
	go Decr(wg, 1000)

	wg.Wait()
	println("Counter:", Counter)
}

func Incr(wg *sync.WaitGroup, r int) {
	for i := 1; i <= r; i++ {
		Counter++
	}
	wg.Done()
}

func Decr(wg *sync.WaitGroup, r int) {
	for i := 1; i <= r; i++ {
		Counter--
	}
	wg.Done()
}
