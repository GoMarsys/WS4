package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutex sync.Mutex

var count int

func main() {
	wg.Add(2)
	go counter("dog")
	go counter("cat")
	wg.Wait()
}

func counter(animal string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		mutex.Lock()
		count++
		fmt.Println(animal, ":", count)
		mutex.Unlock()
	}
	wg.Done()
}

// How to fix the race condition ?
