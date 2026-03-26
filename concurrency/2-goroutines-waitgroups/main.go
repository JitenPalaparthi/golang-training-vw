package main

import (
	"runtime"
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup
)

func init() {
	wg = new(sync.WaitGroup) // new wg pointer
}

func main() {
	wg.Add(1) // incrementing the statte
	go func() {
		println("hello world! go routine")
		wg.Done() // decrementing the state
	}()

	wg.Add(1)
	go func() {
		Evennumbers(10)
		wg.Done()
	}()
	wg.Add(1)
	go Oddnumbers(10, wg)

	println("Hello world by main")
	wg.Wait() // wait until? the state becomes 0
	// time.Sleep(time.Millisecond * 10) // not a solution
	//runtime.Goexit() // it crashes but it is an expected behaviour; It is a work around but not a solution
} // if it comes here , it is os.Exit(0)

func Evennumbers(r int) {
	i := 1
	for {
		if i > r {
			break
			//runtime.Goexit() // on other goroutines it does not crash
		}
		time.Sleep(time.Millisecond * 10)
		if i%2 == 0 {
			println("Even->", i)
		}
		i++
	}
}

func Oddnumbers(r int, wg *sync.WaitGroup) {
	i := 1
	for {
		if i > r {
			//break
			wg.Done()
			runtime.Goexit() // on other goroutines it does not crash
		}
		time.Sleep(time.Millisecond * 10)
		if i%2 != 0 {
			println("Odd->", i)
		}
		i++
	}

}
