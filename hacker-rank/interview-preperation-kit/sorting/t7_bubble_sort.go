package main

import "fmt"

// ref https://www.hackerrank.com/challenges/ctci-bubble-sort
func countSwaps(a []int32) {
	n := len(a)
	swaps := 0
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
				swaps++
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[len(a)-1])
}
