package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Write a program that spawns many goroutines that increment a value
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	inc := 0

	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			incCopy := inc
			runtime.Gosched()
			incCopy++
			inc = incCopy
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(inc)

}
