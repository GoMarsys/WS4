package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go counter("dog")
	go counter("cat")
}

func counter(animal string) {
	for i := 0; i < 40; i++ {
		fmt.Println(animal, ":", i)
	}
	wg.Done()
}
