package main

import (
	"fmt"
	"sync"
)

func WGSample(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func Normal(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go WGSample("WG", &wg)
	Normal("Hi !")

	wg.Wait()
}
