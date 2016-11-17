package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			chan1 <- i
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			chan2 <- i
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			chan3 <- i
		}
	}()

	for {

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
