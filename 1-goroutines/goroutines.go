package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func countItUp(s string) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d - %s\n", i, s)
		time.Sleep(time.Millisecond * 600)
	}
}

func main() {
	people := []string{"Aaron", "Ashly", "Terrance", "Bobby"}
	var wg sync.WaitGroup

	fmt.Println("CPU core count: ", runtime.NumCPU())
	fmt.Println("Goroutine count: ", runtime.NumGoroutine())

	for _, name := range people {
		wg.Add(1)
		go func(s string) {
			countItUp(s)
			wg.Done()
		}(name)
	}

	fmt.Println("Goroutine count: ", runtime.NumGoroutine())
	wg.Wait()
}
