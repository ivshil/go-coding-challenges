package main

// ref https://www.hackerrank.com/challenges/mark-and-toys
func maximumToys(prices []int32, k int32) int32 {
	n := len(prices)

	quicksort(prices, 0, n-1)

	count := int32(0)

	for i := 0; i < n; i++ {
		if k >= prices[i] {
			k -= prices[i]
			count++
		} else {
			break
		}
	}

	return count
}

func quicksort(arr []int32, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
	}
}

func partition(arr []int32, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
