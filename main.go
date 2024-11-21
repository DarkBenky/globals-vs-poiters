package main

import (
	"fmt"
	"time"
)

const arraySize = 1_000_000

// Global array
var globalArray [arraySize]int

// Modify using global variable
func modifyGlobalArray() {
	for i := 0; i < arraySize; i++ {
		globalArray[i]++
	}
}

// Modify using pointer to array
func modifyPointerArray(arr *[arraySize]int) {
	for i := 0; i < arraySize; i++ {
		arr[i]++
	}
}

func main() {
	// Local array
	var localArray [arraySize]int

	// Test global variable access
	start := time.Now()
	modifyGlobalArray()
	fmt.Println("Global array modification time:", time.Since(start))

	// Test pointer access
	start = time.Now()
	modifyPointerArray(&localArray)
	fmt.Println("Pointer array modification time:", time.Since(start))
}
