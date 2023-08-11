package main

// ref kerrank.com/challenges/ctci-array-left-rotation
func rotLeft(a []int32, d int32) []int32 {
	n := int32(len(a))
	result := make([]int32, n)

	for i := int32(0); i < n; i++ {
		newIndex := (i - d + n) % n
		result[newIndex] = a[i]
	}

	return result
}
