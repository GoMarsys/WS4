package main

import "fmt"

func main() {
	counter("dog")
	counter("cat")
}

func counter(animal string) {
	for i := 0; i < 40; i++ {
		fmt.Println(animal, ":", i)
	}
}
