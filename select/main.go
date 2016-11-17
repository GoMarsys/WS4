package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	var counter int32
	counter = 3

	go func() {
		for i := 0; i < 20; i++ {
			chan1 <- i
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		}
		atomic.AddInt32(&counter, -1)
	}()

	go func() {
		for i := 0; i < 20; i++ {
			chan2 <- i
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		}
		atomic.AddInt32(&counter, -1)
	}()

	go func() {
		for i := 0; i < 20; i++ {
			chan3 <- i
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		}
		atomic.AddInt32(&counter, -1)
	}()

	for counter != 0 {

		select {

		case msg_chan1 := <-chan1:
			fmt.Println("Chan1:", msg_chan1)

		case msg_chan2 := <-chan2:
			fmt.Println("Chan2:", msg_chan2)

		case msg_chan3 := <-chan3:
			fmt.Println("Chan3:", msg_chan3)
		default:
		}

	}

}
