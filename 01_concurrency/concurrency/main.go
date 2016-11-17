package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(1)
}

var wg sync.WaitGroup

func main() {
	go counter("dog")
	go counter("cat")
}

func counter(animal string) {
	for i := 0; i < 40; i++ {
		fmt.Println(animal, ":", i)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	}
}

// go routine definition
// context switch between running go routines
//
