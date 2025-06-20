package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("This program counts the frequency of words in a file using both single-threaded and multi-threaded approaches.")
	singleThread()
	usingThreads()
}

func singleThread() {
	start := time.Now()
	fmt.Println("Program started at:", start.Format("2006-01-02 15:04:05"))

	// Open the file
	file, err := os.Open("word_count.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Initialize a map to store word counts
	wordCount := make(map[string]int)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())  // Convert to lowercase
		line = strings.ReplaceAll(line, ".", "") // Remove periods
		line = strings.ReplaceAll(line, ",", "") // Remove commas
		words := strings.Fields(line)            // Split a line into words

		for _, word := range words {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the word count results
	fmt.Println("Word Frequencies:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

	end := time.Now()
	fmt.Println("Program ended at:", end.Format("2006-01-02 15:04:05"))

	duration := end.Sub(start)
	fmt.Println("Total execution time for the single thread:", duration)
}

// usingThreads reads a file line by line, splits each line into words, and
// counts the frequency of each word using multiple threads. It prints the word
// frequencies and the total execution time.
func usingThreads() {

	start := time.Now()
	fmt.Println("Program started at:", start.Format("2006-01-02 15:04:05"))

	file, err := os.Open("word_count.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	wordCount := make(map[string]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go processLine(line, wordCount, &mutex, &wg)
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Word Frequencies:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

	end := time.Now()
	fmt.Println("Program ended at:", end.Format("2006-01-02 15:04:05"))

	duration := end.Sub(start)
	fmt.Println("Total execution time for the multiple thread:", duration)
}

// processLine takes a line of text, cleans it, and updates the word
// counts in the given map. The map is protected by a mutex for
// safe concurrent access.
func processLine(line string, wordCount map[string]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	// Clean and split line
	line = strings.ToLower(line)
	line = strings.ReplaceAll(line, ".", "")
	line = strings.ReplaceAll(line, ",", "")
	words := strings.Fields(line)

	// Lock the map for safe concurrent access
	mutex.Lock()
	defer mutex.Unlock()
	for _, word := range words {
		wordCount[word]++
	}
}
