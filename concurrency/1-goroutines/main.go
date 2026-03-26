package main

import (
	"runtime"
	"time"
)

func main() {
	go func() {
		println("hello world! go routine")
	}()

	go Evennumbers(10)

	println("Hello world by main")
	// time.Sleep(time.Millisecond * 10) // not a solution
	runtime.Goexit() // it crashes but it is an expected behaviour; It is a work around but not a solution
} // if it comes here , it is os.Exit(0)

func Evennumbers(r int) {
	i := 1
	for {
		if i > r {
			runtime.Goexit() // on other goroutines it does not crash
		}
		time.Sleep(time.Second * 1)
		if i%2 == 0 {
			println(i)
		}
		i++
	}
}

// 1. main is also a goroutine
// 2. no goroutine wait for other goroutines to complete its execuition
// 3. the order of execution is cannot guaranteed

// csp by Tony Hoare
// Do not communicate by sharing memory; instead, share memory by communicating.
// share memory by communicating --> channels

// Go we dont create threads, go routines
// go routines small in size, can create technically in few thousands, they just need 2kb to start with
// generally threads require 2mb.
// to create a go routine just use a keyword called go

// there is a os scheduler that schedules the threads.
// m:n multiplexing -> 4000+ threads can be scheduled to run on 12 threads (this is exp on my machine)
// software threads are created by OS, managed by OS.
// To create a OS thread lot of lowlevel primitives taken place,lot syscalls.
// it leads to more latency

// goroutines are green threads
// go routines are managed by Goruntime scheduler
// When run an applicaiton Go created threads(equal to number cores -> 12 on machine), in general
// ex. 12 threads are created

// G --> Gorotine
// T -> Thread
// P -> Process(not OS process Goruntime creates a process)

// when you run a goroutine, it would be scheduled on a thread, generally on a caller threads
// a process maintains a local queue and a global queue
// if a goroutine is blocked it is removed and put back into global queue.
// blocked means due to io call or waiting for epoll on the os.Simialr to async await kind of system
// if any io call , the goroutine is registered on netpool (epool), and once it gets information from os
// it is scheduled back on to the globl queues --> one of the threds would pick it and execute
// if all threads are blocked by any operation, go runtime creates new threads..
// if the work is over the newly created threads are removed.

// work stealing --> if a thread is very busy and the local queue is filled, other threads are free,
// then they steal work from the fello thread local queue

// These algos are keep on being upgraded or refactored so there might be little changes
