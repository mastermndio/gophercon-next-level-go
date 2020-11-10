package main

import "fmt"

func main() {
	c := make(chan string, 3)

	c <- "test"
	c <- "aardvark"
	c <- "aardvark"
	c <- "aardvark"

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
