package main

import (
	"fmt"
	"time"
)

// QuickSort function: Sorts an array using the QuickSort algorithm.
// It recursively partitions the array and sorts the subarrays.
func QuickSort(arr *[]int, left, right int) {
	if left < right {
		// Get the partition index (pivot position)
		pivot := (*arr)[right]
		// Index for the smaller element
		i := left - 1

		// Traverse the array from low to high-1
		for j := left; j < right; j++ {
			// If current element is smaller than the pivot
			if (*arr)[j] < pivot {
				i++ // Move index forward
				// Swap elements to move smaller ones to the left
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}

		// Swap pivot to its correct position
		(*arr)[i+1], (*arr)[right] = (*arr)[right], (*arr)[i+1]

		// Recursively sort the left subarray (before pivot)
		QuickSort(arr, left, i)

		// Recursively sort the right subarray (after pivot)
		QuickSort(arr, i+2, right)
	}
}

// QuickSortWrapper applies QuickSort and measures execution time
func QuickSortWrapper(arr *[]int) {
	start := time.Now() // Start measuring time

	// Apply QuickSort on the entire array
	QuickSort(arr, 0, len(*arr)-1)

	// Calculate elapsed time
	elapsed := time.Since(start)

	// Print execution time in seconds and milliseconds
	fmt.Printf("QuickSort Execution Time: %.6f seconds (%.3f ms)\n", elapsed.Seconds(), float64(elapsed.Nanoseconds())/1e6)
}
