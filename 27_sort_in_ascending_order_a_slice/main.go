package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// GenerateNumbers creates 100,000 random integers and saves them to a file
func GenerateNumbers(filename string, count int) {
	rand.Seed(time.Now().UnixNano())

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 0; i < count; i++ {
		fmt.Fprintln(file, rand.Intn(100000)) // Random numbers from 0 to 999999
	}

	fmt.Println("Generated", count, "random numbers in", filename)
}

// ReadNumbers reads integers from a file
func ReadNumbers(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, scanner.Err()
}

// SaveNumbers saves an array of numbers to a file
func SaveNumbers(filename string, numbers []int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, num := range numbers {
		fmt.Fprintln(file, num)
	}

	fmt.Println("Sorted numbers saved to", filename)
}

func main() {
	filename := "numbers.txt"

	// Check if numbers.txt exists, otherwise generate numbers
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("numbers.txt not found. Generating random numbers...")
		GenerateNumbers(filename, 100)
	} else {
		fmt.Println("numbers.txt found. Using existing numbers.")
	}

	// Read numbers from file
	fmt.Println("Reading numbers from file...")
	numbers, err := ReadNumbers(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// QuickSort Execution
	numbersCopy := append([]int(nil), numbers...) // Copy slice for fair benchmarking
	fmt.Println("\nApplying QuickSort...")
	QuickSortWrapper(&numbersCopy)
	SaveNumbers("sortedQuickSort.txt", numbersCopy) // Save QuickSort result

}
