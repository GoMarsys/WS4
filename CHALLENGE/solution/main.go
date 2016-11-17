package main

import "fmt"

func main() {
	numbers := generate(25)
	for i := range factorial(numbers) {
		fmt.Println(i)
	}
}

func generate(n uint64) chan uint64 {
	channel := make(chan uint64)
	go func() {
		for i := uint64(3); i <= n; i++ {
			channel <- i
		}
		close(channel)
	}()
	return channel
}

func factorial(in chan uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		for i := range in {
			total := uint64(1)
			for z := i; z > 0; z-- {
				total *= z
			}
			out <- total
		}
		close(out)
	}()
	return out
}
