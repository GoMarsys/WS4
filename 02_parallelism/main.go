package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go counter("dog")
	go counter("cat")
	wg.Wait()
}

func counter(animal string) {
	for i := 0; i < 40; i++ {
		fmt.Println(animal, ":", i)
	}
	wg.Done()
}
