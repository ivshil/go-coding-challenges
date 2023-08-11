package main

import "math"

// ref https://www.hackerrank.com/challenges/2d-array

func hourglassSum(arr [][]int32) int32 {
	numRows := len(arr[0])
	numCols := len(arr[1])
	largestHourglass := int32(math.MinInt32)
	for i := 1; i < numCols-1; i++ {
		for j := 1; j < numRows-1; j++ {
			sum := int32(arr[i-1][j] + arr[i][j] + arr[i+1][j] + arr[i-1][j-1] + arr[i-1][j+1] + arr[i+1][j-1] + arr[i+1][j+1])
			if largestHourglass < sum {
				largestHourglass = sum
			}
		}
	}

	return largestHourglass
}
