package main

import "fmt"

func main() {
	for i := range pusher() {
		fmt.Println(i)
	}
}

func pusher() chan int {
	channel := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			channel <- i
		}
		close(channel)
	}()
	return channel
}
