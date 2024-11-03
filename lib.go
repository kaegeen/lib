package main

import "fmt"

// Function to perform linear search
func linearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i // Return index if the target is found
		}
	}
	return -1 // Return -1 if the target is not found
}

func main() {
	var n, target int

	// Prompt the user to enter the number of elements
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n) // Create an array with 'n' elements

	// Input elements of the array
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Input the target element to search for
	fmt.Print("Enter the target element to search for: ")
	fmt.Scan(&target)

	// Perform linear search
	result := linearSearch(arr, target)

	// Display result
	if result != -1 {
		fmt.Printf("Element %d found at index %d.\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array.\n", target)
	}
}
