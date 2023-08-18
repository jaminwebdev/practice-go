package utils

import "fmt"

// Sums any number of integers
func Add(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// Prints any integer you pass to it
func PrintNum(num int) {
	fmt.Println("Current number:", num)
}
