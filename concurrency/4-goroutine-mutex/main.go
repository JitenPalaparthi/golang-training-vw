package main

import "sync"

var (
	Counter = 0
	wg      = new(sync.WaitGroup)
	mu      = new(sync.Mutex)
)

func main() {
	wg.Add(2)
	go Incr(wg, mu, 1000)
	go Decr(wg, mu, 1000)
	wg.Wait()
	println("Counter:", Counter)
}

func Incr(wg *sync.WaitGroup, mu *sync.Mutex, r int) {
	for i := 1; i <= r; i++ {
		mu.Lock()
		Counter++
		mu.Unlock()
	}
	wg.Done()
}

func Decr(wg *sync.WaitGroup, mu *sync.Mutex, r int) {
	for i := 1; i <= r; i++ {
		mu.Lock()
		Counter--
		mu.Unlock()
	}
	wg.Done()
}
