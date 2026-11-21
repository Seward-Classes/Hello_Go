package main

import (
	"fmt"
	"sync"
)

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello, %s! Greetings from a goroutine.\n", name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go greet("Ada", &wg)
	go greet("Linus", &wg)

	wg.Wait()
	fmt.Println("Both goroutines finished.")
}
