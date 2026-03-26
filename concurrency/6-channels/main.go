package main

import "time"

func main() {
	var ch chan int // this is a channel , declared but not instantiated.
	if ch == nil {
		println("chan ch is nil")
		ch = make(chan int) // this is unbuffered channel
	}
	go func() {
		time.Sleep(time.Second * 5)
		ch <- 100 // sender is blocked until the value is received
	}()

	v := <-ch // receiver is blocked until the sender sends the value
	println("hello")
	println(v)

}

// channel means a conduit --> pass the data, can be though of a queue that is enqueue and dequeue from both the ends
// sender and a receiver

// buffered and unbuffored
// unbuffered: a channel can only send one element, once that element is received then the sender can send another element
// until the value is receiver , the sender cannot send a value
// until the sender sends a value the receiver cannot receive a value.
// Sener and receiver are blocked until sender sends and receive receives
// both of them be blocked.

// same applies to bufferes channels, there is buffer so sender can send multiple values
// until the buffer is full
// the block can be minimized
