package main

import "fmt"

func main() {
	channel := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			channel <- i
		}
		done <- true
	}()

	go func() {
		<-done
		close(channel)
	}()

	for i := range channel {
		fmt.Println(i)
	}

}

// factorial challenge next!
