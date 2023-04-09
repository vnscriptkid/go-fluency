package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Define a type for the mapper function
type mapper func(string) map[string]int

// Define a type for the reducer function
type reducer func([]map[string]int) map[string]int

func processFile(file string, m mapper, output chan<- map[string]int, wg *sync.WaitGroup) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Initialize scanner to read the file line by line
	scanner := bufio.NewScanner(f)

	// Process each line in the file using the mapper function
	counts := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		countsMap := m(line)
		for k, v := range countsMap {
			counts[k] += v
		}
	}

	// Send the intermediate results to the output channel
	output <- counts

	// Notify the wait group that the file has been processed
	wg.Done()
}

func mapReduce(files []string, m mapper, r reducer) map[string]int {
	// Initialize wait group and output channel
	var wg sync.WaitGroup
	output := make(chan map[string]int)

	// Process each file using a separate Goroutine
	for _, file := range files {
		wg.Add(1)
		go processFile(file, m, output, &wg)
	}

	// Wait for all Goroutines to complete
	go func() {
		wg.Wait()
		close(output)
	}()

	// Aggregate the intermediate results using the reducer function
	intermediate := []map[string]int{}
	for counts := range output {
		intermediate = append(intermediate, counts)
	}
	finalOutput := r(intermediate)

	return finalOutput
}

func main() {
	// Input files
	files, err := filepath.Glob("input/*.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		fmt.Printf("File %v: %v\n", i+1, file)
	}

	// Define the mapper function
	mapperFunc := func(input string) map[string]int {
		// Split input into words
		words := strings.Fields(input)

		// Initialize map to store word counts
		counts := make(map[string]int)

		// Iterate over words and count occurrences
		for _, word := range words {
			counts[word]++
		}

		return counts
	}

	// Define the reducer function
	reducerFunc := func(counts []map[string]int) map[string]int {
		// Initialize map to store final word counts
		finalCounts := make(map[string]int)

		// Iterate over intermediate counts and aggregate
		for _, m := range counts {
			for k, v := range m {
				finalCounts[k] += v
			}
		}

		return finalCounts
	}

	// Apply MapReduce to the input files
	finalOutput := mapReduce(files, mapperFunc, reducerFunc)

	// Print output
	fmt.Println(finalOutput)
}
