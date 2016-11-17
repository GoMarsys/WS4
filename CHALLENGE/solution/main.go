package main

import "fmt"

func main() {
	for i := range generate(25) {
		fmt.Println(factorial(uint64(i)))
	}
}

func generate(n int) chan int {
	channel := make(chan int)
	go func() {
		for i := 3; i <= n; i++ {
			channel <- i
		}
		close(channel)
	}()
	return channel
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
