package main

import (
	"fmt"
)

func linearSearch(array []int, size int, num int) {
	for i := 0; i < size; i++ {
		if array[i] == num {
			fmt.Println("Number fount at postion: ", i+1)
			return
		}

	}

	fmt.Println("Number not found")
	return
}

// Main function
func main() {
	fmt.Println("Enter the size of the array")
	var size int
	fmt.Scan(&size)

	fmt.Println("Enter the capacity for make mnadatory argument")
	var capacity int
	fmt.Scan(&capacity)

	var num int
	array := make([]int, size, capacity)

	for i := 0; i < size; i++ {
		fmt.Println("Enter the element:")
		fmt.Scan(&array[i])
	}

	fmt.Println("Enter the number to be searcher")
	fmt.Scan(&num)

	linearSearch(array, size, num)
}
