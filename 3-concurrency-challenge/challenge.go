package main

import "fmt"

// Create mapping stucture?
type finalLetterMap map[rune]int

// Create a application that takes a list of strings and lets us know how many times each letter has occured. This should be complete map rather than one for each string.
func CountLetterFrequency(text string) finalLetterMap {
	ourMap := finalLetterMap{}
	for _, letter := range text {
		ourMap[letter]++
	}

	return ourMap
}

// ConcurrentLetterFrequency Create Concurrent function to use the above function
func ConcurrentLetterFrequency(texts []string) finalLetterMap {
	c := make(chan finalLetterMap)
	for _, item := range texts {
		go func(s string) {
			c <- CountLetterFrequency(s)
		}(item)
	}

	completeMap := finalLetterMap{}
	for range texts {
		for k, v := range <-c { //waiting for channel to deliver map
			completeMap[k] += v
		}
	}

	close(c)
	return completeMap
}

// Create function to create big map from each individual map
func main() {
	texts := []string{"Hello, my name is Aaron",
		"Goodbye, I will see you later",
		"What a beautiful day it is"}

	fmt.Println(ConcurrentLetterFrequency(texts))
}
